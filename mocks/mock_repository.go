// Code generated by MockGen. DO NOT EDIT.
// Source: repository/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	entity "transaction-api/entity"
	repository "transaction-api/repository"

	gomock "github.com/golang/mock/gomock"
)

// MockAccountRepository is a mock of AccountRepository interface.
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository.
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance.
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountRepository) Create(account *entity.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAccountRepositoryMockRecorder) Create(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountRepository)(nil).Create), account)
}

// Find mocks base method.
func (m *MockAccountRepository) Find(id int64) (entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockAccountRepositoryMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockAccountRepository)(nil).Find), id)
}

// MockMigrationRepository is a mock of MigrationRepository interface.
type MockMigrationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMigrationRepositoryMockRecorder
}

// MockMigrationRepositoryMockRecorder is the mock recorder for MockMigrationRepository.
type MockMigrationRepositoryMockRecorder struct {
	mock *MockMigrationRepository
}

// NewMockMigrationRepository creates a new mock instance.
func NewMockMigrationRepository(ctrl *gomock.Controller) *MockMigrationRepository {
	mock := &MockMigrationRepository{ctrl: ctrl}
	mock.recorder = &MockMigrationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigrationRepository) EXPECT() *MockMigrationRepositoryMockRecorder {
	return m.recorder
}

// Migrate mocks base method.
func (m *MockMigrationRepository) Migrate(migrationDir string, migrationType repository.MigrationType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrate", migrationDir, migrationType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockMigrationRepositoryMockRecorder) Migrate(migrationDir, migrationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockMigrationRepository)(nil).Migrate), migrationDir, migrationType)
}

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTransactionRepository) Create(transaction *entity.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", transaction)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTransactionRepositoryMockRecorder) Create(transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTransactionRepository)(nil).Create), transaction)
}

// Find mocks base method.
func (m *MockTransactionRepository) Find(id int64) (entity.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(entity.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockTransactionRepositoryMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockTransactionRepository)(nil).Find), id)
}

// FindByAccountID mocks base method.
func (m *MockTransactionRepository) FindByAccountID(id int64) ([]entity.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAccountID", id)
	ret0, _ := ret[0].([]entity.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByAccountID indicates an expected call of FindByAccountID.
func (mr *MockTransactionRepositoryMockRecorder) FindByAccountID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAccountID", reflect.TypeOf((*MockTransactionRepository)(nil).FindByAccountID), id)
}

// Update mocks base method.
func (m *MockTransactionRepository) Update(transaction *entity.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", transaction)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTransactionRepositoryMockRecorder) Update(transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTransactionRepository)(nil).Update), transaction)
}
