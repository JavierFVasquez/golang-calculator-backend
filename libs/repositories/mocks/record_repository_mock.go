// Code generated by mockery v2.28.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/JavierFVasquez/golang-calculator-backend/libs/models"
	mock "github.com/stretchr/testify/mock"
)

// MockRecordRepositoryIF is an autogenerated mock type for the RecordRepositoryIF type
type MockRecordRepositoryIF struct {
	mock.Mock
}

type MockRecordRepositoryIF_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRecordRepositoryIF) EXPECT() *MockRecordRepositoryIF_Expecter {
	return &MockRecordRepositoryIF_Expecter{mock: &_m.Mock}
}

// CreateNewRecord provides a mock function with given fields: ctx, operation, newUserBalance
func (_m *MockRecordRepositoryIF) CreateNewRecord(ctx context.Context, operation models.Operation, newUserBalance float64) (*models.Record, error) {
	ret := _m.Called(ctx, operation, newUserBalance)

	var r0 *models.Record
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Operation, float64) (*models.Record, error)); ok {
		return rf(ctx, operation, newUserBalance)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Operation, float64) *models.Record); ok {
		r0 = rf(ctx, operation, newUserBalance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Record)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Operation, float64) error); ok {
		r1 = rf(ctx, operation, newUserBalance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecordRepositoryIF_CreateNewRecord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateNewRecord'
type MockRecordRepositoryIF_CreateNewRecord_Call struct {
	*mock.Call
}

// CreateNewRecord is a helper method to define mock.On call
//   - ctx context.Context
//   - operation models.Operation
//   - newUserBalance float64
func (_e *MockRecordRepositoryIF_Expecter) CreateNewRecord(ctx interface{}, operation interface{}, newUserBalance interface{}) *MockRecordRepositoryIF_CreateNewRecord_Call {
	return &MockRecordRepositoryIF_CreateNewRecord_Call{Call: _e.mock.On("CreateNewRecord", ctx, operation, newUserBalance)}
}

func (_c *MockRecordRepositoryIF_CreateNewRecord_Call) Run(run func(ctx context.Context, operation models.Operation, newUserBalance float64)) *MockRecordRepositoryIF_CreateNewRecord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Operation), args[2].(float64))
	})
	return _c
}

func (_c *MockRecordRepositoryIF_CreateNewRecord_Call) Return(_a0 *models.Record, _a1 error) *MockRecordRepositoryIF_CreateNewRecord_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecordRepositoryIF_CreateNewRecord_Call) RunAndReturn(run func(context.Context, models.Operation, float64) (*models.Record, error)) *MockRecordRepositoryIF_CreateNewRecord_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRecord provides a mock function with given fields: ctx, recordId
func (_m *MockRecordRepositoryIF) DeleteRecord(ctx context.Context, recordId *string) (*models.Record, error) {
	ret := _m.Called(ctx, recordId)

	var r0 *models.Record
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *string) (*models.Record, error)); ok {
		return rf(ctx, recordId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string) *models.Record); ok {
		r0 = rf(ctx, recordId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Record)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(ctx, recordId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecordRepositoryIF_DeleteRecord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRecord'
type MockRecordRepositoryIF_DeleteRecord_Call struct {
	*mock.Call
}

// DeleteRecord is a helper method to define mock.On call
//   - ctx context.Context
//   - recordId *string
func (_e *MockRecordRepositoryIF_Expecter) DeleteRecord(ctx interface{}, recordId interface{}) *MockRecordRepositoryIF_DeleteRecord_Call {
	return &MockRecordRepositoryIF_DeleteRecord_Call{Call: _e.mock.On("DeleteRecord", ctx, recordId)}
}

func (_c *MockRecordRepositoryIF_DeleteRecord_Call) Run(run func(ctx context.Context, recordId *string)) *MockRecordRepositoryIF_DeleteRecord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*string))
	})
	return _c
}

func (_c *MockRecordRepositoryIF_DeleteRecord_Call) Return(_a0 *models.Record, _a1 error) *MockRecordRepositoryIF_DeleteRecord_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecordRepositoryIF_DeleteRecord_Call) RunAndReturn(run func(context.Context, *string) (*models.Record, error)) *MockRecordRepositoryIF_DeleteRecord_Call {
	_c.Call.Return(run)
	return _c
}

// GetRecordsByUserId provides a mock function with given fields: ctx, id, options
func (_m *MockRecordRepositoryIF) GetRecordsByUserId(ctx context.Context, id string, options *models.PaginationParams) (*models.PaginatedResponse[models.Record], error) {
	ret := _m.Called(ctx, id, options)

	var r0 *models.PaginatedResponse[models.Record]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *models.PaginationParams) (*models.PaginatedResponse[models.Record], error)); ok {
		return rf(ctx, id, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *models.PaginationParams) *models.PaginatedResponse[models.Record]); ok {
		r0 = rf(ctx, id, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PaginatedResponse[models.Record])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *models.PaginationParams) error); ok {
		r1 = rf(ctx, id, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecordRepositoryIF_GetRecordsByUserId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRecordsByUserId'
type MockRecordRepositoryIF_GetRecordsByUserId_Call struct {
	*mock.Call
}

// GetRecordsByUserId is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - options *models.PaginationParams
func (_e *MockRecordRepositoryIF_Expecter) GetRecordsByUserId(ctx interface{}, id interface{}, options interface{}) *MockRecordRepositoryIF_GetRecordsByUserId_Call {
	return &MockRecordRepositoryIF_GetRecordsByUserId_Call{Call: _e.mock.On("GetRecordsByUserId", ctx, id, options)}
}

func (_c *MockRecordRepositoryIF_GetRecordsByUserId_Call) Run(run func(ctx context.Context, id string, options *models.PaginationParams)) *MockRecordRepositoryIF_GetRecordsByUserId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*models.PaginationParams))
	})
	return _c
}

func (_c *MockRecordRepositoryIF_GetRecordsByUserId_Call) Return(_a0 *models.PaginatedResponse[models.Record], _a1 error) *MockRecordRepositoryIF_GetRecordsByUserId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecordRepositoryIF_GetRecordsByUserId_Call) RunAndReturn(run func(context.Context, string, *models.PaginationParams) (*models.PaginatedResponse[models.Record], error)) *MockRecordRepositoryIF_GetRecordsByUserId_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserMostRecentRecord provides a mock function with given fields: ctx
func (_m *MockRecordRepositoryIF) GetUserMostRecentRecord(ctx context.Context) (*models.Record, error) {
	ret := _m.Called(ctx)

	var r0 *models.Record
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*models.Record, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *models.Record); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Record)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecordRepositoryIF_GetUserMostRecentRecord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserMostRecentRecord'
type MockRecordRepositoryIF_GetUserMostRecentRecord_Call struct {
	*mock.Call
}

// GetUserMostRecentRecord is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRecordRepositoryIF_Expecter) GetUserMostRecentRecord(ctx interface{}) *MockRecordRepositoryIF_GetUserMostRecentRecord_Call {
	return &MockRecordRepositoryIF_GetUserMostRecentRecord_Call{Call: _e.mock.On("GetUserMostRecentRecord", ctx)}
}

func (_c *MockRecordRepositoryIF_GetUserMostRecentRecord_Call) Run(run func(ctx context.Context)) *MockRecordRepositoryIF_GetUserMostRecentRecord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRecordRepositoryIF_GetUserMostRecentRecord_Call) Return(_a0 *models.Record, _a1 error) *MockRecordRepositoryIF_GetUserMostRecentRecord_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecordRepositoryIF_GetUserMostRecentRecord_Call) RunAndReturn(run func(context.Context) (*models.Record, error)) *MockRecordRepositoryIF_GetUserMostRecentRecord_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockRecordRepositoryIF interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRecordRepositoryIF creates a new instance of MockRecordRepositoryIF. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRecordRepositoryIF(t mockConstructorTestingTNewMockRecordRepositoryIF) *MockRecordRepositoryIF {
	mock := &MockRecordRepositoryIF{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
