// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./utils_windows.go

// Package mock_networkutils is a generated GoMock package.
package mock_networkutils

import (
	context "context"
	net "net"
	reflect "reflect"
	time "time"

	networkutils "github.com/aws/amazon-ecs-agent/agent/eni/networkutils"
	gomock "github.com/golang/mock/gomock"
)

// MockNetworkUtils is a mock of NetworkUtils interface
type MockNetworkUtils struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkUtilsMockRecorder
}

// MockNetworkUtilsMockRecorder is the mock recorder for MockNetworkUtils
type MockNetworkUtilsMockRecorder struct {
	mock *MockNetworkUtils
}

// NewMockNetworkUtils creates a new mock instance
func NewMockNetworkUtils(ctrl *gomock.Controller) *MockNetworkUtils {
	mock := &MockNetworkUtils{ctrl: ctrl}
	mock.recorder = &MockNetworkUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkUtils) EXPECT() *MockNetworkUtilsMockRecorder {
	return m.recorder
}

// GetInterfaceMACByIndex mocks base method
func (m *MockNetworkUtils) GetInterfaceMACByIndex(arg0 int, arg1 context.Context, arg2 time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInterfaceMACByIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInterfaceMACByIndex indicates an expected call of GetInterfaceMACByIndex
func (mr *MockNetworkUtilsMockRecorder) GetInterfaceMACByIndex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInterfaceMACByIndex", reflect.TypeOf((*MockNetworkUtils)(nil).GetInterfaceMACByIndex), arg0, arg1, arg2)
}

// GetAllNetworkInterfaces mocks base method
func (m *MockNetworkUtils) GetAllNetworkInterfaces() ([]net.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllNetworkInterfaces")
	ret0, _ := ret[0].([]net.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllNetworkInterfaces indicates an expected call of GetAllNetworkInterfaces
func (mr *MockNetworkUtilsMockRecorder) GetAllNetworkInterfaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllNetworkInterfaces", reflect.TypeOf((*MockNetworkUtils)(nil).GetAllNetworkInterfaces))
}

// GetDNSServerAddressList mocks base method
func (m *MockNetworkUtils) GetDNSServerAddressList(macAddress string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDNSServerAddressList", macAddress)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDNSServerAddressList indicates an expected call of GetDNSServerAddressList
func (mr *MockNetworkUtilsMockRecorder) GetDNSServerAddressList(macAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDNSServerAddressList", reflect.TypeOf((*MockNetworkUtils)(nil).GetDNSServerAddressList), macAddress)
}

// ConvertInterfaceAliasToLUID mocks base method
func (m *MockNetworkUtils) ConvertInterfaceAliasToLUID(interfaceAlias string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertInterfaceAliasToLUID", interfaceAlias)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertInterfaceAliasToLUID indicates an expected call of ConvertInterfaceAliasToLUID
func (mr *MockNetworkUtilsMockRecorder) ConvertInterfaceAliasToLUID(interfaceAlias interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertInterfaceAliasToLUID", reflect.TypeOf((*MockNetworkUtils)(nil).ConvertInterfaceAliasToLUID), interfaceAlias)
}

// GetMIBIfEntryFromLUID mocks base method
func (m *MockNetworkUtils) GetMIBIfEntryFromLUID(ifaceLUID uint64) (*networkutils.MibIfRow2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMIBIfEntryFromLUID", ifaceLUID)
	ret0, _ := ret[0].(*networkutils.MibIfRow2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMIBIfEntryFromLUID indicates an expected call of GetMIBIfEntryFromLUID
func (mr *MockNetworkUtilsMockRecorder) GetMIBIfEntryFromLUID(ifaceLUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMIBIfEntryFromLUID", reflect.TypeOf((*MockNetworkUtils)(nil).GetMIBIfEntryFromLUID), ifaceLUID)
}
