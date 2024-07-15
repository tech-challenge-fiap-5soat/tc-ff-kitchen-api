// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockOrderApi is an autogenerated mock type for the OrderApi type
type MockOrderApi struct {
	mock.Mock
}

type MockOrderApi_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOrderApi) EXPECT() *MockOrderApi_Expecter {
	return &MockOrderApi_Expecter{mock: &_m.Mock}
}

// FinishOrder provides a mock function with given fields: orderId
func (_m *MockOrderApi) FinishOrder(orderId string) error {
	ret := _m.Called(orderId)

	if len(ret) == 0 {
		panic("no return value specified for FinishOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(orderId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOrderApi_FinishOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FinishOrder'
type MockOrderApi_FinishOrder_Call struct {
	*mock.Call
}

// FinishOrder is a helper method to define mock.On call
//   - orderId string
func (_e *MockOrderApi_Expecter) FinishOrder(orderId interface{}) *MockOrderApi_FinishOrder_Call {
	return &MockOrderApi_FinishOrder_Call{Call: _e.mock.On("FinishOrder", orderId)}
}

func (_c *MockOrderApi_FinishOrder_Call) Run(run func(orderId string)) *MockOrderApi_FinishOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockOrderApi_FinishOrder_Call) Return(_a0 error) *MockOrderApi_FinishOrder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOrderApi_FinishOrder_Call) RunAndReturn(run func(string) error) *MockOrderApi_FinishOrder_Call {
	_c.Call.Return(run)
	return _c
}

// ReleaseOrder provides a mock function with given fields: orderId
func (_m *MockOrderApi) ReleaseOrder(orderId string) error {
	ret := _m.Called(orderId)

	if len(ret) == 0 {
		panic("no return value specified for ReleaseOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(orderId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOrderApi_ReleaseOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseOrder'
type MockOrderApi_ReleaseOrder_Call struct {
	*mock.Call
}

// ReleaseOrder is a helper method to define mock.On call
//   - orderId string
func (_e *MockOrderApi_Expecter) ReleaseOrder(orderId interface{}) *MockOrderApi_ReleaseOrder_Call {
	return &MockOrderApi_ReleaseOrder_Call{Call: _e.mock.On("ReleaseOrder", orderId)}
}

func (_c *MockOrderApi_ReleaseOrder_Call) Run(run func(orderId string)) *MockOrderApi_ReleaseOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockOrderApi_ReleaseOrder_Call) Return(_a0 error) *MockOrderApi_ReleaseOrder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOrderApi_ReleaseOrder_Call) RunAndReturn(run func(string) error) *MockOrderApi_ReleaseOrder_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOrderApi creates a new instance of MockOrderApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOrderApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOrderApi {
	mock := &MockOrderApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
