// Code generated by MockGen. DO NOT EDIT.
// Source: ../domain/sale.go
//
// Generated by this command:
//
//	mockgen -source=../domain/sale.go -destination=../unit/mocks/sale_mock.go
//

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	domain "app/domain"
	context "context"
	reflect "reflect"

	logrus "github.com/sirupsen/logrus"
	gomock "go.uber.org/mock/gomock"
)

// MockISaleService is a mock of ISaleService interface.
type MockISaleService struct {
	ctrl     *gomock.Controller
	recorder *MockISaleServiceMockRecorder
}

// MockISaleServiceMockRecorder is the mock recorder for MockISaleService.
type MockISaleServiceMockRecorder struct {
	mock *MockISaleService
}

// NewMockISaleService creates a new mock instance.
func NewMockISaleService(ctrl *gomock.Controller) *MockISaleService {
	mock := &MockISaleService{ctrl: ctrl}
	mock.recorder = &MockISaleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISaleService) EXPECT() *MockISaleServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockISaleService) Create(sale *domain.Sale, lg *logrus.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", sale, lg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockISaleServiceMockRecorder) Create(sale, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockISaleService)(nil).Create), sale, lg)
}

// Delete mocks base method.
func (m *MockISaleService) Delete(id int, lg *logrus.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, lg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockISaleServiceMockRecorder) Delete(id, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockISaleService)(nil).Delete), id, lg)
}

// GetByCriteria mocks base method.
func (m *MockISaleService) GetByCriteria(saleCriteria *domain.Sale, offset, limit int, lg *logrus.Logger) ([]domain.Sale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCriteria", saleCriteria, offset, limit, lg)
	ret0, _ := ret[0].([]domain.Sale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCriteria indicates an expected call of GetByCriteria.
func (mr *MockISaleServiceMockRecorder) GetByCriteria(saleCriteria, offset, limit, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCriteria", reflect.TypeOf((*MockISaleService)(nil).GetByCriteria), saleCriteria, offset, limit, lg)
}

// GetById mocks base method.
func (m *MockISaleService) GetById(id int, lg *logrus.Logger) (domain.Sale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id, lg)
	ret0, _ := ret[0].(domain.Sale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockISaleServiceMockRecorder) GetById(id, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockISaleService)(nil).GetById), id, lg)
}

// GetSales mocks base method.
func (m *MockISaleService) GetSales(offset, limit int, lg *logrus.Logger) ([]domain.Sale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSales", offset, limit, lg)
	ret0, _ := ret[0].([]domain.Sale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSales indicates an expected call of GetSales.
func (mr *MockISaleServiceMockRecorder) GetSales(offset, limit, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSales", reflect.TypeOf((*MockISaleService)(nil).GetSales), offset, limit, lg)
}

// Update mocks base method.
func (m *MockISaleService) Update(id int, newState *domain.Sale, lg *logrus.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, newState, lg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockISaleServiceMockRecorder) Update(id, newState, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockISaleService)(nil).Update), id, newState, lg)
}

// MockISaleRepo is a mock of ISaleRepo interface.
type MockISaleRepo struct {
	ctrl     *gomock.Controller
	recorder *MockISaleRepoMockRecorder
}

// MockISaleRepoMockRecorder is the mock recorder for MockISaleRepo.
type MockISaleRepoMockRecorder struct {
	mock *MockISaleRepo
}

// NewMockISaleRepo creates a new mock instance.
func NewMockISaleRepo(ctrl *gomock.Controller) *MockISaleRepo {
	mock := &MockISaleRepo{ctrl: ctrl}
	mock.recorder = &MockISaleRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISaleRepo) EXPECT() *MockISaleRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockISaleRepo) Add(c context.Context, sale *domain.Sale, lg *logrus.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", c, sale, lg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockISaleRepoMockRecorder) Add(c, sale, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockISaleRepo)(nil).Add), c, sale, lg)
}

// Delete mocks base method.
func (m *MockISaleRepo) Delete(c context.Context, id int, lg *logrus.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c, id, lg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockISaleRepoMockRecorder) Delete(c, id, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockISaleRepo)(nil).Delete), c, id, lg)
}

// GetByCriteria mocks base method.
func (m *MockISaleRepo) GetByCriteria(c context.Context, offset, limit int, saleCriteria *domain.Sale, lg *logrus.Logger) ([]domain.Sale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCriteria", c, offset, limit, saleCriteria, lg)
	ret0, _ := ret[0].([]domain.Sale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCriteria indicates an expected call of GetByCriteria.
func (mr *MockISaleRepoMockRecorder) GetByCriteria(c, offset, limit, saleCriteria, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCriteria", reflect.TypeOf((*MockISaleRepo)(nil).GetByCriteria), c, offset, limit, saleCriteria, lg)
}

// GetById mocks base method.
func (m *MockISaleRepo) GetById(c context.Context, id int, lg *logrus.Logger) (domain.Sale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", c, id, lg)
	ret0, _ := ret[0].(domain.Sale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockISaleRepoMockRecorder) GetById(c, id, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockISaleRepo)(nil).GetById), c, id, lg)
}

// GetLimit mocks base method.
func (m *MockISaleRepo) GetLimit(c context.Context, offset, limit int, lg *logrus.Logger) ([]domain.Sale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLimit", c, offset, limit, lg)
	ret0, _ := ret[0].([]domain.Sale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLimit indicates an expected call of GetLimit.
func (mr *MockISaleRepoMockRecorder) GetLimit(c, offset, limit, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLimit", reflect.TypeOf((*MockISaleRepo)(nil).GetLimit), c, offset, limit, lg)
}

// Update mocks base method.
func (m *MockISaleRepo) Update(c context.Context, id int, newState *domain.Sale, lg *logrus.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c, id, newState, lg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockISaleRepoMockRecorder) Update(c, id, newState, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockISaleRepo)(nil).Update), c, id, newState, lg)
}