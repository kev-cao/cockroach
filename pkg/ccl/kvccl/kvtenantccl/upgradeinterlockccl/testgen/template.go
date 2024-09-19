// Copyright 2023 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package main

var test_template = `// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

// Code generated by pkg/ccl/kvccl/kvtenantccl/upgradeinterlockccl/testgen, DO NOT EDIT.

package upgradeinterlockccl

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/ccl/kvccl/kvtenantccl/upgradeinterlockccl/sharedtestutil"
	"github.com/cockroachdb/cockroach/pkg/testutils/skip"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/log"
)
{{- $variants := .Variants -}}
{{- $tests := .Tests -}}
{{- range $testName, $testConfig := $tests -}}
{{- range $variantKey, $variantValue := $variants}}

func TestTenantUpgradeInterlock_{{$variantValue}}_{{$testName}}(t *testing.T) {
	defer leaktest.AfterTest(t)()
	// Times out under race.
	skip.UnderRace(t)
	// Test target takes 100s+ to run.
	skip.UnderShort(t)
	defer log.Scope(t).Close(t)

	runTest(t, {{$variantKey}}, sharedtestutil.Tests["{{$testName}}"])
}
{{- end -}}{{- end -}}
{{ println}}`
