// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package wx is a generated GoMock package.
package wx

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHTTPClient is a mock of HTTPClient interface.
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientMockRecorder
}

// MockHTTPClientMockRecorder is the mock recorder for MockHTTPClient.
type MockHTTPClientMockRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance.
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClient) EXPECT() *MockHTTPClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockHTTPClient) Get(ctx context.Context, reqURL string, options ...HTTPOption) ([]byte, error) {
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
func (mr *MockHTTPClientMockRecorder) Get(ctx, reqURL interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, reqURL}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHTTPClient)(nil).Get), varargs...)
}

// Post mocks base method.
func (m *MockHTTPClient) Post(ctx context.Context, reqURL string, body []byte, options ...HTTPOption) ([]byte, error) {
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
func (mr *MockHTTPClientMockRecorder) Post(ctx, reqURL, body interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, reqURL, body}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockHTTPClient)(nil).Post), varargs...)
}

// PostXML mocks base method.
func (m *MockHTTPClient) PostXML(ctx context.Context, reqURL string, body WXML, options ...HTTPOption) ([]byte, error) {
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
func (mr *MockHTTPClientMockRecorder) PostXML(ctx, reqURL, body interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, reqURL, body}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostXML", reflect.TypeOf((*MockHTTPClient)(nil).PostXML), varargs...)
}

// Upload mocks base method.
func (m *MockHTTPClient) Upload(ctx context.Context, reqURL string, form UploadForm, options ...HTTPOption) ([]byte, error) {
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
func (mr *MockHTTPClientMockRecorder) Upload(ctx, reqURL, form interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, reqURL, form}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockHTTPClient)(nil).Upload), varargs...)
}

// MockUploadForm is a mock of UploadForm interface.
type MockUploadForm struct {
	ctrl     *gomock.Controller
	recorder *MockUploadFormMockRecorder
}

// MockUploadFormMockRecorder is the mock recorder for MockUploadForm.
type MockUploadFormMockRecorder struct {
	mock *MockUploadForm
}

// NewMockUploadForm creates a new mock instance.
func NewMockUploadForm(ctrl *gomock.Controller) *MockUploadForm {
	mock := &MockUploadForm{ctrl: ctrl}
	mock.recorder = &MockUploadFormMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadForm) EXPECT() *MockUploadFormMockRecorder {
	return m.recorder
}

// Buffer mocks base method.
func (m *MockUploadForm) Buffer() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Buffer")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Buffer indicates an expected call of Buffer.
func (mr *MockUploadFormMockRecorder) Buffer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Buffer", reflect.TypeOf((*MockUploadForm)(nil).Buffer))
}

// ExtraFields mocks base method.
func (m *MockUploadForm) ExtraFields() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtraFields")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// ExtraFields indicates an expected call of ExtraFields.
func (mr *MockUploadFormMockRecorder) ExtraFields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtraFields", reflect.TypeOf((*MockUploadForm)(nil).ExtraFields))
}

// FieldName mocks base method.
func (m *MockUploadForm) FieldName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FieldName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FieldName indicates an expected call of FieldName.
func (mr *MockUploadFormMockRecorder) FieldName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FieldName", reflect.TypeOf((*MockUploadForm)(nil).FieldName))
}

// FileName mocks base method.
func (m *MockUploadForm) FileName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FileName indicates an expected call of FileName.
func (mr *MockUploadFormMockRecorder) FileName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileName", reflect.TypeOf((*MockUploadForm)(nil).FileName))
}

// MockAction is a mock of Action interface.
type MockAction struct {
	ctrl     *gomock.Controller
	recorder *MockActionMockRecorder
}

// MockActionMockRecorder is the mock recorder for MockAction.
type MockActionMockRecorder struct {
	mock *MockAction
}

// NewMockAction creates a new mock instance.
func NewMockAction(ctrl *gomock.Controller) *MockAction {
	mock := &MockAction{ctrl: ctrl}
	mock.recorder = &MockActionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAction) EXPECT() *MockActionMockRecorder {
	return m.recorder
}

// Body mocks base method.
func (m *MockAction) Body() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Body indicates an expected call of Body.
func (mr *MockActionMockRecorder) Body() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockAction)(nil).Body))
}

// Decode mocks base method.
func (m *MockAction) Decode() func([]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode")
	ret0, _ := ret[0].(func([]byte) error)
	return ret0
}

// Decode indicates an expected call of Decode.
func (mr *MockActionMockRecorder) Decode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockAction)(nil).Decode))
}

// Method mocks base method.
func (m *MockAction) Method() HTTPMethod {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method")
	ret0, _ := ret[0].(HTTPMethod)
	return ret0
}

// Method indicates an expected call of Method.
func (mr *MockActionMockRecorder) Method() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method", reflect.TypeOf((*MockAction)(nil).Method))
}

// TLS mocks base method.
func (m *MockAction) TLS() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TLS")
	ret0, _ := ret[0].(bool)
	return ret0
}

// TLS indicates an expected call of TLS.
func (mr *MockActionMockRecorder) TLS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TLS", reflect.TypeOf((*MockAction)(nil).TLS))
}

// URL mocks base method.
func (m *MockAction) URL(accessToken ...string) string {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range accessToken {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "URL", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// URL indicates an expected call of URL.
func (mr *MockActionMockRecorder) URL(accessToken ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockAction)(nil).URL), accessToken...)
}

// UploadForm mocks base method.
func (m *MockAction) UploadForm() UploadForm {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadForm")
	ret0, _ := ret[0].(UploadForm)
	return ret0
}

// UploadForm indicates an expected call of UploadForm.
func (mr *MockActionMockRecorder) UploadForm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadForm", reflect.TypeOf((*MockAction)(nil).UploadForm))
}

// WXML mocks base method.
func (m *MockAction) WXML(appid, mchid, nonce string) (WXML, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WXML", appid, mchid, nonce)
	ret0, _ := ret[0].(WXML)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WXML indicates an expected call of WXML.
func (mr *MockActionMockRecorder) WXML(appid, mchid, nonce interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WXML", reflect.TypeOf((*MockAction)(nil).WXML), appid, mchid, nonce)
}
