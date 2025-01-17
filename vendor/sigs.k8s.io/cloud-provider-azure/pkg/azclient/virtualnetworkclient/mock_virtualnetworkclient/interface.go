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
//

// Code generated by MockGen. DO NOT EDIT.
// Source: virtualnetworkclient/interface.go
//
// Generated by this command:
//
//	mockgen -package mock_virtualnetworkclient -source virtualnetworkclient/interface.go -typed -write_generate_directive -copyright_file ../../hack/boilerplate/boilerplate.generatego.txt
//

// Package mock_virtualnetworkclient is a generated GoMock package.
package mock_virtualnetworkclient

import (
	context "context"
	reflect "reflect"

	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	gomock "go.uber.org/mock/gomock"
)

//go:generate mockgen -package mock_virtualnetworkclient -source virtualnetworkclient/interface.go -typed -write_generate_directive -copyright_file ../../hack/boilerplate/boilerplate.generatego.txt

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
	isgomock struct{}
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

// CheckIPAddressAvailability mocks base method.
func (m *MockInterface) CheckIPAddressAvailability(ctx context.Context, resourceGroupName, virtualNetworkName, ipAddress string) (*armnetwork.IPAddressAvailabilityResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIPAddressAvailability", ctx, resourceGroupName, virtualNetworkName, ipAddress)
	ret0, _ := ret[0].(*armnetwork.IPAddressAvailabilityResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIPAddressAvailability indicates an expected call of CheckIPAddressAvailability.
func (mr *MockInterfaceMockRecorder) CheckIPAddressAvailability(ctx, resourceGroupName, virtualNetworkName, ipAddress any) *MockInterfaceCheckIPAddressAvailabilityCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIPAddressAvailability", reflect.TypeOf((*MockInterface)(nil).CheckIPAddressAvailability), ctx, resourceGroupName, virtualNetworkName, ipAddress)
	return &MockInterfaceCheckIPAddressAvailabilityCall{Call: call}
}

// MockInterfaceCheckIPAddressAvailabilityCall wrap *gomock.Call
type MockInterfaceCheckIPAddressAvailabilityCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceCheckIPAddressAvailabilityCall) Return(arg0 *armnetwork.IPAddressAvailabilityResult, arg1 error) *MockInterfaceCheckIPAddressAvailabilityCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceCheckIPAddressAvailabilityCall) Do(f func(context.Context, string, string, string) (*armnetwork.IPAddressAvailabilityResult, error)) *MockInterfaceCheckIPAddressAvailabilityCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceCheckIPAddressAvailabilityCall) DoAndReturn(f func(context.Context, string, string, string) (*armnetwork.IPAddressAvailabilityResult, error)) *MockInterfaceCheckIPAddressAvailabilityCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreateOrUpdate mocks base method.
func (m *MockInterface) CreateOrUpdate(ctx context.Context, resourceGroupName, resourceName string, resourceParam armnetwork.VirtualNetwork) (*armnetwork.VirtualNetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, resourceName, resourceParam)
	ret0, _ := ret[0].(*armnetwork.VirtualNetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockInterfaceMockRecorder) CreateOrUpdate(ctx, resourceGroupName, resourceName, resourceParam any) *MockInterfaceCreateOrUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockInterface)(nil).CreateOrUpdate), ctx, resourceGroupName, resourceName, resourceParam)
	return &MockInterfaceCreateOrUpdateCall{Call: call}
}

// MockInterfaceCreateOrUpdateCall wrap *gomock.Call
type MockInterfaceCreateOrUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceCreateOrUpdateCall) Return(arg0 *armnetwork.VirtualNetwork, arg1 error) *MockInterfaceCreateOrUpdateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceCreateOrUpdateCall) Do(f func(context.Context, string, string, armnetwork.VirtualNetwork) (*armnetwork.VirtualNetwork, error)) *MockInterfaceCreateOrUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceCreateOrUpdateCall) DoAndReturn(f func(context.Context, string, string, armnetwork.VirtualNetwork) (*armnetwork.VirtualNetwork, error)) *MockInterfaceCreateOrUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Delete mocks base method.
func (m *MockInterface) Delete(ctx context.Context, resourceGroupName, resourceName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, resourceGroupName, resourceName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockInterfaceMockRecorder) Delete(ctx, resourceGroupName, resourceName any) *MockInterfaceDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockInterface)(nil).Delete), ctx, resourceGroupName, resourceName)
	return &MockInterfaceDeleteCall{Call: call}
}

// MockInterfaceDeleteCall wrap *gomock.Call
type MockInterfaceDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceDeleteCall) Return(arg0 error) *MockInterfaceDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceDeleteCall) Do(f func(context.Context, string, string) error) *MockInterfaceDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceDeleteCall) DoAndReturn(f func(context.Context, string, string) error) *MockInterfaceDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockInterface) Get(ctx context.Context, resourceGroupName, resourceName string, expand *string) (*armnetwork.VirtualNetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, resourceName, expand)
	ret0, _ := ret[0].(*armnetwork.VirtualNetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(ctx, resourceGroupName, resourceName, expand any) *MockInterfaceGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, resourceGroupName, resourceName, expand)
	return &MockInterfaceGetCall{Call: call}
}

// MockInterfaceGetCall wrap *gomock.Call
type MockInterfaceGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceGetCall) Return(result *armnetwork.VirtualNetwork, rerr error) *MockInterfaceGetCall {
	c.Call = c.Call.Return(result, rerr)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceGetCall) Do(f func(context.Context, string, string, *string) (*armnetwork.VirtualNetwork, error)) *MockInterfaceGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceGetCall) DoAndReturn(f func(context.Context, string, string, *string) (*armnetwork.VirtualNetwork, error)) *MockInterfaceGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockInterface) List(ctx context.Context, resourceGroupName string) ([]*armnetwork.VirtualNetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, resourceGroupName)
	ret0, _ := ret[0].([]*armnetwork.VirtualNetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockInterfaceMockRecorder) List(ctx, resourceGroupName any) *MockInterfaceListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), ctx, resourceGroupName)
	return &MockInterfaceListCall{Call: call}
}

// MockInterfaceListCall wrap *gomock.Call
type MockInterfaceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceListCall) Return(result []*armnetwork.VirtualNetwork, rerr error) *MockInterfaceListCall {
	c.Call = c.Call.Return(result, rerr)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceListCall) Do(f func(context.Context, string) ([]*armnetwork.VirtualNetwork, error)) *MockInterfaceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceListCall) DoAndReturn(f func(context.Context, string) ([]*armnetwork.VirtualNetwork, error)) *MockInterfaceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
