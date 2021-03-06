// Code generated by MockGen. DO NOT EDIT.
// Source: ../model/fset.go

// Package mock is a generated GoMock package.
package mock

import (
	uip "github.com/HyperService-Consortium/go-uip/uip"
	types "github.com/HyperService-Consortium/go-ves/types"
	model "github.com/HyperService-Consortium/go-ves/ves/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// SessionFSet is a mock of SessionFSetI interface
type SessionFSet struct {
	ctrl     *gomock.Controller
	recorder *SessionFSetMockRecorder
}

// SessionFSetMockRecorder is the mock recorder for SessionFSet
type SessionFSetMockRecorder struct {
	mock *SessionFSet
}

// NewSessionFSet creates a new mock instance
func NewSessionFSet(ctrl *gomock.Controller) *SessionFSet {
	mock := &SessionFSet{ctrl: ctrl}
	mock.recorder = &SessionFSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *SessionFSet) EXPECT() *SessionFSetMockRecorder {
	return m.recorder
}

// AckForInit mocks base method
func (m *SessionFSet) AckForInit(arg0 *model.Session, arg1 uip.Account, arg2 uip.Signature) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AckForInit", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AckForInit indicates an expected call of AckForInit
func (mr *SessionFSetMockRecorder) AckForInit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AckForInit", reflect.TypeOf((*SessionFSet)(nil).AckForInit), arg0, arg1, arg2)
}

// FindTransaction mocks base method
func (m *SessionFSet) FindTransaction(arg0 []uint8, arg1 int64) ([]uint8, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTransaction", arg0, arg1)
	ret0, _ := ret[0].([]uint8)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTransaction indicates an expected call of FindTransaction
func (mr *SessionFSetMockRecorder) FindTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTransaction", reflect.TypeOf((*SessionFSet)(nil).FindTransaction), arg0, arg1)
}

// GetAccounts mocks base method
func (m *SessionFSet) GetAccounts(arg0 *model.Session) ([]uip.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", arg0)
	ret0, _ := ret[0].([]uip.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts
func (mr *SessionFSetMockRecorder) GetAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*SessionFSet)(nil).GetAccounts), arg0)
}

// GetAckCount mocks base method
func (m *SessionFSet) GetAckCount(arg0 *model.Session) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAckCount", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAckCount indicates an expected call of GetAckCount
func (mr *SessionFSetMockRecorder) GetAckCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAckCount", reflect.TypeOf((*SessionFSet)(nil).GetAckCount), arg0)
}

// GetTransactingTransaction mocks base method
func (m *SessionFSet) GetTransactingTransaction(arg0 *model.Session) ([]uint8, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactingTransaction", arg0)
	ret0, _ := ret[0].([]uint8)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactingTransaction indicates an expected call of GetTransactingTransaction
func (mr *SessionFSetMockRecorder) GetTransactingTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactingTransaction", reflect.TypeOf((*SessionFSet)(nil).GetTransactingTransaction), arg0)
}

// InitSessionInfo mocks base method
func (m *SessionFSet) InitSessionInfo(arg0 []uint8, arg1 []uip.Instruction, arg2 []*model.SessionAccount) (*model.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitSessionInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitSessionInfo indicates an expected call of InitSessionInfo
func (mr *SessionFSetMockRecorder) InitSessionInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitSessionInfo", reflect.TypeOf((*SessionFSet)(nil).InitSessionInfo), arg0, arg1, arg2)
}

// NotifyAttestation mocks base method
func (m *SessionFSet) NotifyAttestation(arg0 *model.Session, arg1 types.NSBInterface, arg2 uip.BlockChainInterface, arg3 uip.Attestation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyAttestation", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyAttestation indicates an expected call of NotifyAttestation
func (mr *SessionFSetMockRecorder) NotifyAttestation(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyAttestation", reflect.TypeOf((*SessionFSet)(nil).NotifyAttestation), arg0, arg1, arg2, arg3)
}

// ProcessAttestation mocks base method
func (m *SessionFSet) ProcessAttestation(arg0 types.NSBInterface, arg1 uip.BlockChainInterface, arg2 uip.Attestation) (interface{}, interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessAttestation", arg0, arg1, arg2)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ProcessAttestation indicates an expected call of ProcessAttestation
func (mr *SessionFSetMockRecorder) ProcessAttestation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessAttestation", reflect.TypeOf((*SessionFSet)(nil).ProcessAttestation), arg0, arg1, arg2)
}

// SyncFromISC mocks base method
func (m *SessionFSet) SyncFromISC() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncFromISC")
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncFromISC indicates an expected call of SyncFromISC
func (mr *SessionFSetMockRecorder) SyncFromISC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncFromISC", reflect.TypeOf((*SessionFSet)(nil).SyncFromISC))
}
