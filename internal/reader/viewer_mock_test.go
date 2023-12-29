// Code generated by MockGen. DO NOT EDIT.
// Source: internal/reader/view/view.go
//
// Generated by this command:
//
//	mockgen -source=internal/reader/view/view.go -package=reader Viewer
//

// Package reader is a generated GoMock package.
package reader

import (
	reflect "reflect"

	entity "github.com/bow/neon/internal/entity"
	tview "github.com/rivo/tview"
	gomock "go.uber.org/mock/gomock"
)

// MockViewer is a mock of Viewer interface.
type MockViewer struct {
	ctrl     *gomock.Controller
	recorder *MockViewerMockRecorder
}

// MockViewerMockRecorder is the mock recorder for MockViewer.
type MockViewerMockRecorder struct {
	mock *MockViewer
}

// NewMockViewer creates a new mock instance.
func NewMockViewer(ctrl *gomock.Controller) *MockViewer {
	mock := &MockViewer{ctrl: ctrl}
	mock.recorder = &MockViewerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockViewer) EXPECT() *MockViewerMockRecorder {
	return m.recorder
}

// ClearStatusBar mocks base method.
func (m *MockViewer) ClearStatusBar() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearStatusBar")
}

// ClearStatusBar indicates an expected call of ClearStatusBar.
func (mr *MockViewerMockRecorder) ClearStatusBar() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearStatusBar", reflect.TypeOf((*MockViewer)(nil).ClearStatusBar))
}

// CurrentFocus mocks base method.
func (m *MockViewer) CurrentFocus() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentFocus")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// CurrentFocus indicates an expected call of CurrentFocus.
func (mr *MockViewerMockRecorder) CurrentFocus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentFocus", reflect.TypeOf((*MockViewer)(nil).CurrentFocus))
}

// EntriesPane mocks base method.
func (m *MockViewer) EntriesPane() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EntriesPane")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// EntriesPane indicates an expected call of EntriesPane.
func (mr *MockViewerMockRecorder) EntriesPane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntriesPane", reflect.TypeOf((*MockViewer)(nil).EntriesPane))
}

// FeedsPane mocks base method.
func (m *MockViewer) FeedsPane() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeedsPane")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// FeedsPane indicates an expected call of FeedsPane.
func (mr *MockViewerMockRecorder) FeedsPane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeedsPane", reflect.TypeOf((*MockViewer)(nil).FeedsPane))
}

// FocusEntriesPane mocks base method.
func (m *MockViewer) FocusEntriesPane() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FocusEntriesPane")
}

// FocusEntriesPane indicates an expected call of FocusEntriesPane.
func (mr *MockViewerMockRecorder) FocusEntriesPane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FocusEntriesPane", reflect.TypeOf((*MockViewer)(nil).FocusEntriesPane))
}

// FocusFeedsPane mocks base method.
func (m *MockViewer) FocusFeedsPane() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FocusFeedsPane")
}

// FocusFeedsPane indicates an expected call of FocusFeedsPane.
func (mr *MockViewerMockRecorder) FocusFeedsPane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FocusFeedsPane", reflect.TypeOf((*MockViewer)(nil).FocusFeedsPane))
}

// FocusNextPane mocks base method.
func (m *MockViewer) FocusNextPane() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FocusNextPane")
}

// FocusNextPane indicates an expected call of FocusNextPane.
func (mr *MockViewerMockRecorder) FocusNextPane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FocusNextPane", reflect.TypeOf((*MockViewer)(nil).FocusNextPane))
}

// FocusPreviousPane mocks base method.
func (m *MockViewer) FocusPreviousPane() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FocusPreviousPane")
}

// FocusPreviousPane indicates an expected call of FocusPreviousPane.
func (mr *MockViewerMockRecorder) FocusPreviousPane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FocusPreviousPane", reflect.TypeOf((*MockViewer)(nil).FocusPreviousPane))
}

// FocusReadingPane mocks base method.
func (m *MockViewer) FocusReadingPane() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FocusReadingPane")
}

// FocusReadingPane indicates an expected call of FocusReadingPane.
func (mr *MockViewerMockRecorder) FocusReadingPane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FocusReadingPane", reflect.TypeOf((*MockViewer)(nil).FocusReadingPane))
}

// HideIntroPopup mocks base method.
func (m *MockViewer) HideIntroPopup() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HideIntroPopup")
}

// HideIntroPopup indicates an expected call of HideIntroPopup.
func (mr *MockViewerMockRecorder) HideIntroPopup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HideIntroPopup", reflect.TypeOf((*MockViewer)(nil).HideIntroPopup))
}

// MainPage mocks base method.
func (m *MockViewer) MainPage() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MainPage")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// MainPage indicates an expected call of MainPage.
func (mr *MockViewerMockRecorder) MainPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MainPage", reflect.TypeOf((*MockViewer)(nil).MainPage))
}

// NotifyErr mocks base method.
func (m *MockViewer) NotifyErr(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyErr", err)
}

// NotifyErr indicates an expected call of NotifyErr.
func (mr *MockViewerMockRecorder) NotifyErr(err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyErr", reflect.TypeOf((*MockViewer)(nil).NotifyErr), err)
}

// NotifyErrf mocks base method.
func (m *MockViewer) NotifyErrf(text string, a ...any) {
	m.ctrl.T.Helper()
	varargs := []any{text}
	for _, a_2 := range a {
		varargs = append(varargs, a_2)
	}
	m.ctrl.Call(m, "NotifyErrf", varargs...)
}

// NotifyErrf indicates an expected call of NotifyErrf.
func (mr *MockViewerMockRecorder) NotifyErrf(text any, a ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{text}, a...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyErrf", reflect.TypeOf((*MockViewer)(nil).NotifyErrf), varargs...)
}

// NotifyInfof mocks base method.
func (m *MockViewer) NotifyInfof(text string, a ...any) {
	m.ctrl.T.Helper()
	varargs := []any{text}
	for _, a_2 := range a {
		varargs = append(varargs, a_2)
	}
	m.ctrl.Call(m, "NotifyInfof", varargs...)
}

// NotifyInfof indicates an expected call of NotifyInfof.
func (mr *MockViewerMockRecorder) NotifyInfof(text any, a ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{text}, a...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyInfof", reflect.TypeOf((*MockViewer)(nil).NotifyInfof), varargs...)
}

// NotifyWarnf mocks base method.
func (m *MockViewer) NotifyWarnf(text string, a ...any) {
	m.ctrl.T.Helper()
	varargs := []any{text}
	for _, a_2 := range a {
		varargs = append(varargs, a_2)
	}
	m.ctrl.Call(m, "NotifyWarnf", varargs...)
}

// NotifyWarnf indicates an expected call of NotifyWarnf.
func (mr *MockViewerMockRecorder) NotifyWarnf(text any, a ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{text}, a...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyWarnf", reflect.TypeOf((*MockViewer)(nil).NotifyWarnf), varargs...)
}

// ReadingPane mocks base method.
func (m *MockViewer) ReadingPane() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadingPane")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// ReadingPane indicates an expected call of ReadingPane.
func (mr *MockViewerMockRecorder) ReadingPane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadingPane", reflect.TypeOf((*MockViewer)(nil).ReadingPane))
}

// Show mocks base method.
func (m *MockViewer) Show() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show")
	ret0, _ := ret[0].(error)
	return ret0
}

// Show indicates an expected call of Show.
func (mr *MockViewerMockRecorder) Show() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockViewer)(nil).Show))
}

// ShowAboutPopup mocks base method.
func (m *MockViewer) ShowAboutPopup() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShowAboutPopup")
}

// ShowAboutPopup indicates an expected call of ShowAboutPopup.
func (mr *MockViewerMockRecorder) ShowAboutPopup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowAboutPopup", reflect.TypeOf((*MockViewer)(nil).ShowAboutPopup))
}

// ShowFeedsInPane mocks base method.
func (m *MockViewer) ShowFeedsInPane(arg0 <-chan *entity.Feed) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShowFeedsInPane", arg0)
}

// ShowFeedsInPane indicates an expected call of ShowFeedsInPane.
func (mr *MockViewerMockRecorder) ShowFeedsInPane(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowFeedsInPane", reflect.TypeOf((*MockViewer)(nil).ShowFeedsInPane), arg0)
}

// ShowHelpPopup mocks base method.
func (m *MockViewer) ShowHelpPopup() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShowHelpPopup")
}

// ShowHelpPopup indicates an expected call of ShowHelpPopup.
func (mr *MockViewerMockRecorder) ShowHelpPopup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowHelpPopup", reflect.TypeOf((*MockViewer)(nil).ShowHelpPopup))
}

// ShowIntroPopup mocks base method.
func (m *MockViewer) ShowIntroPopup() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShowIntroPopup")
}

// ShowIntroPopup indicates an expected call of ShowIntroPopup.
func (mr *MockViewerMockRecorder) ShowIntroPopup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowIntroPopup", reflect.TypeOf((*MockViewer)(nil).ShowIntroPopup))
}

// ShowStatsPopup mocks base method.
func (m *MockViewer) ShowStatsPopup(arg0 <-chan *entity.Stats) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShowStatsPopup", arg0)
}

// ShowStatsPopup indicates an expected call of ShowStatsPopup.
func (mr *MockViewerMockRecorder) ShowStatsPopup(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowStatsPopup", reflect.TypeOf((*MockViewer)(nil).ShowStatsPopup), arg0)
}

// ToggleStatusBar mocks base method.
func (m *MockViewer) ToggleStatusBar() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ToggleStatusBar")
}

// ToggleStatusBar indicates an expected call of ToggleStatusBar.
func (mr *MockViewerMockRecorder) ToggleStatusBar() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleStatusBar", reflect.TypeOf((*MockViewer)(nil).ToggleStatusBar))
}

// UnfocusPane mocks base method.
func (m *MockViewer) UnfocusPane() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnfocusPane")
}

// UnfocusPane indicates an expected call of UnfocusPane.
func (mr *MockViewerMockRecorder) UnfocusPane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnfocusPane", reflect.TypeOf((*MockViewer)(nil).UnfocusPane))
}
