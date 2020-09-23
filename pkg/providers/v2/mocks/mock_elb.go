// Code generated by MockGen. DO NOT EDIT.
// Source: loadbalancer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	elb "github.com/aws/aws-sdk-go/service/elb"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockELB is a mock of ELB interface
type MockELB struct {
	ctrl     *gomock.Controller
	recorder *MockELBMockRecorder
}

// MockELBMockRecorder is the mock recorder for MockELB
type MockELBMockRecorder struct {
	mock *MockELB
}

// NewMockELB creates a new mock instance
func NewMockELB(ctrl *gomock.Controller) *MockELB {
	mock := &MockELB{ctrl: ctrl}
	mock.recorder = &MockELBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockELB) EXPECT() *MockELBMockRecorder {
	return m.recorder
}

// CreateLoadBalancer mocks base method
func (m *MockELB) CreateLoadBalancer(arg0 *elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoadBalancer", arg0)
	ret0, _ := ret[0].(*elb.CreateLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoadBalancer indicates an expected call of CreateLoadBalancer
func (mr *MockELBMockRecorder) CreateLoadBalancer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoadBalancer", reflect.TypeOf((*MockELB)(nil).CreateLoadBalancer), arg0)
}

// DeleteLoadBalancer mocks base method
func (m *MockELB) DeleteLoadBalancer(arg0 *elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLoadBalancer", arg0)
	ret0, _ := ret[0].(*elb.DeleteLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLoadBalancer indicates an expected call of DeleteLoadBalancer
func (mr *MockELBMockRecorder) DeleteLoadBalancer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLoadBalancer", reflect.TypeOf((*MockELB)(nil).DeleteLoadBalancer), arg0)
}

// DescribeLoadBalancers mocks base method
func (m *MockELB) DescribeLoadBalancers(arg0 *elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeLoadBalancers", arg0)
	ret0, _ := ret[0].(*elb.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancers indicates an expected call of DescribeLoadBalancers
func (mr *MockELBMockRecorder) DescribeLoadBalancers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancers", reflect.TypeOf((*MockELB)(nil).DescribeLoadBalancers), arg0)
}

// RegisterInstancesWithLoadBalancer mocks base method
func (m *MockELB) RegisterInstancesWithLoadBalancer(arg0 *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterInstancesWithLoadBalancer", arg0)
	ret0, _ := ret[0].(*elb.RegisterInstancesWithLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterInstancesWithLoadBalancer indicates an expected call of RegisterInstancesWithLoadBalancer
func (mr *MockELBMockRecorder) RegisterInstancesWithLoadBalancer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInstancesWithLoadBalancer", reflect.TypeOf((*MockELB)(nil).RegisterInstancesWithLoadBalancer), arg0)
}

// DeregisterInstancesFromLoadBalancer mocks base method
func (m *MockELB) DeregisterInstancesFromLoadBalancer(arg0 *elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterInstancesFromLoadBalancer", arg0)
	ret0, _ := ret[0].(*elb.DeregisterInstancesFromLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeregisterInstancesFromLoadBalancer indicates an expected call of DeregisterInstancesFromLoadBalancer
func (mr *MockELBMockRecorder) DeregisterInstancesFromLoadBalancer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterInstancesFromLoadBalancer", reflect.TypeOf((*MockELB)(nil).DeregisterInstancesFromLoadBalancer), arg0)
}

// CreateLoadBalancerPolicy mocks base method
func (m *MockELB) CreateLoadBalancerPolicy(arg0 *elb.CreateLoadBalancerPolicyInput) (*elb.CreateLoadBalancerPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoadBalancerPolicy", arg0)
	ret0, _ := ret[0].(*elb.CreateLoadBalancerPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoadBalancerPolicy indicates an expected call of CreateLoadBalancerPolicy
func (mr *MockELBMockRecorder) CreateLoadBalancerPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoadBalancerPolicy", reflect.TypeOf((*MockELB)(nil).CreateLoadBalancerPolicy), arg0)
}

// SetLoadBalancerPoliciesOfListener mocks base method
func (m *MockELB) SetLoadBalancerPoliciesOfListener(input *elb.SetLoadBalancerPoliciesOfListenerInput) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLoadBalancerPoliciesOfListener", input)
	ret0, _ := ret[0].(*elb.SetLoadBalancerPoliciesOfListenerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetLoadBalancerPoliciesOfListener indicates an expected call of SetLoadBalancerPoliciesOfListener
func (mr *MockELBMockRecorder) SetLoadBalancerPoliciesOfListener(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLoadBalancerPoliciesOfListener", reflect.TypeOf((*MockELB)(nil).SetLoadBalancerPoliciesOfListener), input)
}

// DescribeLoadBalancerPolicies mocks base method
func (m *MockELB) DescribeLoadBalancerPolicies(input *elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeLoadBalancerPolicies", input)
	ret0, _ := ret[0].(*elb.DescribeLoadBalancerPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancerPolicies indicates an expected call of DescribeLoadBalancerPolicies
func (mr *MockELBMockRecorder) DescribeLoadBalancerPolicies(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancerPolicies", reflect.TypeOf((*MockELB)(nil).DescribeLoadBalancerPolicies), input)
}

// CreateLoadBalancerListeners mocks base method
func (m *MockELB) CreateLoadBalancerListeners(arg0 *elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoadBalancerListeners", arg0)
	ret0, _ := ret[0].(*elb.CreateLoadBalancerListenersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoadBalancerListeners indicates an expected call of CreateLoadBalancerListeners
func (mr *MockELBMockRecorder) CreateLoadBalancerListeners(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoadBalancerListeners", reflect.TypeOf((*MockELB)(nil).CreateLoadBalancerListeners), arg0)
}

// DeleteLoadBalancerListeners mocks base method
func (m *MockELB) DeleteLoadBalancerListeners(arg0 *elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLoadBalancerListeners", arg0)
	ret0, _ := ret[0].(*elb.DeleteLoadBalancerListenersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLoadBalancerListeners indicates an expected call of DeleteLoadBalancerListeners
func (mr *MockELBMockRecorder) DeleteLoadBalancerListeners(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLoadBalancerListeners", reflect.TypeOf((*MockELB)(nil).DeleteLoadBalancerListeners), arg0)
}

// ApplySecurityGroupsToLoadBalancer mocks base method
func (m *MockELB) ApplySecurityGroupsToLoadBalancer(arg0 *elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplySecurityGroupsToLoadBalancer", arg0)
	ret0, _ := ret[0].(*elb.ApplySecurityGroupsToLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplySecurityGroupsToLoadBalancer indicates an expected call of ApplySecurityGroupsToLoadBalancer
func (mr *MockELBMockRecorder) ApplySecurityGroupsToLoadBalancer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplySecurityGroupsToLoadBalancer", reflect.TypeOf((*MockELB)(nil).ApplySecurityGroupsToLoadBalancer), arg0)
}

// DescribeLoadBalancerAttributes mocks base method
func (m *MockELB) DescribeLoadBalancerAttributes(arg0 *elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeLoadBalancerAttributes", arg0)
	ret0, _ := ret[0].(*elb.DescribeLoadBalancerAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancerAttributes indicates an expected call of DescribeLoadBalancerAttributes
func (mr *MockELBMockRecorder) DescribeLoadBalancerAttributes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancerAttributes", reflect.TypeOf((*MockELB)(nil).DescribeLoadBalancerAttributes), arg0)
}

// ModifyLoadBalancerAttributes mocks base method
func (m *MockELB) ModifyLoadBalancerAttributes(arg0 *elb.ModifyLoadBalancerAttributesInput) (*elb.ModifyLoadBalancerAttributesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyLoadBalancerAttributes", arg0)
	ret0, _ := ret[0].(*elb.ModifyLoadBalancerAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyLoadBalancerAttributes indicates an expected call of ModifyLoadBalancerAttributes
func (mr *MockELBMockRecorder) ModifyLoadBalancerAttributes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyLoadBalancerAttributes", reflect.TypeOf((*MockELB)(nil).ModifyLoadBalancerAttributes), arg0)
}
