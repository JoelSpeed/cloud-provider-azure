// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetclient (interfaces: Interface)

// Package mock_virtualmachinescalesetclient is a generated GoMock package.
package mock_virtualmachinescalesetclient

import (
	context "context"
	reflect "reflect"

	armcompute "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	gomock "github.com/golang/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockInterface) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 armcompute.VirtualMachineScaleSet) (*armcompute.VirtualMachineScaleSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*armcompute.VirtualMachineScaleSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockInterfaceMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockInterface)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// Delete mocks base method.
func (m *MockInterface) Delete(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockInterface)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockInterface) Get(arg0 context.Context, arg1, arg2 string) (*armcompute.VirtualMachineScaleSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*armcompute.VirtualMachineScaleSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockInterface) List(arg0 context.Context, arg1 string) ([]*armcompute.VirtualMachineScaleSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*armcompute.VirtualMachineScaleSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), arg0, arg1)
}
