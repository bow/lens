// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store.go

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	store "github.com/bow/courier/internal/store"
	gomock "github.com/golang/mock/gomock"
	gofeed "github.com/mmcdole/gofeed"
)

// MockFeedStore is a mock of FeedStore interface.
type MockFeedStore struct {
	ctrl     *gomock.Controller
	recorder *MockFeedStoreMockRecorder
}

// MockFeedStoreMockRecorder is the mock recorder for MockFeedStore.
type MockFeedStoreMockRecorder struct {
	mock *MockFeedStore
}

// NewMockFeedStore creates a new mock instance.
func NewMockFeedStore(ctrl *gomock.Controller) *MockFeedStore {
	mock := &MockFeedStore{ctrl: ctrl}
	mock.recorder = &MockFeedStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeedStore) EXPECT() *MockFeedStoreMockRecorder {
	return m.recorder
}

// AddFeed mocks base method.
func (m *MockFeedStore) AddFeed(arg0 context.Context, arg1 *gofeed.Feed, arg2, arg3 *string, arg4 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFeed", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFeed indicates an expected call of AddFeed.
func (mr *MockFeedStoreMockRecorder) AddFeed(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFeed", reflect.TypeOf((*MockFeedStore)(nil).AddFeed), arg0, arg1, arg2, arg3, arg4)
}

// ListFeeds mocks base method.
func (m *MockFeedStore) ListFeeds(arg0 context.Context) ([]*store.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFeeds", arg0)
	ret0, _ := ret[0].([]*store.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFeeds indicates an expected call of ListFeeds.
func (mr *MockFeedStoreMockRecorder) ListFeeds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFeeds", reflect.TypeOf((*MockFeedStore)(nil).ListFeeds), arg0)
}
