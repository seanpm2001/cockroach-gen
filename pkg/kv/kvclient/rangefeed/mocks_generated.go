// Code generated by MockGen. DO NOT EDIT.
// Source: rangefeed.go

// Package rangefeed is a generated GoMock package.
package rangefeed

import (
	context "context"
	reflect "reflect"

	roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
	hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"
	gomock "github.com/golang/mock/gomock"
)

// MockkvDB is a mock of kvDB interface.
type MockkvDB struct {
	ctrl     *gomock.Controller
	recorder *MockkvDBMockRecorder
}

// MockkvDBMockRecorder is the mock recorder for MockkvDB.
type MockkvDBMockRecorder struct {
	mock *MockkvDB
}

// NewMockkvDB creates a new mock instance.
func NewMockkvDB(ctrl *gomock.Controller) *MockkvDB {
	mock := &MockkvDB{ctrl: ctrl}
	mock.recorder = &MockkvDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockkvDB) EXPECT() *MockkvDBMockRecorder {
	return m.recorder
}

// RangeFeed mocks base method.
func (m *MockkvDB) RangeFeed(ctx context.Context, spans []roachpb.Span, startFrom hlc.Timestamp, withDiff bool, eventC chan<- *roachpb.RangeFeedEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RangeFeed", ctx, spans, startFrom, withDiff, eventC)
	ret0, _ := ret[0].(error)
	return ret0
}

// RangeFeed indicates an expected call of RangeFeed.
func (mr *MockkvDBMockRecorder) RangeFeed(ctx, spans, startFrom, withDiff, eventC interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RangeFeed", reflect.TypeOf((*MockkvDB)(nil).RangeFeed), ctx, spans, startFrom, withDiff, eventC)
}

// Scan mocks base method.
func (m *MockkvDB) Scan(ctx context.Context, spans []roachpb.Span, asOf hlc.Timestamp, rowFn func(roachpb.KeyValue), cfg scanConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", ctx, spans, asOf, rowFn, cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockkvDBMockRecorder) Scan(ctx, spans, asOf, rowFn, cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockkvDB)(nil).Scan), ctx, spans, asOf, rowFn, cfg)
}
