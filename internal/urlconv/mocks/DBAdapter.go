// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DBAdapter is an autogenerated mock type for the DBAdapter type
type DBAdapter struct {
	mock.Mock
}

// CreateID provides a mock function with given fields:
func (_m *DBAdapter) CreateID() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveURL provides a mock function with given fields: short, long, exp
func (_m *DBAdapter) SaveURL(short string, long string, exp int64) error {
	ret := _m.Called(short, long, exp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(short, long, exp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
