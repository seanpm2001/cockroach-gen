// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cockroachdb/cockroach/pkg/roachprod/prometheus (interfaces: Client)

// Package tests is a generated GoMock package.
package tests

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	model "github.com/prometheus/common/model"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Query mocks base method.
func (m *MockClient) Query(arg0 context.Context, arg1 string, arg2 time.Time) (model.Value, v1.Warnings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1, arg2)
	ret0, _ := ret[0].(model.Value)
	ret1, _ := ret[1].(v1.Warnings)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Query indicates an expected call of Query.
func (mr *MockClientMockRecorder) Query(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockClient)(nil).Query), arg0, arg1, arg2)
}

// QueryRange mocks base method.
func (m *MockClient) QueryRange(arg0 context.Context, arg1 string, arg2 v1.Range) (model.Value, v1.Warnings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRange", arg0, arg1, arg2)
	ret0, _ := ret[0].(model.Value)
	ret1, _ := ret[1].(v1.Warnings)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryRange indicates an expected call of QueryRange.
func (mr *MockClientMockRecorder) QueryRange(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRange", reflect.TypeOf((*MockClient)(nil).QueryRange), arg0, arg1, arg2)
}
