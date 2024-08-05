// Code generated by MockGen. DO NOT EDIT.
// Source: internal/api/readerStream/read.go

// Package readerstream is a generated GoMock package.
package readerstream

import (
	context "context"
	reflect "reflect"

	frequency_servicev1 "github.com/dadmaramf/protos/gen/go/frequency_service"
	gomock "github.com/golang/mock/gomock"
)

// MockReaderWriterStream is a mock of ReaderWriterStream interface.
type MockReaderWriterStream struct {
	ctrl     *gomock.Controller
	recorder *MockReaderWriterStreamMockRecorder
}

// MockReaderWriterStreamMockRecorder is the mock recorder for MockReaderWriterStream.
type MockReaderWriterStreamMockRecorder struct {
	mock *MockReaderWriterStream
}

// NewMockReaderWriterStream creates a new mock instance.
func NewMockReaderWriterStream(ctrl *gomock.Controller) *MockReaderWriterStream {
	mock := &MockReaderWriterStream{ctrl: ctrl}
	mock.recorder = &MockReaderWriterStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReaderWriterStream) EXPECT() *MockReaderWriterStreamMockRecorder {
	return m.recorder
}

// Recv mocks base method.
func (m *MockReaderWriterStream) Recv() (*frequency_servicev1.FrequencyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*frequency_servicev1.FrequencyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockReaderWriterStreamMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockReaderWriterStream)(nil).Recv))
}

// WriteAnomalies mocks base method.
func (m *MockReaderWriterStream) WriteAnomalies(ctx context.Context, data *frequency_servicev1.FrequencyResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteAnomalies", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteAnomalies indicates an expected call of WriteAnomalies.
func (mr *MockReaderWriterStreamMockRecorder) WriteAnomalies(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAnomalies", reflect.TypeOf((*MockReaderWriterStream)(nil).WriteAnomalies), ctx, data)
}