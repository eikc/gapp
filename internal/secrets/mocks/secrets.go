// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/secrets/secrets.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gh "github.com/eikc/gapp/internal/gh"
	gomock "github.com/golang/mock/gomock"
	os "os"
	reflect "reflect"
)

// MockFileReader is a mock of FileReader interface
type MockFileReader struct {
	ctrl     *gomock.Controller
	recorder *MockFileReaderMockRecorder
}

// MockFileReaderMockRecorder is the mock recorder for MockFileReader
type MockFileReaderMockRecorder struct {
	mock *MockFileReader
}

// NewMockFileReader creates a new mock instance
func NewMockFileReader(ctrl *gomock.Controller) *MockFileReader {
	mock := &MockFileReader{ctrl: ctrl}
	mock.recorder = &MockFileReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileReader) EXPECT() *MockFileReaderMockRecorder {
	return m.recorder
}

// IsFile mocks base method
func (m *MockFileReader) IsFile(path string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFile", path)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFile indicates an expected call of IsFile
func (mr *MockFileReaderMockRecorder) IsFile(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFile", reflect.TypeOf((*MockFileReader)(nil).IsFile), path)
}

// ReadFile mocks base method
func (m *MockFileReader) ReadFile(path string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", path)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile
func (mr *MockFileReaderMockRecorder) ReadFile(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockFileReader)(nil).ReadFile), path)
}

// ReadDir mocks base method
func (m *MockFileReader) ReadDir(path string) ([]os.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDir", path)
	ret0, _ := ret[0].([]os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDir indicates an expected call of ReadDir
func (mr *MockFileReaderMockRecorder) ReadDir(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDir", reflect.TypeOf((*MockFileReader)(nil).ReadDir), path)
}

// MockActionsClient is a mock of ActionsClient interface
type MockActionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockActionsClientMockRecorder
}

// MockActionsClientMockRecorder is the mock recorder for MockActionsClient
type MockActionsClientMockRecorder struct {
	mock *MockActionsClient
}

// NewMockActionsClient creates a new mock instance
func NewMockActionsClient(ctrl *gomock.Controller) *MockActionsClient {
	mock := &MockActionsClient{ctrl: ctrl}
	mock.recorder = &MockActionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActionsClient) EXPECT() *MockActionsClientMockRecorder {
	return m.recorder
}

// GetPublicKey mocks base method
func (m *MockActionsClient) GetPublicKey(ctx context.Context, owner, repo string) ([]byte, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKey", ctx, owner, repo)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPublicKey indicates an expected call of GetPublicKey
func (mr *MockActionsClientMockRecorder) GetPublicKey(ctx, owner, repo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockActionsClient)(nil).GetPublicKey), ctx, owner, repo)
}

// AddOrUpdateSecret mocks base method
func (m *MockActionsClient) AddOrUpdateSecret(ctx context.Context, owner, repo string, secret gh.SecretParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrUpdateSecret", ctx, owner, repo, secret)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrUpdateSecret indicates an expected call of AddOrUpdateSecret
func (mr *MockActionsClientMockRecorder) AddOrUpdateSecret(ctx, owner, repo, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateSecret", reflect.TypeOf((*MockActionsClient)(nil).AddOrUpdateSecret), ctx, owner, repo, secret)
}

// MockEncryptionWriter is a mock of EncryptionWriter interface
type MockEncryptionWriter struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptionWriterMockRecorder
}

// MockEncryptionWriterMockRecorder is the mock recorder for MockEncryptionWriter
type MockEncryptionWriterMockRecorder struct {
	mock *MockEncryptionWriter
}

// NewMockEncryptionWriter creates a new mock instance
func NewMockEncryptionWriter(ctrl *gomock.Controller) *MockEncryptionWriter {
	mock := &MockEncryptionWriter{ctrl: ctrl}
	mock.recorder = &MockEncryptionWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEncryptionWriter) EXPECT() *MockEncryptionWriterMockRecorder {
	return m.recorder
}

// Encrypt mocks base method
func (m *MockEncryptionWriter) Encrypt(value string, pkey []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", value, pkey)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt
func (mr *MockEncryptionWriterMockRecorder) Encrypt(value, pkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockEncryptionWriter)(nil).Encrypt), value, pkey)
}
