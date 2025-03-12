%{
package parser

import (
  "github.com/cockroachdb/cockroach/pkg/sql/scanner"
  "github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
  "github.com/cockroachdb/cockroach/pkg/util/json"
  "github.com/cockroachdb/cockroach/pkg/util/jsonpath"
)

%}

%{

func setErr(jsonpathlex jsonpathLexer, err error) int {
  jsonpathlex.(*lexer).setErr(err)
  return 1
}

func unimplemented(jsonpathlex jsonpathLexer, feature string) int {
  jsonpathlex.(*lexer).Unimplemented(feature)
  return 1
}

var _ scanner.ScanSymType = &jsonpathSymType{}

func (s *jsonpathSymType) jsonpathScanSymType() {}

// ID implements the scanner.ScanSymType interface.
func (s *jsonpathSymType) ID() int32 {
  return s.id
}

// SetID implements the scanner.ScanSymType interface.
func (s *jsonpathSymType) SetID(id int32) {
  s.id = id
}

// Pos implements the scanner.ScanSymType interface.
func (s *jsonpathSymType) Pos() int32 {
  return s.pos
}

// SetPos implements the scanner.ScanSymType interface.
func (s *jsonpathSymType) SetPos(pos int32) {
  s.pos = pos
}

// Str implements the scanner.ScanSymType interface.
func (s *jsonpathSymType) Str() string {
  return s.str
}

// SetStr implements the scanner.ScanSymType interface.
func (s *jsonpathSymType) SetStr(str string) {
  s.str = str
}

// UnionVal implements the scanner.ScanSymType interface.
func (s *jsonpathSymType) UnionVal() interface{} {
  return s.union.val
}

// SetUnionVal implements the scanner.ScanSymType interface.
func (s *jsonpathSymType) SetUnionVal(val interface{}) {
  s.union.val = val
}

type jsonpathSymUnion struct {
  val interface{}
}

func (u *jsonpathSymUnion) jsonpath() jsonpath.Jsonpath {
  return u.val.(jsonpath.Jsonpath)
}

func (u *jsonpathSymUnion) path() jsonpath.Path {
  return u.val.(jsonpath.Path)
}

func (u *jsonpathSymUnion) paths() jsonpath.Paths {
  return u.val.(jsonpath.Paths)
}

func (u *jsonpathSymUnion) pathArr() []jsonpath.Path {
  return u.val.([]jsonpath.Path)
}

func (u *jsonpathSymUnion) bool() bool {
  return u.val.(bool)
}

func (u *jsonpathSymUnion) numVal() *tree.NumVal {
  return u.val.(*tree.NumVal)
}

func (u *jsonpathSymUnion) arrayList() jsonpath.ArrayList {
  return u.val.(jsonpath.ArrayList)
}

%}

%union{
  id      int32
  pos     int32
  str     string
  union   jsonpathSymUnion
}

/*
 * Basic non-keyword token types.  These are hard-wired into the core lexer.
 * They must be listed first so that their numeric codes do not depend on
 * the set of keywords.  Keep this list in sync with backend/parser/gram.y!
 *
 * Some of these are not directly referenced in this file, but they must be
 * here anyway.
 */
%token <str> IDENT UIDENT FCONST SCONST USCONST BCONST XCONST Op
%token <*tree.NumVal> ICONST PARAM
%token <str> TYPECAST DOT_DOT COLON_EQUALS EQUALS_GREATER
%token <str> LESS_EQUALS GREATER_EQUALS NOT_EQUALS
%token <str> ERROR

%token <str> STRICT
%token <str> LAX

%token <str> VARIABLE
%token <str> TO

%type <jsonpath.Jsonpath> jsonpath
%type <jsonpath.Path> expr_or_predicate
%type <jsonpath.Path> expr
%type <jsonpath.Path> accessor_op
%type <jsonpath.Path> path_primary
%type <jsonpath.Path> key
%type <jsonpath.Path> array_accessor
%type <jsonpath.Path> scalar_value
%type <jsonpath.Path> index_elem
%type <[]jsonpath.Path> accessor_expr
%type <[]jsonpath.Path> index_list
%type <str> key_name
%type <str> any_identifier
%type <str> unreserved_keyword
%type <bool> mode

%%

jsonpath:
  mode expr_or_predicate
  {
    jp := jsonpath.Jsonpath{Strict: $1.bool(), Path: $2.path()}
    jsonpathlex.(*lexer).SetJsonpath(jp)
  }
;

mode:
  STRICT
  {
    $$.val = true
  }
| LAX
  {
    $$.val = false
  }
| /* empty */
  {
    $$.val = false
  }
;

expr_or_predicate:
  expr
  {
    $$.val = $1.path()
  }
;

expr:
  accessor_expr
  {
    $$.val = jsonpath.Paths($1.pathArr())
  }
;

accessor_expr:
  path_primary
  {
    $$.val = []jsonpath.Path{$1.path()}
  }
| accessor_expr accessor_op
  {
    $$.val = append($1.pathArr(), $2.path())
  }
;

path_primary:
  '$'
  {
    $$.val = jsonpath.Root{}
  }
| scalar_value
  {
    $$.val = $1.path()
  }
// TODO(normanchenn): support LAST for array ranges.
;

accessor_op:
  '.' key
  {
    $$.val = $2.path()
  }
| array_accessor
  {
    $$.val = $1.path()
  }
;

key:
  key_name
  {
    $$.val = jsonpath.Key($1)
  }
;

key_name:
  any_identifier
  {
    $$ = $1
  }
;

array_accessor:
  '[' '*' ']'
  {
    $$.val = jsonpath.Wildcard{}
  }
| '[' index_list ']'
  {
    $$.val = $2.path()
  }
;

index_list:
  index_elem
  {
    $$.val = jsonpath.ArrayList{$1.path()}
  }
| index_list ',' index_elem
  {
    $$.val = append($1.arrayList(), $3.path())
  }
;

index_elem:
  expr
  {
    $$.val = $1.path()
  }
| expr TO expr
  {
    $$.val = jsonpath.ArrayIndexRange{
      Start: $1.path(),
      End: $3.path(),
    }
  }
;

scalar_value:
  VARIABLE
  {
    $$.val = jsonpath.Scalar{Type: jsonpath.ScalarVariable, Variable: $1}
  }
| ICONST
  {
    i, err := $1.numVal().AsInt64()
    if err != nil {
      return setErr(jsonpathlex, err)
    }
    $$.val = jsonpath.Scalar{Type: jsonpath.ScalarInt, Value: json.FromInt64(i)}
  }
| FCONST
  {
    return unimplemented(jsonpathlex, "float consts")
  }
// TODO(normanchenn): support strings, bools, null.
;

any_identifier:
  IDENT
| unreserved_keyword
;

unreserved_keyword:
  LAX
| STRICT
| TO
;

%%
