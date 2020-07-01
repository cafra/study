// Code generated by MockGen. DO NOT EDIT.
// Source: gomock.go

// Package mock_unit is a generated GoMock package.
package mock_unit

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMyInterface is a mock of MyInterface interface
type MockMyInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMyInterfaceMockRecorder
}

// MockMyInterfaceMockRecorder is the mock recorder for MockMyInterface
type MockMyInterfaceMockRecorder struct {
	mock *MockMyInterface
}

// NewMockMyInterface creates a new mock instance
func NewMockMyInterface(ctrl *gomock.Controller) *MockMyInterface {
	mock := &MockMyInterface{ctrl: ctrl}
	mock.recorder = &MockMyInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMyInterface) EXPECT() *MockMyInterfaceMockRecorder {
	return m.recorder
}

// HelloWorld mocks base method
func (m *MockMyInterface) HelloWorld(x int64, y string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HelloWorld", x, y)
}

// HelloWorld indicates an expected call of HelloWorld
func (mr *MockMyInterfaceMockRecorder) HelloWorld(x, y interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloWorld", reflect.TypeOf((*MockMyInterface)(nil).HelloWorld), x, y)
}