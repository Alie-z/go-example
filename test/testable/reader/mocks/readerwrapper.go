// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jianghushinian/blog-go-example/test/testable/reader (interfaces: ReaderWrapper)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockReaderWrapper is a mock of ReaderWrapper interface.
type MockReaderWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockReaderWrapperMockRecorder
}

// MockReaderWrapperMockRecorder is the mock recorder for MockReaderWrapper.
type MockReaderWrapperMockRecorder struct {
	mock *MockReaderWrapper
}

// NewMockReaderWrapper creates a new mock instance.
func NewMockReaderWrapper(ctrl *gomock.Controller) *MockReaderWrapper {
	mock := &MockReaderWrapper{ctrl: ctrl}
	mock.recorder = &MockReaderWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReaderWrapper) EXPECT() *MockReaderWrapperMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockReaderWrapper) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockReaderWrapperMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReaderWrapper)(nil).Read), arg0)
}