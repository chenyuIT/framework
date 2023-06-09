// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ResponseStatus is an autogenerated mock type for the ResponseStatus type
type ResponseStatus struct {
	mock.Mock
}

// Data provides a mock function with given fields: contentType, data
func (_m *ResponseStatus) Data(contentType string, data []byte) {
	_m.Called(contentType, data)
}

// Json provides a mock function with given fields: obj
func (_m *ResponseStatus) Json(obj interface{}) {
	_m.Called(obj)
}

// String provides a mock function with given fields: format, values
func (_m *ResponseStatus) String(format string, values ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, values...)
	_m.Called(_ca...)
}

// NewResponseStatus creates a new instance of ResponseStatus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResponseStatus(t interface {
	mock.TestingT
	Cleanup(func())
}) *ResponseStatus {
	mock := &ResponseStatus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
