// Code generated by MockGen. DO NOT EDIT.
// Source: ../domain/manager.go
//
// Generated by this command:
//
//	mockgen -source=../domain/manager.go -destination=../unit/mocks/manager_mock.go
//

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	domain "app/domain"
	context "context"
	reflect "reflect"
	time "time"

	logrus "github.com/sirupsen/logrus"
	gomock "go.uber.org/mock/gomock"
)

// MockIManagerRepo is a mock of IManagerRepo interface.
type MockIManagerRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIManagerRepoMockRecorder
}

// MockIManagerRepoMockRecorder is the mock recorder for MockIManagerRepo.
type MockIManagerRepoMockRecorder struct {
	mock *MockIManagerRepo
}

// NewMockIManagerRepo creates a new mock instance.
func NewMockIManagerRepo(ctrl *gomock.Controller) *MockIManagerRepo {
	mock := &MockIManagerRepo{ctrl: ctrl}
	mock.recorder = &MockIManagerRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIManagerRepo) EXPECT() *MockIManagerRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIManagerRepo) Add(c context.Context, mngr *domain.Manager, lg *logrus.Logger) (domain.Manager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", c, mngr, lg)
	ret0, _ := ret[0].(domain.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockIManagerRepoMockRecorder) Add(c, mngr, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIManagerRepo)(nil).Add), c, mngr, lg)
}

// Delete mocks base method.
func (m *MockIManagerRepo) Delete(c context.Context, id int, lg *logrus.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c, id, lg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIManagerRepoMockRecorder) Delete(c, id, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIManagerRepo)(nil).Delete), c, id, lg)
}

// GetAllRequests mocks base method.
func (m *MockIManagerRepo) GetAllRequests(c context.Context, id int, lg *logrus.Logger) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRequests", c, id, lg)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRequests indicates an expected call of GetAllRequests.
func (mr *MockIManagerRepoMockRecorder) GetAllRequests(c, id, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRequests", reflect.TypeOf((*MockIManagerRepo)(nil).GetAllRequests), c, id, lg)
}

// GetByDepartment mocks base method.
func (m *MockIManagerRepo) GetByDepartment(c context.Context, department string, lg *logrus.Logger) ([]domain.Manager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDepartment", c, department, lg)
	ret0, _ := ret[0].([]domain.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDepartment indicates an expected call of GetByDepartment.
func (mr *MockIManagerRepoMockRecorder) GetByDepartment(c, department, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDepartment", reflect.TypeOf((*MockIManagerRepo)(nil).GetByDepartment), c, department, lg)
}

// GetById mocks base method.
func (m *MockIManagerRepo) GetById(c context.Context, id int, lg *logrus.Logger) (domain.Manager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", c, id, lg)
	ret0, _ := ret[0].(domain.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIManagerRepoMockRecorder) GetById(c, id, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIManagerRepo)(nil).GetById), c, id, lg)
}

// GetByNameSurname mocks base method.
func (m *MockIManagerRepo) GetByNameSurname(c context.Context, name, surname string, lg *logrus.Logger) ([]domain.Manager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNameSurname", c, name, surname, lg)
	ret0, _ := ret[0].([]domain.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNameSurname indicates an expected call of GetByNameSurname.
func (mr *MockIManagerRepoMockRecorder) GetByNameSurname(c, name, surname, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNameSurname", reflect.TypeOf((*MockIManagerRepo)(nil).GetByNameSurname), c, name, surname, lg)
}

// GetLimit mocks base method.
func (m *MockIManagerRepo) GetLimit(c context.Context, offset, limit int, lg *logrus.Logger) ([]domain.Manager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLimit", c, offset, limit, lg)
	ret0, _ := ret[0].([]domain.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLimit indicates an expected call of GetLimit.
func (mr *MockIManagerRepoMockRecorder) GetLimit(c, offset, limit, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLimit", reflect.TypeOf((*MockIManagerRepo)(nil).GetLimit), c, offset, limit, lg)
}

// GetNumberServedRequests mocks base method.
func (m *MockIManagerRepo) GetNumberServedRequests(c context.Context, id int, lg *logrus.Logger) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumberServedRequests", c, id, lg)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNumberServedRequests indicates an expected call of GetNumberServedRequests.
func (mr *MockIManagerRepoMockRecorder) GetNumberServedRequests(c, id, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumberServedRequests", reflect.TypeOf((*MockIManagerRepo)(nil).GetNumberServedRequests), c, id, lg)
}

// GetSumOnPeriod mocks base method.
func (m *MockIManagerRepo) GetSumOnPeriod(c context.Context, id int, from, to time.Time, lg *logrus.Logger) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSumOnPeriod", c, id, from, to, lg)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSumOnPeriod indicates an expected call of GetSumOnPeriod.
func (mr *MockIManagerRepoMockRecorder) GetSumOnPeriod(c, id, from, to, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSumOnPeriod", reflect.TypeOf((*MockIManagerRepo)(nil).GetSumOnPeriod), c, id, from, to, lg)
}

// Update mocks base method.
func (m *MockIManagerRepo) Update(c context.Context, id int, newState *domain.Manager, lg *logrus.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c, id, newState, lg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIManagerRepoMockRecorder) Update(c, id, newState, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIManagerRepo)(nil).Update), c, id, newState, lg)
}

// MockIManagerService is a mock of IManagerService interface.
type MockIManagerService struct {
	ctrl     *gomock.Controller
	recorder *MockIManagerServiceMockRecorder
}

// MockIManagerServiceMockRecorder is the mock recorder for MockIManagerService.
type MockIManagerServiceMockRecorder struct {
	mock *MockIManagerService
}

// NewMockIManagerService creates a new mock instance.
func NewMockIManagerService(ctrl *gomock.Controller) *MockIManagerService {
	mock := &MockIManagerService{ctrl: ctrl}
	mock.recorder = &MockIManagerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIManagerService) EXPECT() *MockIManagerServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIManagerService) Create(mngr *domain.Manager, lg *logrus.Logger) (domain.Manager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", mngr, lg)
	ret0, _ := ret[0].(domain.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIManagerServiceMockRecorder) Create(mngr, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIManagerService)(nil).Create), mngr, lg)
}

// Delete mocks base method.
func (m *MockIManagerService) Delete(id int, lg *logrus.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, lg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIManagerServiceMockRecorder) Delete(id, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIManagerService)(nil).Delete), id, lg)
}

// GetByDepartment mocks base method.
func (m *MockIManagerService) GetByDepartment(department string, lg *logrus.Logger) ([]domain.Manager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDepartment", department, lg)
	ret0, _ := ret[0].([]domain.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDepartment indicates an expected call of GetByDepartment.
func (mr *MockIManagerServiceMockRecorder) GetByDepartment(department, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDepartment", reflect.TypeOf((*MockIManagerService)(nil).GetByDepartment), department, lg)
}

// GetById mocks base method.
func (m *MockIManagerService) GetById(id int, lg *logrus.Logger) (domain.Manager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id, lg)
	ret0, _ := ret[0].(domain.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIManagerServiceMockRecorder) GetById(id, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIManagerService)(nil).GetById), id, lg)
}

// GetByNameSurname mocks base method.
func (m *MockIManagerService) GetByNameSurname(name, surname string, lg *logrus.Logger) ([]domain.Manager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNameSurname", name, surname, lg)
	ret0, _ := ret[0].([]domain.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNameSurname indicates an expected call of GetByNameSurname.
func (mr *MockIManagerServiceMockRecorder) GetByNameSurname(name, surname, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNameSurname", reflect.TypeOf((*MockIManagerService)(nil).GetByNameSurname), name, surname, lg)
}

// GetStatisticsForManager mocks base method.
func (m *MockIManagerService) GetStatisticsForManager(id int, from, to time.Time, lg *logrus.Logger) (domain.Statistics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatisticsForManager", id, from, to, lg)
	ret0, _ := ret[0].(domain.Statistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatisticsForManager indicates an expected call of GetStatisticsForManager.
func (mr *MockIManagerServiceMockRecorder) GetStatisticsForManager(id, from, to, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatisticsForManager", reflect.TypeOf((*MockIManagerService)(nil).GetStatisticsForManager), id, from, to, lg)
}

// GetStatisticsForManagers mocks base method.
func (m *MockIManagerService) GetStatisticsForManagers(from, to time.Time, offset, limit int, lg *logrus.Logger) ([]domain.Statistics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatisticsForManagers", from, to, offset, limit, lg)
	ret0, _ := ret[0].([]domain.Statistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatisticsForManagers indicates an expected call of GetStatisticsForManagers.
func (mr *MockIManagerServiceMockRecorder) GetStatisticsForManagers(from, to, offset, limit, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatisticsForManagers", reflect.TypeOf((*MockIManagerService)(nil).GetStatisticsForManagers), from, to, offset, limit, lg)
}

// Update mocks base method.
func (m *MockIManagerService) Update(id int, newState *domain.Manager, lg *logrus.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, newState, lg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIManagerServiceMockRecorder) Update(id, newState, lg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIManagerService)(nil).Update), id, newState, lg)
}