// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/internal/operators (interfaces: API)

// Package operators is a generated GoMock package.
package operators

import (
	gomock "github.com/golang/mock/gomock"
	common "github.com/openshift/assisted-service/internal/common"
	models "github.com/openshift/assisted-service/models"
	reflect "reflect"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// AnyOperatorEnabled mocks base method
func (m *MockAPI) AnyOperatorEnabled(arg0 *common.Cluster) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnyOperatorEnabled", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AnyOperatorEnabled indicates an expected call of AnyOperatorEnabled
func (mr *MockAPIMockRecorder) AnyOperatorEnabled(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnyOperatorEnabled", reflect.TypeOf((*MockAPI)(nil).AnyOperatorEnabled), arg0)
}

// GenerateManifests mocks base method
func (m *MockAPI) GenerateManifests(arg0 *common.Cluster) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateManifests", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateManifests indicates an expected call of GenerateManifests
func (mr *MockAPIMockRecorder) GenerateManifests(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateManifests", reflect.TypeOf((*MockAPI)(nil).GenerateManifests), arg0)
}

// GetMonitoredOperatorsList mocks base method
func (m *MockAPI) GetMonitoredOperatorsList() []*models.MonitoredOperator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMonitoredOperatorsList")
	ret0, _ := ret[0].([]*models.MonitoredOperator)
	return ret0
}

// GetMonitoredOperatorsList indicates an expected call of GetMonitoredOperatorsList
func (mr *MockAPIMockRecorder) GetMonitoredOperatorsList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMonitoredOperatorsList", reflect.TypeOf((*MockAPI)(nil).GetMonitoredOperatorsList))
}

// GetOperatorByName mocks base method
func (m *MockAPI) GetOperatorByName(arg0 string) (*models.MonitoredOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorByName", arg0)
	ret0, _ := ret[0].(*models.MonitoredOperator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorByName indicates an expected call of GetOperatorByName
func (mr *MockAPIMockRecorder) GetOperatorByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorByName", reflect.TypeOf((*MockAPI)(nil).GetOperatorByName), arg0)
}

// GetOperatorStatusInfo mocks base method
func (m *MockAPI) GetOperatorStatusInfo(arg0 *common.Cluster, arg1 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorStatusInfo", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOperatorStatusInfo indicates an expected call of GetOperatorStatusInfo
func (mr *MockAPIMockRecorder) GetOperatorStatusInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorStatusInfo", reflect.TypeOf((*MockAPI)(nil).GetOperatorStatusInfo), arg0, arg1)
}

// GetOperatorsByType mocks base method
func (m *MockAPI) GetOperatorsByType(arg0 models.OperatorType) []*models.MonitoredOperator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsByType", arg0)
	ret0, _ := ret[0].([]*models.MonitoredOperator)
	return ret0
}

// GetOperatorsByType indicates an expected call of GetOperatorsByType
func (mr *MockAPIMockRecorder) GetOperatorsByType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsByType", reflect.TypeOf((*MockAPI)(nil).GetOperatorsByType), arg0)
}

// ValidateOCSRequirements mocks base method
func (m *MockAPI) ValidateOCSRequirements(arg0 *common.Cluster) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateOCSRequirements", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// ValidateOCSRequirements indicates an expected call of ValidateOCSRequirements
func (mr *MockAPIMockRecorder) ValidateOCSRequirements(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateOCSRequirements", reflect.TypeOf((*MockAPI)(nil).ValidateOCSRequirements), arg0)
}
