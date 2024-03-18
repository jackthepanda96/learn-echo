// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// TodoController is an autogenerated mock type for the TodoController type
type TodoController struct {
	mock.Mock
}

// Add provides a mock function with given fields:
func (_m *TodoController) Add() echo.HandlerFunc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// NewTodoController creates a new instance of TodoController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTodoController(t interface {
	mock.TestingT
	Cleanup(func())
}) *TodoController {
	mock := &TodoController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}