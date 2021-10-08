// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package wx is a generated GoMock package.
package wx

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	yiigo "github.com/shenghui0779/yiigo"
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

// Get mocks base method.
func (m *MockClient) Get(ctx context.Context, reqURL string, options ...yiigo.HTTPOption) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, reqURL}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClientMockRecorder) Get(ctx, reqURL interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, reqURL}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), varargs...)
}

// Post mocks base method.
func (m *MockClient) Post(ctx context.Context, reqURL string, body []byte, options ...yiigo.HTTPOption) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, reqURL, body}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Post", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockClientMockRecorder) Post(ctx, reqURL, body interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, reqURL, body}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockClient)(nil).Post), varargs...)
}

// PostXML mocks base method.
func (m *MockClient) PostXML(ctx context.Context, reqURL string, body WXML, options ...yiigo.HTTPOption) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, reqURL, body}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostXML", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostXML indicates an expected call of PostXML.
func (mr *MockClientMockRecorder) PostXML(ctx, reqURL, body interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, reqURL, body}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostXML", reflect.TypeOf((*MockClient)(nil).PostXML), varargs...)
}

// SetLogger mocks base method.
func (m *MockClient) SetLogger(l Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogger", l)
}

// SetLogger indicates an expected call of SetLogger.
func (mr *MockClientMockRecorder) SetLogger(l interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogger", reflect.TypeOf((*MockClient)(nil).SetLogger), l)
}

// Upload mocks base method.
func (m *MockClient) Upload(ctx context.Context, reqURL string, form yiigo.UploadForm, options ...yiigo.HTTPOption) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, reqURL, form}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Upload", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockClientMockRecorder) Upload(ctx, reqURL, form interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, reqURL, form}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockClient)(nil).Upload), varargs...)
}

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Log mocks base method.
func (m *MockLogger) Log(ctx context.Context, data *LogData) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Log", ctx, data)
}

// Log indicates an expected call of Log.
func (mr *MockLoggerMockRecorder) Log(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLogger)(nil).Log), ctx, data)
}
