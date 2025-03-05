// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cockroachdb/cockroach/pkg/cmd/roachtest/test (interfaces: Test)

// Package clusterstats is a generated GoMock package.
package clusterstats

import (
	context "context"
	reflect "reflect"

	task "github.com/cockroachdb/cockroach/pkg/cmd/roachtest/roachtestutil/task"
	test "github.com/cockroachdb/cockroach/pkg/cmd/roachtest/test"
	logger "github.com/cockroachdb/cockroach/pkg/roachprod/logger"
	version "github.com/cockroachdb/cockroach/pkg/util/version"
	gomock "github.com/golang/mock/gomock"
)

// MockTest is a mock of Test interface.
type MockTest struct {
	ctrl     *gomock.Controller
	recorder *MockTestMockRecorder
}

// MockTestMockRecorder is the mock recorder for MockTest.
type MockTestMockRecorder struct {
	mock *MockTest
}

// NewMockTest creates a new mock instance.
func NewMockTest(ctrl *gomock.Controller) *MockTest {
	mock := &MockTest{ctrl: ctrl}
	mock.recorder = &MockTestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTest) EXPECT() *MockTestMockRecorder {
	return m.recorder
}

// AddParam mocks base method.
func (m *MockTest) AddParam(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddParam", arg0, arg1)
}

// AddParam indicates an expected call of AddParam.
func (mr *MockTestMockRecorder) AddParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddParam", reflect.TypeOf((*MockTest)(nil).AddParam), arg0, arg1)
}

// ArtifactsDir mocks base method.
func (m *MockTest) ArtifactsDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArtifactsDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// ArtifactsDir indicates an expected call of ArtifactsDir.
func (mr *MockTestMockRecorder) ArtifactsDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArtifactsDir", reflect.TypeOf((*MockTest)(nil).ArtifactsDir))
}

// BuildVersion mocks base method.
func (m *MockTest) BuildVersion() *version.Version {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildVersion")
	ret0, _ := ret[0].(*version.Version)
	return ret0
}

// BuildVersion indicates an expected call of BuildVersion.
func (mr *MockTestMockRecorder) BuildVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildVersion", reflect.TypeOf((*MockTest)(nil).BuildVersion))
}

// Cockroach mocks base method.
func (m *MockTest) Cockroach() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cockroach")
	ret0, _ := ret[0].(string)
	return ret0
}

// Cockroach indicates an expected call of Cockroach.
func (mr *MockTestMockRecorder) Cockroach() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cockroach", reflect.TypeOf((*MockTest)(nil).Cockroach))
}

// DeprecatedWorkload mocks base method.
func (m *MockTest) DeprecatedWorkload() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeprecatedWorkload")
	ret0, _ := ret[0].(string)
	return ret0
}

// DeprecatedWorkload indicates an expected call of DeprecatedWorkload.
func (mr *MockTestMockRecorder) DeprecatedWorkload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeprecatedWorkload", reflect.TypeOf((*MockTest)(nil).DeprecatedWorkload))
}

// Error mocks base method.
func (m *MockTest) Error(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockTestMockRecorder) Error(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockTest)(nil).Error), arg0...)
}

// Errorf mocks base method.
func (m *MockTest) Errorf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockTestMockRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockTest)(nil).Errorf), varargs...)
}

// ExportOpenmetrics mocks base method.
func (m *MockTest) ExportOpenmetrics() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportOpenmetrics")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExportOpenmetrics indicates an expected call of ExportOpenmetrics.
func (mr *MockTestMockRecorder) ExportOpenmetrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportOpenmetrics", reflect.TypeOf((*MockTest)(nil).ExportOpenmetrics))
}

// FailNow mocks base method.
func (m *MockTest) FailNow() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FailNow")
}

// FailNow indicates an expected call of FailNow.
func (mr *MockTestMockRecorder) FailNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailNow", reflect.TypeOf((*MockTest)(nil).FailNow))
}

// Failed mocks base method.
func (m *MockTest) Failed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Failed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Failed indicates an expected call of Failed.
func (mr *MockTestMockRecorder) Failed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Failed", reflect.TypeOf((*MockTest)(nil).Failed))
}

// Fatal mocks base method.
func (m *MockTest) Fatal(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockTestMockRecorder) Fatal(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockTest)(nil).Fatal), arg0...)
}

// Fatalf mocks base method.
func (m *MockTest) Fatalf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MockTestMockRecorder) Fatalf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockTest)(nil).Fatalf), varargs...)
}

// GetRunId mocks base method.
func (m *MockTest) GetRunId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRunId indicates an expected call of GetRunId.
func (mr *MockTestMockRecorder) GetRunId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunId", reflect.TypeOf((*MockTest)(nil).GetRunId))
}

// Go mocks base method.
func (m *MockTest) Go(arg0 task.Func, arg1 ...task.Option) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Go", varargs...)
}

// Go indicates an expected call of Go.
func (mr *MockTestMockRecorder) Go(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Go", reflect.TypeOf((*MockTest)(nil).Go), varargs...)
}

// GoCoverArtifactsDir mocks base method.
func (m *MockTest) GoCoverArtifactsDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GoCoverArtifactsDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// GoCoverArtifactsDir indicates an expected call of GoCoverArtifactsDir.
func (mr *MockTestMockRecorder) GoCoverArtifactsDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoCoverArtifactsDir", reflect.TypeOf((*MockTest)(nil).GoCoverArtifactsDir))
}

// GoWithCancel mocks base method.
func (m *MockTest) GoWithCancel(arg0 task.Func, arg1 ...task.Option) context.CancelFunc {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GoWithCancel", varargs...)
	ret0, _ := ret[0].(context.CancelFunc)
	return ret0
}

// GoWithCancel indicates an expected call of GoWithCancel.
func (mr *MockTestMockRecorder) GoWithCancel(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoWithCancel", reflect.TypeOf((*MockTest)(nil).GoWithCancel), varargs...)
}

// Helper mocks base method.
func (m *MockTest) Helper() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Helper")
}

// Helper indicates an expected call of Helper.
func (mr *MockTestMockRecorder) Helper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Helper", reflect.TypeOf((*MockTest)(nil).Helper))
}

// IsBuildVersion mocks base method.
func (m *MockTest) IsBuildVersion(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBuildVersion", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBuildVersion indicates an expected call of IsBuildVersion.
func (mr *MockTestMockRecorder) IsBuildVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBuildVersion", reflect.TypeOf((*MockTest)(nil).IsBuildVersion), arg0)
}

// IsDebug mocks base method.
func (m *MockTest) IsDebug() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDebug")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDebug indicates an expected call of IsDebug.
func (mr *MockTestMockRecorder) IsDebug() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDebug", reflect.TypeOf((*MockTest)(nil).IsDebug))
}

// L mocks base method.
func (m *MockTest) L() *logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "L")
	ret0, _ := ret[0].(*logger.Logger)
	return ret0
}

// L indicates an expected call of L.
func (mr *MockTestMockRecorder) L() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "L", reflect.TypeOf((*MockTest)(nil).L))
}

// Monitor mocks base method.
func (m *MockTest) Monitor() test.Monitor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Monitor")
	ret0, _ := ret[0].(test.Monitor)
	return ret0
}

// Monitor indicates an expected call of Monitor.
func (mr *MockTestMockRecorder) Monitor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Monitor", reflect.TypeOf((*MockTest)(nil).Monitor))
}

// Name mocks base method.
func (m *MockTest) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockTestMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockTest)(nil).Name))
}

// NewErrorGroup mocks base method.
func (m *MockTest) NewErrorGroup(arg0 ...task.Option) task.ErrorGroup {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewErrorGroup", varargs...)
	ret0, _ := ret[0].(task.ErrorGroup)
	return ret0
}

// NewErrorGroup indicates an expected call of NewErrorGroup.
func (mr *MockTestMockRecorder) NewErrorGroup(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewErrorGroup", reflect.TypeOf((*MockTest)(nil).NewErrorGroup), arg0...)
}

// NewGroup mocks base method.
func (m *MockTest) NewGroup(arg0 ...task.Option) task.Group {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewGroup", varargs...)
	ret0, _ := ret[0].(task.Group)
	return ret0
}

// NewGroup indicates an expected call of NewGroup.
func (mr *MockTestMockRecorder) NewGroup(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGroup", reflect.TypeOf((*MockTest)(nil).NewGroup), arg0...)
}

// Owner mocks base method.
func (m *MockTest) Owner() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Owner")
	ret0, _ := ret[0].(string)
	return ret0
}

// Owner indicates an expected call of Owner.
func (mr *MockTestMockRecorder) Owner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Owner", reflect.TypeOf((*MockTest)(nil).Owner))
}

// PerfArtifactsDir mocks base method.
func (m *MockTest) PerfArtifactsDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerfArtifactsDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// PerfArtifactsDir indicates an expected call of PerfArtifactsDir.
func (mr *MockTestMockRecorder) PerfArtifactsDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerfArtifactsDir", reflect.TypeOf((*MockTest)(nil).PerfArtifactsDir))
}

// Progress mocks base method.
func (m *MockTest) Progress(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Progress", arg0)
}

// Progress indicates an expected call of Progress.
func (mr *MockTestMockRecorder) Progress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Progress", reflect.TypeOf((*MockTest)(nil).Progress), arg0)
}

// RuntimeAssertionsCockroach mocks base method.
func (m *MockTest) RuntimeAssertionsCockroach() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuntimeAssertionsCockroach")
	ret0, _ := ret[0].(string)
	return ret0
}

// RuntimeAssertionsCockroach indicates an expected call of RuntimeAssertionsCockroach.
func (mr *MockTestMockRecorder) RuntimeAssertionsCockroach() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuntimeAssertionsCockroach", reflect.TypeOf((*MockTest)(nil).RuntimeAssertionsCockroach))
}

// Skip mocks base method.
func (m *MockTest) Skip(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Skip", varargs...)
}

// Skip indicates an expected call of Skip.
func (mr *MockTestMockRecorder) Skip(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skip", reflect.TypeOf((*MockTest)(nil).Skip), arg0...)
}

// SkipInit mocks base method.
func (m *MockTest) SkipInit() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SkipInit")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SkipInit indicates an expected call of SkipInit.
func (mr *MockTestMockRecorder) SkipInit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SkipInit", reflect.TypeOf((*MockTest)(nil).SkipInit))
}

// Skipf mocks base method.
func (m *MockTest) Skipf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Skipf", varargs...)
}

// Skipf indicates an expected call of Skipf.
func (mr *MockTestMockRecorder) Skipf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skipf", reflect.TypeOf((*MockTest)(nil).Skipf), varargs...)
}

// SnapshotPrefix mocks base method.
func (m *MockTest) SnapshotPrefix() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotPrefix")
	ret0, _ := ret[0].(string)
	return ret0
}

// SnapshotPrefix indicates an expected call of SnapshotPrefix.
func (mr *MockTestMockRecorder) SnapshotPrefix() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotPrefix", reflect.TypeOf((*MockTest)(nil).SnapshotPrefix))
}

// Spec mocks base method.
func (m *MockTest) Spec() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Spec")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Spec indicates an expected call of Spec.
func (mr *MockTestMockRecorder) Spec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Spec", reflect.TypeOf((*MockTest)(nil).Spec))
}

// StandardCockroach mocks base method.
func (m *MockTest) StandardCockroach() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StandardCockroach")
	ret0, _ := ret[0].(string)
	return ret0
}

// StandardCockroach indicates an expected call of StandardCockroach.
func (mr *MockTestMockRecorder) StandardCockroach() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StandardCockroach", reflect.TypeOf((*MockTest)(nil).StandardCockroach))
}

// Status mocks base method.
func (m *MockTest) Status(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Status", varargs...)
}

// Status indicates an expected call of Status.
func (mr *MockTestMockRecorder) Status(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockTest)(nil).Status), arg0...)
}

// VersionsBinaryOverride mocks base method.
func (m *MockTest) VersionsBinaryOverride() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VersionsBinaryOverride")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// VersionsBinaryOverride indicates an expected call of VersionsBinaryOverride.
func (mr *MockTestMockRecorder) VersionsBinaryOverride() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VersionsBinaryOverride", reflect.TypeOf((*MockTest)(nil).VersionsBinaryOverride))
}

// WorkerProgress mocks base method.
func (m *MockTest) WorkerProgress(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WorkerProgress", arg0)
}

// WorkerProgress indicates an expected call of WorkerProgress.
func (mr *MockTestMockRecorder) WorkerProgress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkerProgress", reflect.TypeOf((*MockTest)(nil).WorkerProgress), arg0)
}

// WorkerStatus mocks base method.
func (m *MockTest) WorkerStatus(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "WorkerStatus", varargs...)
}

// WorkerStatus indicates an expected call of WorkerStatus.
func (mr *MockTestMockRecorder) WorkerStatus(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkerStatus", reflect.TypeOf((*MockTest)(nil).WorkerStatus), arg0...)
}
