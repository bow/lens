// Code generated by MockGen. DO NOT EDIT.
// Source: internal/reader/backend/backend.go
//
// Generated by this command:
//
//	mockgen -source=internal/reader/backend/backend.go -package=reader Backend
//

// Package reader is a generated GoMock package.
package reader

import (
	context "context"
	reflect "reflect"

	entity "github.com/bow/neon/internal/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// GetStatsF mocks base method.
func (m *MockBackend) GetStatsF(arg0 context.Context) func() (*entity.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatsF", arg0)
	ret0, _ := ret[0].(func() (*entity.Stats, error))
	return ret0
}

// GetStatsF indicates an expected call of GetStatsF.
func (mr *MockBackendMockRecorder) GetStatsF(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatsF", reflect.TypeOf((*MockBackend)(nil).GetStatsF), arg0)
}

// ListFeedsF mocks base method.
func (m *MockBackend) ListFeedsF(arg0 context.Context) func() ([]*entity.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFeedsF", arg0)
	ret0, _ := ret[0].(func() ([]*entity.Feed, error))
	return ret0
}

// ListFeedsF indicates an expected call of ListFeedsF.
func (mr *MockBackendMockRecorder) ListFeedsF(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFeedsF", reflect.TypeOf((*MockBackend)(nil).ListFeedsF), arg0)
}

// PullFeeds mocks base method.
func (m *MockBackend) PullFeeds(arg0 context.Context, arg1 []entity.ID) <-chan entity.PullResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullFeeds", arg0, arg1)
	ret0, _ := ret[0].(<-chan entity.PullResult)
	return ret0
}

// PullFeeds indicates an expected call of PullFeeds.
func (mr *MockBackendMockRecorder) PullFeeds(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullFeeds", reflect.TypeOf((*MockBackend)(nil).PullFeeds), arg0, arg1)
}

// StringF mocks base method.
func (m *MockBackend) StringF() func() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StringF")
	ret0, _ := ret[0].(func() string)
	return ret0
}

// StringF indicates an expected call of StringF.
func (mr *MockBackendMockRecorder) StringF() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StringF", reflect.TypeOf((*MockBackend)(nil).StringF))
}
