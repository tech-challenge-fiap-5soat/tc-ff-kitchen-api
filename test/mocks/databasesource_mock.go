// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockDatabaseSource is an autogenerated mock type for the DatabaseSource type
type MockDatabaseSource struct {
	mock.Mock
}

type MockDatabaseSource_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDatabaseSource) EXPECT() *MockDatabaseSource_Expecter {
	return &MockDatabaseSource_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: identifier
func (_m *MockDatabaseSource) Delete(identifier string) (interface{}, error) {
	ret := _m.Called(identifier)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (interface{}, error)); ok {
		return rf(identifier)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(identifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatabaseSource_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockDatabaseSource_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - identifier string
func (_e *MockDatabaseSource_Expecter) Delete(identifier interface{}) *MockDatabaseSource_Delete_Call {
	return &MockDatabaseSource_Delete_Call{Call: _e.mock.On("Delete", identifier)}
}

func (_c *MockDatabaseSource_Delete_Call) Run(run func(identifier string)) *MockDatabaseSource_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDatabaseSource_Delete_Call) Return(id interface{}, err error) *MockDatabaseSource_Delete_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockDatabaseSource_Delete_Call) RunAndReturn(run func(string) (interface{}, error)) *MockDatabaseSource_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields: fieldName, fieldValue
func (_m *MockDatabaseSource) FindAll(fieldName string, fieldValue string) ([]interface{}, error) {
	ret := _m.Called(fieldName, fieldValue)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]interface{}, error)); ok {
		return rf(fieldName, fieldValue)
	}
	if rf, ok := ret.Get(0).(func(string, string) []interface{}); ok {
		r0 = rf(fieldName, fieldValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(fieldName, fieldValue)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatabaseSource_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockDatabaseSource_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
//   - fieldName string
//   - fieldValue string
func (_e *MockDatabaseSource_Expecter) FindAll(fieldName interface{}, fieldValue interface{}) *MockDatabaseSource_FindAll_Call {
	return &MockDatabaseSource_FindAll_Call{Call: _e.mock.On("FindAll", fieldName, fieldValue)}
}

func (_c *MockDatabaseSource_FindAll_Call) Run(run func(fieldName string, fieldValue string)) *MockDatabaseSource_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockDatabaseSource_FindAll_Call) Return(_a0 []interface{}, _a1 error) *MockDatabaseSource_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDatabaseSource_FindAll_Call) RunAndReturn(run func(string, string) ([]interface{}, error)) *MockDatabaseSource_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindOne provides a mock function with given fields: key, value
func (_m *MockDatabaseSource) FindOne(key string, value string) (interface{}, error) {
	ret := _m.Called(key, value)

	if len(ret) == 0 {
		panic("no return value specified for FindOne")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (interface{}, error)); ok {
		return rf(key, value)
	}
	if rf, ok := ret.Get(0).(func(string, string) interface{}); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(key, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatabaseSource_FindOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOne'
type MockDatabaseSource_FindOne_Call struct {
	*mock.Call
}

// FindOne is a helper method to define mock.On call
//   - key string
//   - value string
func (_e *MockDatabaseSource_Expecter) FindOne(key interface{}, value interface{}) *MockDatabaseSource_FindOne_Call {
	return &MockDatabaseSource_FindOne_Call{Call: _e.mock.On("FindOne", key, value)}
}

func (_c *MockDatabaseSource_FindOne_Call) Run(run func(key string, value string)) *MockDatabaseSource_FindOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockDatabaseSource_FindOne_Call) Return(_a0 interface{}, _a1 error) *MockDatabaseSource_FindOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDatabaseSource_FindOne_Call) RunAndReturn(run func(string, string) (interface{}, error)) *MockDatabaseSource_FindOne_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: data
func (_m *MockDatabaseSource) Save(data interface{}) (interface{}, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatabaseSource_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockDatabaseSource_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - data interface{}
func (_e *MockDatabaseSource_Expecter) Save(data interface{}) *MockDatabaseSource_Save_Call {
	return &MockDatabaseSource_Save_Call{Call: _e.mock.On("Save", data)}
}

func (_c *MockDatabaseSource_Save_Call) Run(run func(data interface{})) *MockDatabaseSource_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockDatabaseSource_Save_Call) Return(id interface{}, err error) *MockDatabaseSource_Save_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockDatabaseSource_Save_Call) RunAndReturn(run func(interface{}) (interface{}, error)) *MockDatabaseSource_Save_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: identifier, data
func (_m *MockDatabaseSource) Update(identifier string, data interface{}) (interface{}, error) {
	ret := _m.Called(identifier, data)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, interface{}) (interface{}, error)); ok {
		return rf(identifier, data)
	}
	if rf, ok := ret.Get(0).(func(string, interface{}) interface{}); ok {
		r0 = rf(identifier, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(identifier, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatabaseSource_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockDatabaseSource_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - identifier string
//   - data interface{}
func (_e *MockDatabaseSource_Expecter) Update(identifier interface{}, data interface{}) *MockDatabaseSource_Update_Call {
	return &MockDatabaseSource_Update_Call{Call: _e.mock.On("Update", identifier, data)}
}

func (_c *MockDatabaseSource_Update_Call) Run(run func(identifier string, data interface{})) *MockDatabaseSource_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *MockDatabaseSource_Update_Call) Return(id interface{}, err error) *MockDatabaseSource_Update_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *MockDatabaseSource_Update_Call) RunAndReturn(run func(string, interface{}) (interface{}, error)) *MockDatabaseSource_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDatabaseSource creates a new instance of MockDatabaseSource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDatabaseSource(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDatabaseSource {
	mock := &MockDatabaseSource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}