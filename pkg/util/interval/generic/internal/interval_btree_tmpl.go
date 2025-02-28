// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

//go:build ignore

package internal

import (
	"bytes"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
)

// nilT is a nil instance of the Template type.
var nilT T

const (
	degree   = 16
	maxItems = 2*degree - 1
	minItems = degree - 1
)

// compare returns a value indicating the sort order relationship between
// a and b. The comparison is performed lexicographically on
//
//	(a.Key(), a.EndKey(), a.ID())
//
// and
//
//	(b.Key(), b.EndKey(), b.ID())
//
// tuples.
//
// Given c = compare(a, b):
//
//	c == -1  if (a.Key(), a.EndKey(), a.ID()) <  (b.Key(), b.EndKey(), b.ID())
//	c ==  0  if (a.Key(), a.EndKey(), a.ID()) == (b.Key(), b.EndKey(), b.ID())
//	c ==  1  if (a.Key(), a.EndKey(), a.ID()) >  (b.Key(), b.EndKey(), b.ID())
func compare(a, b T) int {
	c := bytes.Compare(a.Key(), b.Key())
	if c != 0 {
		return c
	}
	c = bytes.Compare(a.EndKey(), b.EndKey())
	if c != 0 {
		return c
	}
	if a.ID() < b.ID() {
		return -1
	} else if a.ID() > b.ID() {
		return 1
	} else {
		return 0
	}
}

// keyBound represents the upper-bound of a key range.
type keyBound struct {
	key []byte
	inc bool
}

func (b keyBound) compare(o keyBound) int {
	c := bytes.Compare(b.key, o.key)
	if c != 0 {
		return c
	}
	if b.inc == o.inc {
		return 0
	}
	if b.inc {
		return 1
	}
	return -1
}

func (b keyBound) contains(a T) bool {
	c := bytes.Compare(a.Key(), b.key)
	if c == 0 {
		return b.inc
	}
	return c < 0
}

func upperBound(c T) keyBound {
	if len(c.EndKey()) != 0 {
		return keyBound{key: c.EndKey()}
	}
	return keyBound{key: c.Key(), inc: true}
}

type node struct {
	ref   int32
	count int16

	// These fields form a keyBound, but by inlining them into node we can avoid
	// the extra word that would be needed to pad out maxInc if it were part of
	// its own struct.
	maxInc bool
	maxKey []byte

	items [maxItems]T

	// The children array pointer is only populated for interior nodes; it is nil
	// for leaf nodes.
	children *childrenArray
}

type childrenArray = [maxItems + 1]*node

var leafPool = sync.Pool{
	New: func() interface{} {
		return new(node)
	},
}

var nodePool = sync.Pool{
	New: func() interface{} {
		type interiorNode struct {
			node
			children childrenArray
		}
		n := new(interiorNode)
		n.node.children = &n.children
		return &n.node
	},
}

func newLeafNode() *node {
	n := leafPool.Get().(*node)
	n.ref = 1
	return n
}

func newNode() *node {
	n := nodePool.Get().(*node)
	n.ref = 1
	return n
}

// mut creates and returns a mutable node reference. If the node is not shared
// with any other trees then it can be modified in place. Otherwise, it must be
// cloned to ensure unique ownership. In this way, we enforce a copy-on-write
// policy which transparently incorporates the idea of local mutations, like
// Clojure's transients or Haskell's ST monad, where nodes are only copied
// during the first time that they are modified between Clone operations.
//
// When a node is cloned, the provided pointer will be redirected to the new
// mutable node.
func mut(n **node) *node {
	if atomic.LoadInt32(&(*n).ref) == 1 {
		// Exclusive ownership. Can mutate in place.
		return *n
	}
	// If we do not have unique ownership over the node then we
	// clone it to gain unique ownership. After doing so, we can
	// release our reference to the old node. We pass recursive
	// as true because even though we just observed the node's
	// reference count to be greater than 1, we might be racing
	// with another call to decRef on this node.
	c := (*n).clone()
	(*n).decRef(true /* recursive */)
	*n = c
	return *n
}

// leaf returns true if this is a leaf node.
func (n *node) leaf() bool {
	return n.children == nil
}

// max returns the maximum keyBound in the subtree rooted at this node.
func (n *node) max() keyBound {
	return keyBound{
		key: n.maxKey,
		inc: n.maxInc,
	}
}

// setMax sets the maximum keyBound for the subtree rooted at this node.
func (n *node) setMax(k keyBound) {
	n.maxKey = k.key
	n.maxInc = k.inc
}

// incRef acquires a reference to the node.
func (n *node) incRef() {
	atomic.AddInt32(&n.ref, 1)
}

// decRef releases a reference to the node. If requested, the method
// will recurse into child nodes and decrease their refcounts as well.
func (n *node) decRef(recursive bool) {
	if atomic.AddInt32(&n.ref, -1) > 0 {
		// Other references remain. Can't free.
		return
	}
	// Clear and release node into memory pool.
	if n.leaf() {
		*n = node{}
		leafPool.Put(n)
	} else {
		// Release child references first, if requested.
		if recursive {
			for i := int16(0); i <= n.count; i++ {
				n.children[i].decRef(true /* recursive */)
			}
		}
		*n = node{children: n.children}
		*n.children = childrenArray{}
		nodePool.Put(n)
	}
}

// clone creates a clone of the receiver with a single reference count.
func (n *node) clone() *node {
	var c *node
	if n.leaf() {
		c = newLeafNode()
	} else {
		c = newNode()
	}
	// NB: copy field-by-field without touching n.ref to avoid
	// triggering the race detector and looking like a data race.
	c.count = n.count
	c.maxKey = n.maxKey
	c.maxInc = n.maxInc
	c.items = n.items
	if !c.leaf() {
		// Copy children and increase each refcount.
		*c.children = *n.children
		for i := int16(0); i <= c.count; i++ {
			c.children[i].incRef()
		}
	}
	return c
}

func (n *node) insertAt(index int, item T, nd *node) {
	if index < int(n.count) {
		copy(n.items[index+1:n.count+1], n.items[index:n.count])
		if !n.leaf() {
			copy(n.children[index+2:n.count+2], n.children[index+1:n.count+1])
		}
	}
	n.items[index] = item
	if !n.leaf() {
		n.children[index+1] = nd
	}
	n.count++
}

func (n *node) pushBack(item T, nd *node) {
	n.items[n.count] = item
	if !n.leaf() {
		n.children[n.count+1] = nd
	}
	n.count++
}

func (n *node) pushFront(item T, nd *node) {
	if !n.leaf() {
		copy(n.children[1:n.count+2], n.children[:n.count+1])
		n.children[0] = nd
	}
	copy(n.items[1:n.count+1], n.items[:n.count])
	n.items[0] = item
	n.count++
}

// removeAt removes a value at a given index, pulling all subsequent values
// back.
func (n *node) removeAt(index int) (T, *node) {
	var child *node
	if !n.leaf() {
		child = n.children[index+1]
		copy(n.children[index+1:n.count], n.children[index+2:n.count+1])
		n.children[n.count] = nil
	}
	n.count--
	out := n.items[index]
	copy(n.items[index:n.count], n.items[index+1:n.count+1])
	n.items[n.count] = nilT
	return out, child
}

// popBack removes and returns the last element in the list.
func (n *node) popBack() (T, *node) {
	n.count--
	out := n.items[n.count]
	n.items[n.count] = nilT
	if n.leaf() {
		return out, nil
	}
	child := n.children[n.count+1]
	n.children[n.count+1] = nil
	return out, child
}

// popFront removes and returns the first element in the list.
func (n *node) popFront() (T, *node) {
	n.count--
	var child *node
	if !n.leaf() {
		child = n.children[0]
		copy(n.children[:n.count+1], n.children[1:n.count+2])
		n.children[n.count+1] = nil
	}
	out := n.items[0]
	copy(n.items[:n.count], n.items[1:n.count+1])
	n.items[n.count] = nilT
	return out, child
}

// find returns the index where the given item should be inserted into this
// list. 'found' is true if the item already exists in the list at the given
// index.
func (n *node) find(item T) (index int, found bool) {
	// Logic copied from sort.Search. Inlining this gave
	// an 11% speedup on BenchmarkBTreeDeleteInsert.
	i, j := 0, int(n.count)
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		v := compare(item, n.items[h])
		if v == 0 {
			return h, true
		} else if v > 0 {
			i = h + 1
		} else {
			j = h
		}
	}
	return i, false
}

// split splits the given node at the given index. The current node shrinks,
// and this function returns the item that existed at that index and a new
// node containing all items/children after it.
//
// Before:
//
//	+-----------+
//	|   x y z   |
//	+--/-/-\-\--+
//
// After:
//
//	+-----------+
//	|     y     |
//	+----/-\----+
//	    /   \
//	   v     v
//
// +-----------+     +-----------+
// |         x |     | z         |
// +-----------+     +-----------+
func (n *node) split(i int) (T, *node) {
	out := n.items[i]
	var next *node
	if n.leaf() {
		next = newLeafNode()
	} else {
		next = newNode()
	}
	next.count = n.count - int16(i+1)
	copy(next.items[:], n.items[i+1:n.count])
	for j := int16(i); j < n.count; j++ {
		n.items[j] = nilT
	}
	if !n.leaf() {
		copy(next.children[:], n.children[i+1:n.count+1])
		for j := int16(i + 1); j <= n.count; j++ {
			n.children[j] = nil
		}
	}
	n.count = int16(i)

	nextMax := next.findUpperBound()
	next.setMax(nextMax)
	nMax := n.max()
	if nMax.compare(nextMax) != 0 && nMax.compare(upperBound(out)) != 0 {
		// If upper bound wasn't from new node or item
		// at index i, it must still be from old node.
	} else {
		n.setMax(n.findUpperBound())
	}
	return out, next
}

// insert inserts an item into the subtree rooted at this node, making sure no
// nodes in the subtree exceed maxItems items. Returns true if an existing item
// was replaced and false if an item was inserted. Also returns whether the
// node's upper bound changes.
func (n *node) insert(item T) (replaced, newBound bool) {
	i, found := n.find(item)
	if found {
		n.items[i] = item
		return true, false
	}
	if n.leaf() {
		n.insertAt(i, item, nil)
		return false, n.adjustUpperBoundOnInsertion(item, nil)
	}
	if n.children[i].count >= maxItems {
		splitLa, splitNode := mut(&n.children[i]).split(maxItems / 2)
		n.insertAt(i, splitLa, splitNode)

		switch v := compare(item, n.items[i]); {
		case v < 0:
			// no change, we want first split node
		case v > 0:
			i++ // we want second split node
		default:
			n.items[i] = item
			return true, false
		}
	}
	replaced, newBound = mut(&n.children[i]).insert(item)
	if newBound {
		newBound = n.adjustUpperBoundOnInsertion(item, nil)
	}
	return replaced, newBound
}

// removeMax removes and returns the maximum item from the subtree rooted at
// this node.
func (n *node) removeMax() T {
	if n.leaf() {
		n.count--
		out := n.items[n.count]
		n.items[n.count] = nilT
		n.adjustUpperBoundOnRemoval(out, nil)
		return out
	}
	// Recurse into max child.
	i := int(n.count)
	if n.children[i].count <= minItems {
		// Child not large enough to remove from.
		n.rebalanceOrMerge(i)
		return n.removeMax() // redo
	}
	child := mut(&n.children[i])
	out := child.removeMax()
	n.adjustUpperBoundOnRemoval(out, nil)
	return out
}

// remove removes an item from the subtree rooted at this node. Returns the item
// that was removed or nil if no matching item was found. Also returns whether
// the node's upper bound changes.
func (n *node) remove(item T) (out T, newBound bool) {
	i, found := n.find(item)
	if n.leaf() {
		if found {
			out, _ = n.removeAt(i)
			return out, n.adjustUpperBoundOnRemoval(out, nil)
		}
		return nilT, false
	}
	if n.children[i].count <= minItems {
		// Child not large enough to remove from.
		n.rebalanceOrMerge(i)
		return n.remove(item) // redo
	}
	child := mut(&n.children[i])
	if found {
		// Replace the item being removed with the max item in our left child.
		out = n.items[i]
		n.items[i] = child.removeMax()
		return out, n.adjustUpperBoundOnRemoval(out, nil)
	}
	// Latch is not in this node and child is large enough to remove from.
	out, newBound = child.remove(item)
	if newBound {
		newBound = n.adjustUpperBoundOnRemoval(out, nil)
	}
	return out, newBound
}

// rebalanceOrMerge grows child 'i' to ensure it has sufficient room to remove
// an item from it while keeping it at or above minItems.
func (n *node) rebalanceOrMerge(i int) {
	switch {
	case i > 0 && n.children[i-1].count > minItems:
		// Rebalance from left sibling.
		//
		//          +-----------+
		//          |     y     |
		//          +----/-\----+
		//              /   \
		//             v     v
		// +-----------+     +-----------+
		// |         x |     |           |
		// +----------\+     +-----------+
		//             \
		//              v
		//              a
		//
		// After:
		//
		//          +-----------+
		//          |     x     |
		//          +----/-\----+
		//              /   \
		//             v     v
		// +-----------+     +-----------+
		// |           |     | y         |
		// +-----------+     +/----------+
		//                   /
		//                  v
		//                  a
		//
		left := mut(&n.children[i-1])
		child := mut(&n.children[i])
		xLa, grandChild := left.popBack()
		yLa := n.items[i-1]
		child.pushFront(yLa, grandChild)
		n.items[i-1] = xLa

		left.adjustUpperBoundOnRemoval(xLa, grandChild)
		child.adjustUpperBoundOnInsertion(yLa, grandChild)

	case i < int(n.count) && n.children[i+1].count > minItems:
		// Rebalance from right sibling.
		//
		//          +-----------+
		//          |     y     |
		//          +----/-\----+
		//              /   \
		//             v     v
		// +-----------+     +-----------+
		// |           |     | x         |
		// +-----------+     +/----------+
		//                   /
		//                  v
		//                  a
		//
		// After:
		//
		//          +-----------+
		//          |     x     |
		//          +----/-\----+
		//              /   \
		//             v     v
		// +-----------+     +-----------+
		// |         y |     |           |
		// +----------\+     +-----------+
		//             \
		//              v
		//              a
		//
		right := mut(&n.children[i+1])
		child := mut(&n.children[i])
		xLa, grandChild := right.popFront()
		yLa := n.items[i]
		child.pushBack(yLa, grandChild)
		n.items[i] = xLa

		right.adjustUpperBoundOnRemoval(xLa, grandChild)
		child.adjustUpperBoundOnInsertion(yLa, grandChild)

	default:
		// Merge with either the left or right sibling.
		//
		//          +-----------+
		//          |   u y v   |
		//          +----/-\----+
		//              /   \
		//             v     v
		// +-----------+     +-----------+
		// |         x |     | z         |
		// +-----------+     +-----------+
		//
		// After:
		//
		//          +-----------+
		//          |    u v    |
		//          +-----|-----+
		//                |
		//                v
		//          +-----------+
		//          |   x y z   |
		//          +-----------+
		//
		if i >= int(n.count) {
			i = int(n.count - 1)
		}
		child := mut(&n.children[i])
		// Make mergeChild mutable, bumping the refcounts on its children if necessary.
		_ = mut(&n.children[i+1])
		mergeLa, mergeChild := n.removeAt(i)
		child.items[child.count] = mergeLa
		copy(child.items[child.count+1:], mergeChild.items[:mergeChild.count])
		if !child.leaf() {
			copy(child.children[child.count+1:], mergeChild.children[:mergeChild.count+1])
		}
		child.count += mergeChild.count + 1

		child.adjustUpperBoundOnInsertion(mergeLa, mergeChild)
		mergeChild.decRef(false /* recursive */)
	}
}

// findUpperBound returns the largest end key node range, assuming that its
// children have correct upper bounds already set.
func (n *node) findUpperBound() keyBound {
	var max keyBound
	for i := int16(0); i < n.count; i++ {
		up := upperBound(n.items[i])
		if max.compare(up) < 0 {
			max = up
		}
	}
	if !n.leaf() {
		for i := int16(0); i <= n.count; i++ {
			up := n.children[i].max()
			if max.compare(up) < 0 {
				max = up
			}
		}
	}
	return max
}

// adjustUpperBoundOnInsertion adjusts the upper key bound for this node given
// an item and an optional child node that was inserted. Returns true is the
// upper bound was changed and false if not.
func (n *node) adjustUpperBoundOnInsertion(item T, child *node) bool {
	up := upperBound(item)
	if child != nil {
		if childMax := child.max(); up.compare(childMax) < 0 {
			up = childMax
		}
	}
	if n.max().compare(up) < 0 {
		n.setMax(up)
		return true
	}
	return false
}

// adjustUpperBoundOnRemoval adjusts the upper key bound for this node given an
// item and an optional child node that was removed. Returns true is the upper
// bound was changed and false if not.
func (n *node) adjustUpperBoundOnRemoval(item T, child *node) bool {
	up := upperBound(item)
	if child != nil {
		if childMax := child.max(); up.compare(childMax) < 0 {
			up = childMax
		}
	}
	if n.max().compare(up) == 0 {
		// up was previous upper bound of n.
		max := n.findUpperBound()
		n.setMax(max)
		return max.compare(up) != 0
	}
	return false
}

// btree is an implementation of an augmented interval B-Tree.
//
// btree stores items in an ordered structure, allowing easy insertion,
// removal, and iteration. It represents intervals and permits an interval
// search operation following the approach laid out in CLRS, Chapter 14.
// The B-Tree stores items in order based on their start key and each
// B-Tree node maintains the upper-bound end key of all items in its
// subtree.
//
// Write operations are not safe for concurrent mutation by multiple
// goroutines, but Read operations are.
type btree struct {
	root   *node
	length int
}

// Reset removes all items from the btree. In doing so, it allows memory
// held by the btree to be recycled. Failure to call this method before
// letting a btree be GCed is safe in that it won't cause a memory leak,
// but it will prevent btree nodes from being efficiently re-used.
func (t *btree) Reset() {
	if t.root != nil {
		t.root.decRef(true /* recursive */)
		t.root = nil
	}
	t.length = 0
}

// Clone clones the btree, lazily. It does so in constant time.
func (t *btree) Clone() btree {
	c := *t
	if c.root != nil {
		// Incrementing the reference count on the root node is sufficient to
		// ensure that no node in the cloned tree can be mutated by an actor
		// holding a reference to the original tree and vice versa. This
		// property is upheld because the root node in the receiver btree and
		// the returned btree will both necessarily have a reference count of at
		// least 2 when this method returns. All tree mutations recursively
		// acquire mutable node references (see mut) as they traverse down the
		// tree. The act of acquiring a mutable node reference performs a clone
		// if a node's reference count is greater than one. Cloning a node (see
		// clone) increases the reference count on each of its children,
		// ensuring that they have a reference count of at least 2. This, in
		// turn, ensures that any of the child nodes that are modified will also
		// be copied-on-write, recursively ensuring the immutability property
		// over the entire tree.
		c.root.incRef()
	}
	return c
}

// Delete removes an item equal to the passed in item from the tree.
func (t *btree) Delete(item T) {
	if t.root == nil || t.root.count == 0 {
		return
	}
	if out, _ := mut(&t.root).remove(item); out != nilT {
		t.length--
	}
	if t.root.count == 0 {
		old := t.root
		if t.root.leaf() {
			t.root = nil
		} else {
			t.root = t.root.children[0]
		}
		old.decRef(false /* recursive */)
	}
}

// Set adds the given item to the tree. If an item in the tree already equals
// the given one, it is replaced with the new item.
func (t *btree) Set(item T) {
	if t.root == nil {
		t.root = newLeafNode()
	} else if t.root.count >= maxItems {
		splitLa, splitNode := mut(&t.root).split(maxItems / 2)
		newRoot := newNode()
		newRoot.count = 1
		newRoot.items[0] = splitLa
		newRoot.children[0] = t.root
		newRoot.children[1] = splitNode
		newRoot.setMax(newRoot.findUpperBound())
		t.root = newRoot
	}
	if replaced, _ := mut(&t.root).insert(item); !replaced {
		t.length++
	}
}

// MakeIter returns a new iterator object. It is not safe to continue using an
// iterator after modifications are made to the tree. If modifications are made,
// create a new iterator.
func (t *btree) MakeIter() iterator {
	return iterator{r: t.root, pos: -1}
}

// Height returns the height of the tree.
func (t *btree) Height() int {
	if t.root == nil {
		return 0
	}
	h := 1
	n := t.root
	for !n.leaf() {
		n = n.children[0]
		h++
	}
	return h
}

// Len returns the number of items currently in the tree.
func (t *btree) Len() int {
	return t.length
}

// String returns a string description of the tree. The format is
// similar to the https://en.wikipedia.org/wiki/Newick_format.
func (t *btree) String() string {
	if t.length == 0 {
		return ";"
	}
	var b strings.Builder
	t.root.writeString(&b)
	return b.String()
}

func (n *node) writeString(b *strings.Builder) {
	if n.leaf() {
		for i := int16(0); i < n.count; i++ {
			if i != 0 {
				b.WriteString(",")
			}
			b.WriteString(n.items[i].String())
		}
		return
	}
	for i := int16(0); i <= n.count; i++ {
		b.WriteString("(")
		n.children[i].writeString(b)
		b.WriteString(")")
		if i < n.count {
			b.WriteString(n.items[i].String())
		}
	}
}

// iterStack represents a stack of (node, pos) tuples, which captures
// iteration state as an iterator descends a btree.
type iterStack struct {
	a    iterStackArr
	aLen int16 // -1 when using s
	s    []iterFrame
}

// Used to avoid allocations for stacks below a certain size.
type iterStackArr [3]iterFrame

type iterFrame struct {
	n   *node
	pos int16
}

func (is *iterStack) push(f iterFrame) {
	if is.aLen == -1 {
		is.s = append(is.s, f)
	} else if int(is.aLen) == len(is.a) {
		is.s = make([]iterFrame, int(is.aLen)+1, 2*int(is.aLen))
		copy(is.s, is.a[:])
		is.s[int(is.aLen)] = f
		is.aLen = -1
	} else {
		is.a[is.aLen] = f
		is.aLen++
	}
}

func (is *iterStack) pop() iterFrame {
	if is.aLen == -1 {
		f := is.s[len(is.s)-1]
		is.s = is.s[:len(is.s)-1]
		return f
	}
	is.aLen--
	return is.a[is.aLen]
}

func (is *iterStack) len() int {
	if is.aLen == -1 {
		return len(is.s)
	}
	return int(is.aLen)
}

func (is *iterStack) reset() {
	if is.aLen == -1 {
		is.s = is.s[:0]
	} else {
		is.aLen = 0
	}
}

// iterator is responsible for search and traversal within a btree.
type iterator struct {
	r   *node
	n   *node
	pos int16
	s   iterStack
	o   overlapScan
}

func (i *iterator) reset() {
	i.n = i.r
	i.pos = -1
	i.s.reset()
	i.o = overlapScan{}
}

func (i *iterator) descend(n *node, pos int16) {
	i.s.push(iterFrame{n: n, pos: pos})
	i.n = n.children[pos]
	i.pos = 0
}

// ascend ascends up to the current node's parent and resets the position
// to the one previously set for this parent node.
func (i *iterator) ascend() {
	f := i.s.pop()
	i.n = f.n
	i.pos = f.pos
}

// SeekGE seeks to the first item greater-than or equal to the provided
// item.
func (i *iterator) SeekGE(item T) {
	i.reset()
	if i.n == nil {
		return
	}
	for {
		pos, found := i.n.find(item)
		i.pos = int16(pos)
		if found {
			return
		}
		if i.n.leaf() {
			if i.pos == i.n.count {
				i.Next()
			}
			return
		}
		i.descend(i.n, i.pos)
	}
}

// SeekLT seeks to the first item less-than the provided item.
func (i *iterator) SeekLT(item T) {
	i.reset()
	if i.n == nil {
		return
	}
	for {
		pos, found := i.n.find(item)
		i.pos = int16(pos)
		if found || i.n.leaf() {
			i.Prev()
			return
		}
		i.descend(i.n, i.pos)
	}
}

// First seeks to the first item in the btree.
func (i *iterator) First() {
	i.reset()
	if i.n == nil {
		return
	}
	for !i.n.leaf() {
		i.descend(i.n, 0)
	}
	i.pos = 0
}

// Last seeks to the last item in the btree.
func (i *iterator) Last() {
	i.reset()
	if i.n == nil {
		return
	}
	for !i.n.leaf() {
		i.descend(i.n, i.n.count)
	}
	i.pos = i.n.count - 1
}

// Next positions the iterator to the item immediately following
// its current position.
func (i *iterator) Next() {
	if i.n == nil {
		return
	}

	if i.n.leaf() {
		i.pos++
		if i.pos < i.n.count {
			return
		}
		for i.s.len() > 0 && i.pos >= i.n.count {
			i.ascend()
		}
		return
	}

	i.descend(i.n, i.pos+1)
	for !i.n.leaf() {
		i.descend(i.n, 0)
	}
	i.pos = 0
}

// Prev positions the iterator to the item immediately preceding
// its current position.
func (i *iterator) Prev() {
	if i.n == nil {
		return
	}

	if i.n.leaf() {
		i.pos--
		if i.pos >= 0 {
			return
		}
		for i.s.len() > 0 && i.pos < 0 {
			i.ascend()
			i.pos--
		}
		return
	}

	i.descend(i.n, i.pos)
	for !i.n.leaf() {
		i.descend(i.n, i.n.count)
	}
	i.pos = i.n.count - 1
}

// Valid returns whether the iterator is positioned at a valid position.
func (i *iterator) Valid() bool {
	return i.pos >= 0 && i.pos < i.n.count
}

// Cur returns the item at the iterator's current position. It is illegal
// to call Cur if the iterator is not valid.
func (i *iterator) Cur() T {
	return i.n.items[i.pos]
}

// An overlap scan is a scan over all items that overlap with the provided item
// (start key inclusive, end key exclusive) in order of the overlapping items'
// start keys. The goal of the scan is to minimize the number of key comparisons
// performed in total. The algorithm operates based on the following two
// invariants maintained by augmented interval btree:
//  1. all items are sorted in the btree based on their start key.
//  2. all btree nodes maintain the upper bound end key of all items
//     in their subtree.
//
// An overlapping scan can be performed either in the forward or reverse
// direction. The algorithm for each is slightly different. First, we present
// each algorithm, and then we talk about the differences (and the reasons for
// for them).
//
// -----------------------------------------------------------------------------
// Forward Overlap Scan Algorithm
// -----------------------------------------------------------------------------
//
// The scan algorithm starts in "unconstrained minimum" and "unconstrained
// maximum" states. To enter a "constrained minimum" state, the scan must reach
// items in the tree with start keys above the search range's start key.
// Because items in the tree are sorted by start key, once the scan enters the
// "constrained minimum" state it will remain there. To enter a "constrained
// maximum" state, the scan must determine the first child btree node in a given
// subtree that can have items with start keys above the search range's end
// key. The scan then remains in the "constrained maximum" state until it
// traverse into this child node, at which point it moves to the "unconstrained
// maximum" state again.
//
// The scan algorithm works like a standard btree forward scan with the
// following augmentations:
//  1. before traversing the tree, the scan performs a binary search on the
//     root node's items to determine a "soft" lower-bound constraint position
//     and a "hard" upper-bound constraint position in the root's children.
//  2. when traversing into a child node in the lower or upper bound constraint
//     position, the constraint is refined by searching the child's items.
//  3. the initial traversal down the tree follows the left-most children
//     whose upper bound end keys are equal to or greater than the start key
//     of the search range. The children followed will be equal to or less
//     than the soft lower bound constraint.
//  4. once the initial traversal completes and the scan is in the left-most
//     btree node whose upper bound overlaps the search range, key comparisons
//     must be performed with each item in the tree. This is necessary because
//     any of these items may have end keys that cause them to overlap with the
//     search range.
//  5. once the scan reaches the lower bound constraint position (the first item
//     with a start key equal to or greater than the search range's start key),
//     it can begin scanning without performing key comparisons. This is allowed
//     because all items from this point forward will have end keys that are
//     greater than the search range's start key.
//  6. once the scan reaches the upper bound constraint position, it terminates.
//     It does so because the item at this position is the first item with a
//     start key larger than the search range's end key.
//
// -----------------------------------------------------------------------------
// Reverse Overlap Scan Algorithm
// -----------------------------------------------------------------------------
//
// The scan algorithm starts in a "unconstrained minimum" and "unconstrained
// maximum" state. To enter the "constrained maximum" state, the scan must reach
// items in the tree with start keys below the search range's end key. Once the
// scan enters the "constrained maximum" state, it will remain there. To enter
// the "constrained minimum" state, the scan must reach items in the tree with
// start keys above the search range's start key. To enter a "constrained
// minimum" state, the scan must determine the first child btree node in a given
// subtree that can have items with start keys less than the search range's
// start key. The scan then remains in the "constrained minimum" state, until it
// traverses into this child node, at which point it moves to the "unconstrained
// minimum" state again.
//
// The scan algorithm works like a standard btree reverse scan with the
// following augmentations:
//  1. before traversing the tree, the scan performs a binary search on the root
//     node's items to determine a "soft" lower-bound constraint position and a
//     "hard" upper-bound constraint position in the root's children.
//  2. when traversing into a child node in the lower or upper bound constraint
//     position the constraint is refined by searching the child's items.
//  3. the initial traversal down the tree follows the right-most children whose
//     upper bound end keys are equal to or greater than the start key of the search
//     range. The children followed will be equal to or less than the hard
//     upper-bound constraint.
//  4. once the initial traversal completes and the scan is in the right-most
//     btree node whose upper bound overlaps the search range, then jumps directly
//     to the first item before the upper bound constraint position. This is because
//     the upper bound constraint position is exclusive.
//  5. as long as the scan hasn't reached the lower bound constraint position
//     (the first item with a start key equal or greater than the search range's
//     start key), it can continue scanning without performing key comparisions.
//     This is allowed because all items until the lower bound constraint position
//     is reached will have end keys greater than the search range's start key, and
//     we know that all items in the tree have start keys less than the search
//     range's end key.
//  6. once the scan reaches the lower bound constraint position, key comparisons
//     must be performed with each item in the tree. This is necessary because even
//     though the scan is dealing with items that have start keys less than the
//     serach range's start key, it is possible that these items have end keys that
//     cause them to overlap with the search range.
//
// -----------------------------------------------------------------------------
// Differences between forward and reverse overlap scans
// -----------------------------------------------------------------------------
//
//  1. The forward overlapping scan can terminate early. It can do so once it
//     reaches the hard upper bound constaint. This is the first item whose start
//     key is greater than or equal to the search range's end key.
//  2. The reverse overlaping scan can directly "jump" directly to the item right
//     before the hard upper bound constraint. This allows it to skip any subtree/
//     items that are past the search range's end key.
//  3. Once the forward overlapping scan enters the inConstrMin condition, it
//     stays there. The scan reaches inConstrMin at the first item whose start key
//     is greater than the search range's start key. While inConstrMin is true, no
//     key comparisons need to be performed[1].
//  4. In contrast, the reverse overlapping starts off[2] in the inConstrMin
//     state (where no key comparisons are needed). Once it exitst this state, key
//     comparisons are needed.
//
// [1] This is because the search interval's end key must be greater than the
// item's start key. Otherwise, we'd have terminated the scan. So, the item must
// overlap with the search range, and we can skip key comparisons.
// [2] After it's jumped to the item right before the "hard" upper bound
// constraint.
//
// There's a few reasons for the differences here:
//  1. All items in the b-tree are sorted by start key. This means that we have a
//     "hard" upper bound, but a "soft" lower bound. For forward scans, this means
//     that we can terminate early when this "hard" upper bound is reached. For
//     reverse scans, this means that we can directly jump to the last item the scan
//     will ever see.
//  2. The number of children at an interior node is one more than the number of
//     items. For a forward scan, the scan starts off at the 0th index regardless of
//     the type of node. For a reverse scan, however, the scan must start off at the
//     last child for interior nodes, and the last item for leaf-nodes. These are "off
//     by one", which requires some care to get right.
type overlapScan struct {
	// The "soft" lower-bound constraint.
	constrMinN   *node
	constrMinPos int16
	inConstrMin  bool

	// The "hard" upper-bound constraint.
	constrMaxN   *node
	constrMaxPos int16
}

// FirstOverlap seeks to the first item in the btree that overlaps with the
// provided search item.
func (i *iterator) FirstOverlap(item T) {
	i.reset()
	if i.n == nil {
		return
	}
	i.pos = 0
	i.constrainMinSearchBounds(item)
	i.constrainMaxSearchBounds(item)
	i.findNextOverlap(item)
}

// LastOverlap seeks to the last item in the btree that overlaps with the
// provided search item.
func (i *iterator) LastOverlap(item T) {
	i.reset()
	if i.n == nil {
		return
	}
	i.pos = i.n.count
	i.o.inConstrMin = true // reverse scans start inConstrMin
	i.constrainMinSearchBounds(item)
	i.constrainMaxSearchBounds(item)
	i.findPrevOverlap(item)
}

// NextOverlap positions the iterator to the item immediately following
// its current position that overlaps with the search item.
func (i *iterator) NextOverlap(item T) {
	if i.n == nil {
		return
	}
	i.pos++
	i.findNextOverlap(item)
}

// PrevOverlap positions the iterator to the item immediately preceding
// its current position that overlaps with the search item.
func (i *iterator) PrevOverlap(item T) {
	if i.n == nil {
		return
	}
	i.findPrevOverlap(item)
}

// constrainMinSearchBounds sets the "soft" lower-bound constraint. This is the
// first item whose start key is greater than or equal to the supplied search
// range's start key.
//
//	| search range:           [-------------)
//	| items:         [----) [----) [----) [----) [----)
//	|                                ^
//	|                                |
//	|                                +---constrMinPos
func (i *iterator) constrainMinSearchBounds(item T) {
	k := item.Key()
	j := sort.Search(int(i.n.count), func(j int) bool {
		return bytes.Compare(k, i.n.items[j].Key()) <= 0
	})
	i.o.constrMinN = i.n
	i.o.constrMinPos = int16(j)
}

// constrainMaxSearchBounds sets the "hard" upper-bound constraint. This is the
// first item whose start key is greater than the supplied search range's end
// key.
//
//	| search range:           [--------------)
//	| items:         [----) [----) [----) [----) [----)
//	|                                               ^
//	|                                               |
//	|                                               +---constrMaxPos
func (i *iterator) constrainMaxSearchBounds(item T) {
	up := upperBound(item)
	j := sort.Search(int(i.n.count), func(j int) bool {
		return !up.contains(i.n.items[j])
	})
	i.o.constrMaxN = i.n
	i.o.constrMaxPos = int16(j)
}

func (i *iterator) findNextOverlap(item T) {
	for {
		if i.pos > i.n.count {
			// Iterate up tree.
			i.ascend()
		} else if !i.n.leaf() {
			// Iterate down tree.
			if i.o.inConstrMin || i.n.children[i.pos].max().contains(item) {
				par := i.n
				pos := i.pos
				i.descend(par, pos)

				// Refine the constraint bounds, if necessary.
				if par == i.o.constrMinN && pos == i.o.constrMinPos {
					i.constrainMinSearchBounds(item)
				}
				if par == i.o.constrMaxN && pos == i.o.constrMaxPos {
					i.constrainMaxSearchBounds(item)
				}
				continue
			}
		}

		// Check search bounds.
		if i.n == i.o.constrMaxN && i.pos == i.o.constrMaxPos {
			// Invalid. Past possible overlaps.
			i.pos = i.n.count
			return
		}
		if i.n == i.o.constrMinN && i.pos == i.o.constrMinPos {
			// The scan reached the soft lower-bound constraint.
			i.o.inConstrMin = true
		}

		// Iterate across node.
		if i.pos < i.n.count {
			// Check for overlapping item.
			if i.o.inConstrMin {
				// Fast-path to avoid span comparison. i.o.inConstrMin
				// tells us that all items have end keys above our search
				// span's start key.
				return
			}
			if upperBound(i.n.items[i.pos]).contains(item) {
				return
			}
		}
		i.pos++
	}
}

func (i *iterator) findPrevOverlap(item T) {
	for {
		// First off, we want to avoid exploring any items that are past the
		// search range's end key.
		if i.n == i.o.constrMaxN && i.pos > i.o.constrMaxPos {
			// The item at i.o.constrMaxPos is the first item with a start key
			// that's greater than the search range's end key. As such, it is
			// the left-most item that does not overlap with the search range --
			// we want to "jump" straight to the preceding position. However, we
			// can't just position the iterator to i.o.constrMaxPos - 1; if this
			// is an interior node, we might need to descend into the child on
			// the right of i.o.constrMaxPos - 1 first. So, we set the position
			// to that of the child node, and let the loop decide whether to
			// descend further or not.
			i.pos = i.o.constrMaxPos
		}

		if i.pos < 0 {
			if i.s.len() > 0 {
				// Iterate up tree, if possible.
				i.ascend()
			} else {
				// We've reached the root, so there's no more ascending to do;
				// the iterator is already invalid, so simply return.
				return
			}
		} else if !i.n.leaf() {
			// Iterate down tree, but only if there's any hope of finding an
			// overlap. In particular, if the max key found in the child subtree
			// is less than the search item's start key, there's no way we'll
			// find an overlap.
			if i.o.inConstrMin || i.n.children[i.pos].max().contains(item) {
				par := i.n
				pos := i.pos
				i.descend(par, pos)
				// Position the iterator to the last child. It's fine if we
				// descended to a leaf node; we'll be positioned to the last
				// item in the leaf node in the next iteration.
				i.pos = i.n.count

				// Refine the constraint bounds, if necessary.
				if par == i.o.constrMinN && pos == i.o.constrMinPos {
					i.constrainMinSearchBounds(item)
				}
				if par == i.o.constrMaxN && pos == i.o.constrMaxPos {
					i.constrainMaxSearchBounds(item)
				}
				continue
			}
		}

		if i.n == i.o.constrMinN && i.pos == i.o.constrMinPos {
			// The item at i.o.constrMinPos is the left-most item whose start
			// key is greater than or equal to the search range's start key. As
			// such, it's the last item we can eschew key comparisons for. The
			// i.pos-- is going to transition us from a constrained minimum to
			// an unconstrained minimum scan; update state to reflect this.
			i.o.inConstrMin = false
		}

		// NB: We decrement the position between descending into a child node
		// and checking an item. This is in contrast to the forward scan, where
		// the increment happens after doing both.
		//
		// This is because the number of children at an interior node is one
		// more than the number of items. As such, positioning the iterator at
		// the last child is different than positioning it at the last item
		// (it's off by one). On the other hand, positioning the iterator at the
		// first item or the first child is the same index (0).
		i.pos--

		// Iterate across node.
		if i.pos >= 0 {
			if i.o.inConstrMin {
				// Fast-path to avoid span comparison. i.o.inConstrMin
				// tells us that the item has a start key that's greater than
				// the search range's start key.
				return
			}
			if upperBound(i.n.items[i.pos]).contains(item) {
				// We're not in the constrained minimum state, which means the
				// item's start key is less than the search range's start key.
				// We must check if the item's end key results in an overlap
				// (i.e. if the end key is contained in the search range).
				return
			}
		}
	}
}
