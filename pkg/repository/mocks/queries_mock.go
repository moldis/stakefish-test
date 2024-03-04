// Code generated by mockery v2.23.2. DO NOT EDIT.

package repository

import (
	context "context"

	model "code.stakefish.test/service/ip_validator/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// MockQueries is an autogenerated mock type for the Queries type
type MockQueries struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, in
func (_m *MockQueries) Create(ctx context.Context, in *model.Query) error {
	ret := _m.Called(ctx, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Query) error); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: ctx, skip, limit
func (_m *MockQueries) List(ctx context.Context, skip int64, limit int64) ([]model.Query, error) {
	ret := _m.Called(ctx, skip, limit)

	var r0 []model.Query
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) ([]model.Query, error)); ok {
		return rf(ctx, skip, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) []model.Query); ok {
		r0 = rf(ctx, skip, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Query)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, skip, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockQueries interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockQueries creates a new instance of MockQueries. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockQueries(t mockConstructorTestingTNewMockQueries) *MockQueries {
	mock := &MockQueries{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}