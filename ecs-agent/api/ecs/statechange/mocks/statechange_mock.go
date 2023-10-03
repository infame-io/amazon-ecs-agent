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
// Source: github.com/aws/amazon-ecs-agent/ecs-agent/api/ecs/statechange (interfaces: ContainerMetadataGetter,TaskMetadataGetter)

// Package mock_statechange is a generated GoMock package.
package mock_statechange

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockContainerMetadataGetter is a mock of ContainerMetadataGetter interface.
type MockContainerMetadataGetter struct {
	ctrl     *gomock.Controller
	recorder *MockContainerMetadataGetterMockRecorder
}

// MockContainerMetadataGetterMockRecorder is the mock recorder for MockContainerMetadataGetter.
type MockContainerMetadataGetterMockRecorder struct {
	mock *MockContainerMetadataGetter
}

// NewMockContainerMetadataGetter creates a new mock instance.
func NewMockContainerMetadataGetter(ctrl *gomock.Controller) *MockContainerMetadataGetter {
	mock := &MockContainerMetadataGetter{ctrl: ctrl}
	mock.recorder = &MockContainerMetadataGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContainerMetadataGetter) EXPECT() *MockContainerMetadataGetterMockRecorder {
	return m.recorder
}

// GetContainerIsEssential mocks base method.
func (m *MockContainerMetadataGetter) GetContainerIsEssential() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerIsEssential")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetContainerIsEssential indicates an expected call of GetContainerIsEssential.
func (mr *MockContainerMetadataGetterMockRecorder) GetContainerIsEssential() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerIsEssential", reflect.TypeOf((*MockContainerMetadataGetter)(nil).GetContainerIsEssential))
}

// GetContainerIsNil mocks base method.
func (m *MockContainerMetadataGetter) GetContainerIsNil() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerIsNil")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetContainerIsNil indicates an expected call of GetContainerIsNil.
func (mr *MockContainerMetadataGetterMockRecorder) GetContainerIsNil() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerIsNil", reflect.TypeOf((*MockContainerMetadataGetter)(nil).GetContainerIsNil))
}

// GetContainerRuntimeID mocks base method.
func (m *MockContainerMetadataGetter) GetContainerRuntimeID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerRuntimeID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetContainerRuntimeID indicates an expected call of GetContainerRuntimeID.
func (mr *MockContainerMetadataGetterMockRecorder) GetContainerRuntimeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerRuntimeID", reflect.TypeOf((*MockContainerMetadataGetter)(nil).GetContainerRuntimeID))
}

// GetContainerSentStatusString mocks base method.
func (m *MockContainerMetadataGetter) GetContainerSentStatusString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerSentStatusString")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetContainerSentStatusString indicates an expected call of GetContainerSentStatusString.
func (mr *MockContainerMetadataGetterMockRecorder) GetContainerSentStatusString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerSentStatusString", reflect.TypeOf((*MockContainerMetadataGetter)(nil).GetContainerSentStatusString))
}

// MockTaskMetadataGetter is a mock of TaskMetadataGetter interface.
type MockTaskMetadataGetter struct {
	ctrl     *gomock.Controller
	recorder *MockTaskMetadataGetterMockRecorder
}

// MockTaskMetadataGetterMockRecorder is the mock recorder for MockTaskMetadataGetter.
type MockTaskMetadataGetterMockRecorder struct {
	mock *MockTaskMetadataGetter
}

// NewMockTaskMetadataGetter creates a new mock instance.
func NewMockTaskMetadataGetter(ctrl *gomock.Controller) *MockTaskMetadataGetter {
	mock := &MockTaskMetadataGetter{ctrl: ctrl}
	mock.recorder = &MockTaskMetadataGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskMetadataGetter) EXPECT() *MockTaskMetadataGetterMockRecorder {
	return m.recorder
}

// GetTaskExecutionStoppedAt mocks base method.
func (m *MockTaskMetadataGetter) GetTaskExecutionStoppedAt() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskExecutionStoppedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTaskExecutionStoppedAt indicates an expected call of GetTaskExecutionStoppedAt.
func (mr *MockTaskMetadataGetterMockRecorder) GetTaskExecutionStoppedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskExecutionStoppedAt", reflect.TypeOf((*MockTaskMetadataGetter)(nil).GetTaskExecutionStoppedAt))
}

// GetTaskIsNil mocks base method.
func (m *MockTaskMetadataGetter) GetTaskIsNil() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskIsNil")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetTaskIsNil indicates an expected call of GetTaskIsNil.
func (mr *MockTaskMetadataGetterMockRecorder) GetTaskIsNil() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskIsNil", reflect.TypeOf((*MockTaskMetadataGetter)(nil).GetTaskIsNil))
}

// GetTaskPullStartedAt mocks base method.
func (m *MockTaskMetadataGetter) GetTaskPullStartedAt() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskPullStartedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTaskPullStartedAt indicates an expected call of GetTaskPullStartedAt.
func (mr *MockTaskMetadataGetterMockRecorder) GetTaskPullStartedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskPullStartedAt", reflect.TypeOf((*MockTaskMetadataGetter)(nil).GetTaskPullStartedAt))
}

// GetTaskPullStoppedAt mocks base method.
func (m *MockTaskMetadataGetter) GetTaskPullStoppedAt() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskPullStoppedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTaskPullStoppedAt indicates an expected call of GetTaskPullStoppedAt.
func (mr *MockTaskMetadataGetterMockRecorder) GetTaskPullStoppedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskPullStoppedAt", reflect.TypeOf((*MockTaskMetadataGetter)(nil).GetTaskPullStoppedAt))
}

// GetTaskSentStatusString mocks base method.
func (m *MockTaskMetadataGetter) GetTaskSentStatusString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskSentStatusString")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTaskSentStatusString indicates an expected call of GetTaskSentStatusString.
func (mr *MockTaskMetadataGetterMockRecorder) GetTaskSentStatusString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskSentStatusString", reflect.TypeOf((*MockTaskMetadataGetter)(nil).GetTaskSentStatusString))
}
