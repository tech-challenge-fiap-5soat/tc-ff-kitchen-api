// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockPublisher is an autogenerated mock type for the Publisher type
type MockPublisher struct {
	mock.Mock
}

type MockPublisher_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPublisher) EXPECT() *MockPublisher_Expecter {
	return &MockPublisher_Expecter{mock: &_m.Mock}
}

// PublishMessage provides a mock function with given fields: queueName, message
func (_m *MockPublisher) PublishMessage(queueName string, message string) error {
	ret := _m.Called(queueName, message)

	if len(ret) == 0 {
		panic("no return value specified for PublishMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(queueName, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPublisher_PublishMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PublishMessage'
type MockPublisher_PublishMessage_Call struct {
	*mock.Call
}

// PublishMessage is a helper method to define mock.On call
//   - queueName string
//   - message string
func (_e *MockPublisher_Expecter) PublishMessage(queueName interface{}, message interface{}) *MockPublisher_PublishMessage_Call {
	return &MockPublisher_PublishMessage_Call{Call: _e.mock.On("PublishMessage", queueName, message)}
}

func (_c *MockPublisher_PublishMessage_Call) Run(run func(queueName string, message string)) *MockPublisher_PublishMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockPublisher_PublishMessage_Call) Return(_a0 error) *MockPublisher_PublishMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublisher_PublishMessage_Call) RunAndReturn(run func(string, string) error) *MockPublisher_PublishMessage_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPublisher creates a new instance of MockPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPublisher(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPublisher {
	mock := &MockPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}