// Code generated by MockGen. DO NOT EDIT.
// Source: cloud.go

// Package mocks is a generated GoMock package.
package mocks

import (
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEC2 is a mock of EC2 interface
type MockEC2 struct {
	ctrl     *gomock.Controller
	recorder *MockEC2MockRecorder
}

// MockEC2MockRecorder is the mock recorder for MockEC2
type MockEC2MockRecorder struct {
	mock *MockEC2
}

// NewMockEC2 creates a new mock instance
func NewMockEC2(ctrl *gomock.Controller) *MockEC2 {
	mock := &MockEC2{ctrl: ctrl}
	mock.recorder = &MockEC2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEC2) EXPECT() *MockEC2MockRecorder {
	return m.recorder
}

// DescribeInstances mocks base method
func (m *MockEC2) DescribeInstances(request *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeInstances", request)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstances indicates an expected call of DescribeInstances
func (mr *MockEC2MockRecorder) DescribeInstances(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstances", reflect.TypeOf((*MockEC2)(nil).DescribeInstances), request)
}

// DescribeSecurityGroups mocks base method
func (m *MockEC2) DescribeSecurityGroups(request *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSecurityGroups", request)
	ret0, _ := ret[0].(*ec2.DescribeSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSecurityGroups indicates an expected call of DescribeSecurityGroups
func (mr *MockEC2MockRecorder) DescribeSecurityGroups(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecurityGroups", reflect.TypeOf((*MockEC2)(nil).DescribeSecurityGroups), request)
}

// DeleteSecurityGroup mocks base method
func (m *MockEC2) DeleteSecurityGroup(request *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecurityGroup", request)
	ret0, _ := ret[0].(*ec2.DeleteSecurityGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSecurityGroup indicates an expected call of DeleteSecurityGroup
func (mr *MockEC2MockRecorder) DeleteSecurityGroup(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecurityGroup", reflect.TypeOf((*MockEC2)(nil).DeleteSecurityGroup), request)
}

// AuthorizeSecurityGroupIngress mocks base method
func (m *MockEC2) AuthorizeSecurityGroupIngress(arg0 *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizeSecurityGroupIngress", arg0)
	ret0, _ := ret[0].(*ec2.AuthorizeSecurityGroupIngressOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorizeSecurityGroupIngress indicates an expected call of AuthorizeSecurityGroupIngress
func (mr *MockEC2MockRecorder) AuthorizeSecurityGroupIngress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizeSecurityGroupIngress", reflect.TypeOf((*MockEC2)(nil).AuthorizeSecurityGroupIngress), arg0)
}

// RevokeSecurityGroupIngress mocks base method
func (m *MockEC2) RevokeSecurityGroupIngress(arg0 *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeSecurityGroupIngress", arg0)
	ret0, _ := ret[0].(*ec2.RevokeSecurityGroupIngressOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeSecurityGroupIngress indicates an expected call of RevokeSecurityGroupIngress
func (mr *MockEC2MockRecorder) RevokeSecurityGroupIngress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeSecurityGroupIngress", reflect.TypeOf((*MockEC2)(nil).RevokeSecurityGroupIngress), arg0)
}

// MockEC2Metadata is a mock of EC2Metadata interface
type MockEC2Metadata struct {
	ctrl     *gomock.Controller
	recorder *MockEC2MetadataMockRecorder
}

// MockEC2MetadataMockRecorder is the mock recorder for MockEC2Metadata
type MockEC2MetadataMockRecorder struct {
	mock *MockEC2Metadata
}

// NewMockEC2Metadata creates a new mock instance
func NewMockEC2Metadata(ctrl *gomock.Controller) *MockEC2Metadata {
	mock := &MockEC2Metadata{ctrl: ctrl}
	mock.recorder = &MockEC2MetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEC2Metadata) EXPECT() *MockEC2MetadataMockRecorder {
	return m.recorder
}

// GetMetadata mocks base method
func (m *MockEC2Metadata) GetMetadata(path string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata", path)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata
func (mr *MockEC2MetadataMockRecorder) GetMetadata(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockEC2Metadata)(nil).GetMetadata), path)
}
