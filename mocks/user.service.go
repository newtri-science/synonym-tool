// Code generated by MockGen. DO NOT EDIT.
// Source: services/user.service.go
//
// Generated by this command:
//
//	mockgen -source services/user.service.go -destination mocks/user.service.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/newtri-science/synonym-tool/model"
	gomock "go.uber.org/mock/gomock"
)

// MockUserServicer is a mock of UserServicer interface.
type MockUserServicer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServicerMockRecorder
}

// MockUserServicerMockRecorder is the mock recorder for MockUserServicer.
type MockUserServicerMockRecorder struct {
	mock *MockUserServicer
}

// NewMockUserServicer creates a new mock instance.
func NewMockUserServicer(ctrl *gomock.Controller) *MockUserServicer {
	mock := &MockUserServicer{ctrl: ctrl}
	mock.recorder = &MockUserServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServicer) EXPECT() *MockUserServicerMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserServicer) AddUser(user model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserServicerMockRecorder) AddUser(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserServicer)(nil).AddUser), user)
}

// Count mocks base method.
func (m *MockUserServicer) Count() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockUserServicerMockRecorder) Count() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockUserServicer)(nil).Count))
}

// DeleteUser mocks base method.
func (m *MockUserServicer) DeleteUser(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServicerMockRecorder) DeleteUser(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserServicer)(nil).DeleteUser), id)
}

// GetAllUsers mocks base method.
func (m *MockUserServicer) GetAllUsers() ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserServicerMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserServicer)(nil).GetAllUsers))
}

// GetByEmail mocks base method.
func (m *MockUserServicer) GetByEmail(email string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", email)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUserServicerMockRecorder) GetByEmail(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUserServicer)(nil).GetByEmail), email)
}

// GetById mocks base method.
func (m *MockUserServicer) GetById(id int) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUserServicerMockRecorder) GetById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUserServicer)(nil).GetById), id)
}
