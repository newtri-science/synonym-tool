// Code generated by MockGen. DO NOT EDIT.
// Source: utils/casbin_enforcer.go
//
// Generated by this command:
//
//	mockgen -source utils/casbin_enforcer.go -destination mocks/casbin_enforcer.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	echo "github.com/labstack/echo/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockCasbinEnforcer is a mock of CasbinEnforcer interface.
type MockCasbinEnforcer struct {
	ctrl     *gomock.Controller
	recorder *MockCasbinEnforcerMockRecorder
}

// MockCasbinEnforcerMockRecorder is the mock recorder for MockCasbinEnforcer.
type MockCasbinEnforcerMockRecorder struct {
	mock *MockCasbinEnforcer
}

// NewMockCasbinEnforcer creates a new mock instance.
func NewMockCasbinEnforcer(ctrl *gomock.Controller) *MockCasbinEnforcer {
	mock := &MockCasbinEnforcer{ctrl: ctrl}
	mock.recorder = &MockCasbinEnforcerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCasbinEnforcer) EXPECT() *MockCasbinEnforcerMockRecorder {
	return m.recorder
}

// Enforce mocks base method.
func (m *MockCasbinEnforcer) Enforce(role, resource, method string) *echo.HTTPError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enforce", role, resource, method)
	ret0, _ := ret[0].(*echo.HTTPError)
	return ret0
}

// Enforce indicates an expected call of Enforce.
func (mr *MockCasbinEnforcerMockRecorder) Enforce(role, resource, method any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enforce", reflect.TypeOf((*MockCasbinEnforcer)(nil).Enforce), role, resource, method)
}
