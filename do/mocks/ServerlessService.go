// Code generated by MockGen. DO NOT EDIT.
// Source: serverless.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	exec "os/exec"
	reflect "reflect"

	do "github.com/digitalocean/doctl/do"
	gomock "github.com/golang/mock/gomock"
)

// MockServerlessService is a mock of ServerlessService interface.
type MockServerlessService struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessServiceMockRecorder
}

// MockServerlessServiceMockRecorder is the mock recorder for MockServerlessService.
type MockServerlessServiceMockRecorder struct {
	mock *MockServerlessService
}

// NewMockServerlessService creates a new mock instance.
func NewMockServerlessService(ctrl *gomock.Controller) *MockServerlessService {
	mock := &MockServerlessService{ctrl: ctrl}
	mock.recorder = &MockServerlessServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessService) EXPECT() *MockServerlessServiceMockRecorder {
	return m.recorder
}

// CheckServerlessStatus mocks base method.
func (m *MockServerlessService) CheckServerlessStatus(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckServerlessStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckServerlessStatus indicates an expected call of CheckServerlessStatus.
func (mr *MockServerlessServiceMockRecorder) CheckServerlessStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckServerlessStatus", reflect.TypeOf((*MockServerlessService)(nil).CheckServerlessStatus), arg0)
}

// Cmd mocks base method.
func (m *MockServerlessService) Cmd(arg0 string, arg1 []string) (*exec.Cmd, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cmd", arg0, arg1)
	ret0, _ := ret[0].(*exec.Cmd)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cmd indicates an expected call of Cmd.
func (mr *MockServerlessServiceMockRecorder) Cmd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cmd", reflect.TypeOf((*MockServerlessService)(nil).Cmd), arg0, arg1)
}

// Exec mocks base method.
func (m *MockServerlessService) Exec(arg0 *exec.Cmd) (do.ServerlessOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", arg0)
	ret0, _ := ret[0].(do.ServerlessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockServerlessServiceMockRecorder) Exec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockServerlessService)(nil).Exec), arg0)
}

// GetHostInfo mocks base method.
func (m *MockServerlessService) GetHostInfo(arg0 string) (do.ServerlessHostInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostInfo", arg0)
	ret0, _ := ret[0].(do.ServerlessHostInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostInfo indicates an expected call of GetHostInfo.
func (mr *MockServerlessServiceMockRecorder) GetHostInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostInfo", reflect.TypeOf((*MockServerlessService)(nil).GetHostInfo), arg0)
}

// GetServerlessNamespace mocks base method.
func (m *MockServerlessService) GetServerlessNamespace(arg0 context.Context) (do.ServerlessCredentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerlessNamespace", arg0)
	ret0, _ := ret[0].(do.ServerlessCredentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerlessNamespace indicates an expected call of GetServerlessNamespace.
func (mr *MockServerlessServiceMockRecorder) GetServerlessNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerlessNamespace", reflect.TypeOf((*MockServerlessService)(nil).GetServerlessNamespace), arg0)
}

// InstallServerless mocks base method.
func (m *MockServerlessService) InstallServerless(arg0 string, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallServerless", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallServerless indicates an expected call of InstallServerless.
func (mr *MockServerlessServiceMockRecorder) InstallServerless(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallServerless", reflect.TypeOf((*MockServerlessService)(nil).InstallServerless), arg0, arg1)
}

// Stream mocks base method.
func (m *MockServerlessService) Stream(arg0 *exec.Cmd) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stream indicates an expected call of Stream.
func (mr *MockServerlessServiceMockRecorder) Stream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockServerlessService)(nil).Stream), arg0)
}

// WriteCredentials mocks base method.
func (m *MockServerlessService) WriteCredentials(arg0 do.ServerlessCredentials) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteCredentials", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteCredentials indicates an expected call of WriteCredentials.
func (mr *MockServerlessServiceMockRecorder) WriteCredentials(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteCredentials", reflect.TypeOf((*MockServerlessService)(nil).WriteCredentials), arg0)
}
