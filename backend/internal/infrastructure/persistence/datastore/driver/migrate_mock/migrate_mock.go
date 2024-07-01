// Code generated by MockGen. DO NOT EDIT.
// Source: migrate.go
//
// Generated by this command:
//
//	mockgen -source=migrate.go -destination=migrate_mock/migrate_mock.go -package migrate_mock
//

// Package migrate_mock is a generated GoMock package.
package migrate_mock

import (
	reflect "reflect"

	migrate "github.com/golang-migrate/migrate/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockLibMigrator is a mock of LibMigrator interface.
type MockLibMigrator struct {
	ctrl     *gomock.Controller
	recorder *MockLibMigratorMockRecorder
}

// MockLibMigratorMockRecorder is the mock recorder for MockLibMigrator.
type MockLibMigratorMockRecorder struct {
	mock *MockLibMigrator
}

// NewMockLibMigrator creates a new mock instance.
func NewMockLibMigrator(ctrl *gomock.Controller) *MockLibMigrator {
	mock := &MockLibMigrator{ctrl: ctrl}
	mock.recorder = &MockLibMigratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLibMigrator) EXPECT() *MockLibMigratorMockRecorder {
	return m.recorder
}

// New mocks base method.
func (m *MockLibMigrator) New(arg0, arg1 string) (*migrate.Migrate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0, arg1)
	ret0, _ := ret[0].(*migrate.Migrate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockLibMigratorMockRecorder) New(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockLibMigrator)(nil).New), arg0, arg1)
}

// MockMigrateClient is a mock of MigrateClient interface.
type MockMigrateClient struct {
	ctrl     *gomock.Controller
	recorder *MockMigrateClientMockRecorder
}

// MockMigrateClientMockRecorder is the mock recorder for MockMigrateClient.
type MockMigrateClientMockRecorder struct {
	mock *MockMigrateClient
}

// NewMockMigrateClient creates a new mock instance.
func NewMockMigrateClient(ctrl *gomock.Controller) *MockMigrateClient {
	mock := &MockMigrateClient{ctrl: ctrl}
	mock.recorder = &MockMigrateClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigrateClient) EXPECT() *MockMigrateClientMockRecorder {
	return m.recorder
}

// Up mocks base method.
func (m *MockMigrateClient) Up() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Up")
	ret0, _ := ret[0].(error)
	return ret0
}

// Up indicates an expected call of Up.
func (mr *MockMigrateClientMockRecorder) Up() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Up", reflect.TypeOf((*MockMigrateClient)(nil).Up))
}