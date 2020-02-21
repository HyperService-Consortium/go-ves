// Code generated by MockGen. DO NOT EDIT.
// Source: ../model/internal/abstraction/session.go

// Package mock is a generated GoMock package.
package mock

import (
	abstraction "github.com/HyperService-Consortium/go-ves/ves/model/internal/abstraction"
	database "github.com/HyperService-Consortium/go-ves/ves/model/internal/database"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// SessionDB is a mock of SessionDB interface
type SessionDB struct {
	ctrl     *gomock.Controller
	recorder *SessionDBMockRecorder
}

// SessionDBMockRecorder is the mock recorder for SessionDB
type SessionDBMockRecorder struct {
	mock *SessionDB
}

// NewSessionDB creates a new mock instance
func NewSessionDB(ctrl *gomock.Controller) *SessionDB {
	mock := &SessionDB{ctrl: ctrl}
	mock.recorder = &SessionDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *SessionDB) EXPECT() *SessionDBMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *SessionDB) Create(s *database.Session) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", s)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *SessionDBMockRecorder) Create(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*SessionDB)(nil).Create), s)
}

// Update mocks base method
func (m *SessionDB) Update(s *database.Session) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", s)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *SessionDBMockRecorder) Update(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*SessionDB)(nil).Update), s)
}

// UpdateFields mocks base method
func (m *SessionDB) UpdateFields(s *database.Session, fields []string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFields", s, fields)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFields indicates an expected call of UpdateFields
func (mr *SessionDBMockRecorder) UpdateFields(s, fields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFields", reflect.TypeOf((*SessionDB)(nil).UpdateFields), s, fields)
}

// Delete mocks base method
func (m *SessionDB) Delete(s *database.Session) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", s)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *SessionDBMockRecorder) Delete(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*SessionDB)(nil).Delete), s)
}

// Filter mocks base method
func (m *SessionDB) Filter(f *database.SessionFilter) ([]database.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", f)
	ret0, _ := ret[0].([]database.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *SessionDBMockRecorder) Filter(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*SessionDB)(nil).Filter), f)
}

// FilterI mocks base method
func (m *SessionDB) FilterI(f interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterI", f)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterI indicates an expected call of FilterI
func (mr *SessionDBMockRecorder) FilterI(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterI", reflect.TypeOf((*SessionDB)(nil).FilterI), f)
}

// ID mocks base method
func (m *SessionDB) ID(id uint) (*database.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID", id)
	ret0, _ := ret[0].(*database.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ID indicates an expected call of ID
func (mr *SessionDBMockRecorder) ID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*SessionDB)(nil).ID), id)
}

// QueryGUID mocks base method
func (m *SessionDB) QueryGUID(iscAddress string) (*database.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryGUID", iscAddress)
	ret0, _ := ret[0].(*database.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryGUID indicates an expected call of QueryGUID
func (mr *SessionDBMockRecorder) QueryGUID(iscAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryGUID", reflect.TypeOf((*SessionDB)(nil).QueryGUID), iscAddress)
}

// QueryGUIDByBytes mocks base method
func (m *SessionDB) QueryGUIDByBytes(iscAddress []byte) (*database.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryGUIDByBytes", iscAddress)
	ret0, _ := ret[0].(*database.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryGUIDByBytes indicates an expected call of QueryGUIDByBytes
func (mr *SessionDBMockRecorder) QueryGUIDByBytes(iscAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryGUIDByBytes", reflect.TypeOf((*SessionDB)(nil).QueryGUIDByBytes), iscAddress)
}

// Query mocks base method
func (m *SessionDB) Query(opts ...abstraction.SessionQueryOption) ([]database.Session, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].([]database.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *SessionDBMockRecorder) Query(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*SessionDB)(nil).Query), opts...)
}

// Scan mocks base method
func (m *SessionDB) Scan(desc interface{}, opts ...abstraction.SessionQueryOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{desc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan
func (mr *SessionDBMockRecorder) Scan(desc interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{desc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*SessionDB)(nil).Scan), varargs...)
}

// MockSessionQueryOption is a mock of SessionQueryOption interface
type MockSessionQueryOption struct {
	ctrl     *gomock.Controller
	recorder *MockSessionQueryOptionMockRecorder
}

// MockSessionQueryOptionMockRecorder is the mock recorder for MockSessionQueryOption
type MockSessionQueryOptionMockRecorder struct {
	mock *MockSessionQueryOption
}

// NewMockSessionQueryOption creates a new mock instance
func NewMockSessionQueryOption(ctrl *gomock.Controller) *MockSessionQueryOption {
	mock := &MockSessionQueryOption{ctrl: ctrl}
	mock.recorder = &MockSessionQueryOptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionQueryOption) EXPECT() *MockSessionQueryOptionMockRecorder {
	return m.recorder
}

// implementsSessionQuery mocks base method
func (m *MockSessionQueryOption) implementsSessionQuery() abstraction.SessionQueryOption {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "implementsSessionQuery")
	ret0, _ := ret[0].(abstraction.SessionQueryOption)
	return ret0
}

// implementsSessionQuery indicates an expected call of implementsSessionQuery
func (mr *MockSessionQueryOptionMockRecorder) implementsSessionQuery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "implementsSessionQuery", reflect.TypeOf((*MockSessionQueryOption)(nil).implementsSessionQuery))
}
