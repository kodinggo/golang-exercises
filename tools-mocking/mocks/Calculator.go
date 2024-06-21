// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Calculator is an autogenerated mock type for the Calculator type
type Calculator struct {
	mock.Mock
}

// Count provides a mock function with given fields: n1, n2, op
func (_m *Calculator) Count(n1 int, n2 int, op rune) (int, error) {
	ret := _m.Called(n1, n2, op)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, rune) (int, error)); ok {
		return rf(n1, n2, op)
	}
	if rf, ok := ret.Get(0).(func(int, int, rune) int); ok {
		r0 = rf(n1, n2, op)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int, int, rune) error); ok {
		r1 = rf(n1, n2, op)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCalculator creates a new instance of Calculator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCalculator(t interface {
	mock.TestingT
	Cleanup(func())
}) *Calculator {
	mock := &Calculator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}