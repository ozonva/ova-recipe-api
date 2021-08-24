// Code generated by MockGen. DO NOT EDIT.
// Source: ova-recipe-api/internal/ticker (interfaces: Ticker)

// Package ticker is a generated GoMock package.
package ticker

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockTicker is a mock of Ticker interface.
type MockTicker struct {
	ctrl     *gomock.Controller
	recorder *MockTickerMockRecorder
}

// MockTickerMockRecorder is the mock recorder for MockTicker.
type MockTickerMockRecorder struct {
	mock *MockTicker
}

// NewMockTicker creates a new mock instance.
func NewMockTicker(ctrl *gomock.Controller) *MockTicker {
	mock := &MockTicker{ctrl: ctrl}
	mock.recorder = &MockTickerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicker) EXPECT() *MockTickerMockRecorder {
	return m.recorder
}

// Chanel mocks base method.
func (m *MockTicker) Chanel() <-chan time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chanel")
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

// Chanel indicates an expected call of Chanel.
func (mr *MockTickerMockRecorder) Chanel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chanel", reflect.TypeOf((*MockTicker)(nil).Chanel))
}

// Stop mocks base method.
func (m *MockTicker) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockTickerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTicker)(nil).Stop))
}