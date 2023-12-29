// Code generated by MockGen. DO NOT EDIT.
// Source: default_ddd/app/internal/adapters/port (interfaces: CreateUserEndpoint,CreateProductEndpoint,CreateCartItemEndpoint,CreateOrderEndpoint)

// Package mocks is a generated GoMock package.
package mocks

import (
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCreateUserEndpoint is a mock of CreateUserEndpoint interface
type MockCreateUserEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockCreateUserEndpointMockRecorder
}

// MockCreateUserEndpointMockRecorder is the mock recorder for MockCreateUserEndpoint
type MockCreateUserEndpointMockRecorder struct {
	mock *MockCreateUserEndpoint
}

// NewMockCreateUserEndpoint creates a new mock instance
func NewMockCreateUserEndpoint(ctrl *gomock.Controller) *MockCreateUserEndpoint {
	mock := &MockCreateUserEndpoint{ctrl: ctrl}
	mock.recorder = &MockCreateUserEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCreateUserEndpoint) EXPECT() *MockCreateUserEndpointMockRecorder {
	return m.recorder
}

// ExecuteCreateUserEndpoint mocks base method
func (m *MockCreateUserEndpoint) ExecuteCreateUserEndpoint(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExecuteCreateUserEndpoint", arg0)
}

// ExecuteCreateUserEndpoint indicates an expected call of ExecuteCreateUserEndpoint
func (mr *MockCreateUserEndpointMockRecorder) ExecuteCreateUserEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteCreateUserEndpoint", reflect.TypeOf((*MockCreateUserEndpoint)(nil).ExecuteCreateUserEndpoint), arg0)
}

// MockCreateProductEndpoint is a mock of CreateProductEndpoint interface
type MockCreateProductEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockCreateProductEndpointMockRecorder
}

// MockCreateProductEndpointMockRecorder is the mock recorder for MockCreateProductEndpoint
type MockCreateProductEndpointMockRecorder struct {
	mock *MockCreateProductEndpoint
}

// NewMockCreateProductEndpoint creates a new mock instance
func NewMockCreateProductEndpoint(ctrl *gomock.Controller) *MockCreateProductEndpoint {
	mock := &MockCreateProductEndpoint{ctrl: ctrl}
	mock.recorder = &MockCreateProductEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCreateProductEndpoint) EXPECT() *MockCreateProductEndpointMockRecorder {
	return m.recorder
}

// ExecuteCreateProductEndpoint mocks base method
func (m *MockCreateProductEndpoint) ExecuteCreateProductEndpoint(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExecuteCreateProductEndpoint", arg0)
}

// ExecuteCreateProductEndpoint indicates an expected call of ExecuteCreateProductEndpoint
func (mr *MockCreateProductEndpointMockRecorder) ExecuteCreateProductEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteCreateProductEndpoint", reflect.TypeOf((*MockCreateProductEndpoint)(nil).ExecuteCreateProductEndpoint), arg0)
}

// MockCreateCartItemEndpoint is a mock of CreateCartItemEndpoint interface
type MockCreateCartItemEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockCreateCartItemEndpointMockRecorder
}

// MockCreateCartItemEndpointMockRecorder is the mock recorder for MockCreateCartItemEndpoint
type MockCreateCartItemEndpointMockRecorder struct {
	mock *MockCreateCartItemEndpoint
}

// NewMockCreateCartItemEndpoint creates a new mock instance
func NewMockCreateCartItemEndpoint(ctrl *gomock.Controller) *MockCreateCartItemEndpoint {
	mock := &MockCreateCartItemEndpoint{ctrl: ctrl}
	mock.recorder = &MockCreateCartItemEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCreateCartItemEndpoint) EXPECT() *MockCreateCartItemEndpointMockRecorder {
	return m.recorder
}

// ExecuteCreateCartItemEndpoint mocks base method
func (m *MockCreateCartItemEndpoint) ExecuteCreateCartItemEndpoint(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExecuteCreateCartItemEndpoint", arg0)
}

// ExecuteCreateCartItemEndpoint indicates an expected call of ExecuteCreateCartItemEndpoint
func (mr *MockCreateCartItemEndpointMockRecorder) ExecuteCreateCartItemEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteCreateCartItemEndpoint", reflect.TypeOf((*MockCreateCartItemEndpoint)(nil).ExecuteCreateCartItemEndpoint), arg0)
}

// MockCreateOrderEndpoint is a mock of CreateOrderEndpoint interface
type MockCreateOrderEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockCreateOrderEndpointMockRecorder
}

// MockCreateOrderEndpointMockRecorder is the mock recorder for MockCreateOrderEndpoint
type MockCreateOrderEndpointMockRecorder struct {
	mock *MockCreateOrderEndpoint
}

// NewMockCreateOrderEndpoint creates a new mock instance
func NewMockCreateOrderEndpoint(ctrl *gomock.Controller) *MockCreateOrderEndpoint {
	mock := &MockCreateOrderEndpoint{ctrl: ctrl}
	mock.recorder = &MockCreateOrderEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCreateOrderEndpoint) EXPECT() *MockCreateOrderEndpointMockRecorder {
	return m.recorder
}

// ExecuteCreateOrderEndpoint mocks base method
func (m *MockCreateOrderEndpoint) ExecuteCreateOrderEndpoint(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExecuteCreateOrderEndpoint", arg0)
}

// ExecuteCreateOrderEndpoint indicates an expected call of ExecuteCreateOrderEndpoint
func (mr *MockCreateOrderEndpointMockRecorder) ExecuteCreateOrderEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteCreateOrderEndpoint", reflect.TypeOf((*MockCreateOrderEndpoint)(nil).ExecuteCreateOrderEndpoint), arg0)
}