// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	http "github.com/chenyuIT/framework/contracts/http"
	mock "github.com/stretchr/testify/mock"
)

// RateLimiter is an autogenerated mock type for the RateLimiter type
type RateLimiter struct {
	mock.Mock
}

// For provides a mock function with given fields: name, callback
func (_m *RateLimiter) For(name string, callback func(http.Context) http.Limit) {
	_m.Called(name, callback)
}

// ForWithLimits provides a mock function with given fields: name, callback
func (_m *RateLimiter) ForWithLimits(name string, callback func(http.Context) []http.Limit) {
	_m.Called(name, callback)
}

// Limiter provides a mock function with given fields: name
func (_m *RateLimiter) Limiter(name string) func(http.Context) []http.Limit {
	ret := _m.Called(name)

	var r0 func(http.Context) []http.Limit
	if rf, ok := ret.Get(0).(func(string) func(http.Context) []http.Limit); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(http.Context) []http.Limit)
		}
	}

	return r0
}

type mockConstructorTestingTNewRateLimiter interface {
	mock.TestingT
	Cleanup(func())
}

// NewRateLimiter creates a new instance of RateLimiter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRateLimiter(t mockConstructorTestingTNewRateLimiter) *RateLimiter {
	mock := &RateLimiter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}