// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package answersheet_usecase is a generated GoMock package.
package answersheet_usecase

import (
	context "context"
	sql "database/sql"
	entities "picket/src/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// BeginTransaction mocks base method.
func (m *MockIRepository) BeginTransaction(ctx context.Context, opts ...*sql.TxOptions) context.Context {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BeginTransaction", varargs...)
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// BeginTransaction indicates an expected call of BeginTransaction.
func (mr *MockIRepositoryMockRecorder) BeginTransaction(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTransaction", reflect.TypeOf((*MockIRepository)(nil).BeginTransaction), varargs...)
}

// Commit mocks base method.
func (m *MockIRepository) Commit(ctx context.Context) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", ctx)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockIRepositoryMockRecorder) Commit(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockIRepository)(nil).Commit), ctx)
}

// GetLatestEvent mocks base method.
func (m *MockIRepository) GetLatestEvent(ctx context.Context, userId, testId, number int) ([]entities.AnswerSheetEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestEvent", ctx, userId, testId, number)
	ret0, _ := ret[0].([]entities.AnswerSheetEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestEvent indicates an expected call of GetLatestEvent.
func (mr *MockIRepositoryMockRecorder) GetLatestEvent(ctx, userId, testId, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestEvent", reflect.TypeOf((*MockIRepository)(nil).GetLatestEvent), ctx, userId, testId, number)
}

// Rollback mocks base method.
func (m *MockIRepository) Rollback(ctx context.Context) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", ctx)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockIRepositoryMockRecorder) Rollback(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockIRepository)(nil).Rollback), ctx)
}

// SendEvent mocks base method.
func (m *MockIRepository) SendEvent(ctx context.Context, event *entities.AnswerSheetEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEvent", ctx, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEvent indicates an expected call of SendEvent.
func (mr *MockIRepositoryMockRecorder) SendEvent(ctx, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEvent", reflect.TypeOf((*MockIRepository)(nil).SendEvent), ctx, event)
}

// MockITestUsecase is a mock of ITestUsecase interface.
type MockITestUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockITestUsecaseMockRecorder
}

// MockITestUsecaseMockRecorder is the mock recorder for MockITestUsecase.
type MockITestUsecaseMockRecorder struct {
	mock *MockITestUsecase
}

// NewMockITestUsecase creates a new mock instance.
func NewMockITestUsecase(ctrl *gomock.Controller) *MockITestUsecase {
	mock := &MockITestUsecase{ctrl: ctrl}
	mock.recorder = &MockITestUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITestUsecase) EXPECT() *MockITestUsecaseMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockITestUsecase) GetById(ctx context.Context, id int) (*entities.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*entities.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockITestUsecaseMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockITestUsecase)(nil).GetById), ctx, id)
}

// GetContent mocks base method.
func (m *MockITestUsecase) GetContent(ctx context.Context, testId int) (*entities.TestContent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContent", ctx, testId)
	ret0, _ := ret[0].(*entities.TestContent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContent indicates an expected call of GetContent.
func (mr *MockITestUsecaseMockRecorder) GetContent(ctx, testId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContent", reflect.TypeOf((*MockITestUsecase)(nil).GetContent), ctx, testId)
}