// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	user "21-api/features/user"

	mock "github.com/stretchr/testify/mock"
)

// UserModel is an autogenerated mock type for the UserModel type
type UserModel struct {
	mock.Mock
}

// GetUserByHP provides a mock function with given fields: hp
func (_m *UserModel) GetUserByHP(hp string) (user.User, error) {
	ret := _m.Called(hp)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByHP")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (user.User, error)); ok {
		return rf(hp)
	}
	if rf, ok := ret.Get(0).(func(string) user.User); ok {
		r0 = rf(hp)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUser provides a mock function with given fields: newData
func (_m *UserModel) InsertUser(newData user.User) error {
	ret := _m.Called(newData)

	if len(ret) == 0 {
		panic("no return value specified for InsertUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(user.User) error); ok {
		r0 = rf(newData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: hp
func (_m *UserModel) Login(hp string) (user.User, error) {
	ret := _m.Called(hp)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (user.User, error)); ok {
		return rf(hp)
	}
	if rf, ok := ret.Get(0).(func(string) user.User); ok {
		r0 = rf(hp)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: hp, data
func (_m *UserModel) UpdateUser(hp string, data user.User) error {
	ret := _m.Called(hp, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, user.User) error); ok {
		r0 = rf(hp, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserModel creates a new instance of UserModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserModel(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserModel {
	mock := &UserModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
