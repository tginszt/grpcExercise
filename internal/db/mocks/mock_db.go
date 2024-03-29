// Code generated by MockGen. DO NOT EDIT.
// Source: grpcExercise/internal/db (interfaces: Storage)

// Package mocks is a generated GoMock package.
package mocks

import (
	users "grpcExercise/internal/users"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockStorage) CreateUser(arg0 *users.User) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateUser", arg0)
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStorageMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStorage)(nil).CreateUser), arg0)
}

// DeleteUser mocks base method.
func (m *MockStorage) DeleteUser(arg0 *users.Id) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStorageMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStorage)(nil).DeleteUser), arg0)
}

// ReadUser mocks base method.
func (m *MockStorage) ReadUser(arg0 *users.Id) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUser", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// ReadUser indicates an expected call of ReadUser.
func (mr *MockStorageMockRecorder) ReadUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUser", reflect.TypeOf((*MockStorage)(nil).ReadUser), arg0)
}

// ReadUsers mocks base method.
func (m *MockStorage) ReadUsers() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUsers")
	ret0, _ := ret[0].(string)
	return ret0
}

// ReadUsers indicates an expected call of ReadUsers.
func (mr *MockStorageMockRecorder) ReadUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUsers", reflect.TypeOf((*MockStorage)(nil).ReadUsers))
}

// UpdateUser mocks base method.
func (m *MockStorage) UpdateUser(arg0 *users.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStorageMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStorage)(nil).UpdateUser), arg0)
}
