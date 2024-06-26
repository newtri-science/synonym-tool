// Code generated by MockGen. DO NOT EDIT.
// Source: utils/crypto.go
//
// Generated by this command:
//
//	mockgen -source utils/crypto.go -destination mocks/crypto.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCryptoer is a mock of Cryptoer interface.
type MockCryptoer struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoerMockRecorder
}

// MockCryptoerMockRecorder is the mock recorder for MockCryptoer.
type MockCryptoerMockRecorder struct {
	mock *MockCryptoer
}

// NewMockCryptoer creates a new mock instance.
func NewMockCryptoer(ctrl *gomock.Controller) *MockCryptoer {
	mock := &MockCryptoer{ctrl: ctrl}
	mock.recorder = &MockCryptoerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoer) EXPECT() *MockCryptoerMockRecorder {
	return m.recorder
}

// CompareHashAndPassword mocks base method.
func (m *MockCryptoer) CompareHashAndPassword(hashedPassword, password []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareHashAndPassword", hashedPassword, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompareHashAndPassword indicates an expected call of CompareHashAndPassword.
func (mr *MockCryptoerMockRecorder) CompareHashAndPassword(hashedPassword, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareHashAndPassword", reflect.TypeOf((*MockCryptoer)(nil).CompareHashAndPassword), hashedPassword, password)
}

// GenerateFromPassword mocks base method.
func (m *MockCryptoer) GenerateFromPassword(password []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateFromPassword", password)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateFromPassword indicates an expected call of GenerateFromPassword.
func (mr *MockCryptoerMockRecorder) GenerateFromPassword(password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateFromPassword", reflect.TypeOf((*MockCryptoer)(nil).GenerateFromPassword), password)
}
