// Code generated by mockery v2.28.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
	mock "github.com/stretchr/testify/mock"
)

// MockRecordServiceIF is an autogenerated mock type for the RecordServiceIF type
type MockRecordServiceIF struct {
	mock.Mock
}

type MockRecordServiceIF_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRecordServiceIF) EXPECT() *MockRecordServiceIF_Expecter {
	return &MockRecordServiceIF_Expecter{mock: &_m.Mock}
}

// DeleteRecord provides a mock function with given fields: ctx, recordId
func (_m *MockRecordServiceIF) DeleteRecord(ctx context.Context, recordId *string) (*models.Record, *error) {
	ret := _m.Called(ctx, recordId)

	var r0 *models.Record
	var r1 *error
	if rf, ok := ret.Get(0).(func(context.Context, *string) (*models.Record, *error)); ok {
		return rf(ctx, recordId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string) *models.Record); ok {
		r0 = rf(ctx, recordId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Record)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string) *error); ok {
		r1 = rf(ctx, recordId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*error)
		}
	}

	return r0, r1
}

// MockRecordServiceIF_DeleteRecord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRecord'
type MockRecordServiceIF_DeleteRecord_Call struct {
	*mock.Call
}

// DeleteRecord is a helper method to define mock.On call
//   - ctx context.Context
//   - recordId *string
func (_e *MockRecordServiceIF_Expecter) DeleteRecord(ctx interface{}, recordId interface{}) *MockRecordServiceIF_DeleteRecord_Call {
	return &MockRecordServiceIF_DeleteRecord_Call{Call: _e.mock.On("DeleteRecord", ctx, recordId)}
}

func (_c *MockRecordServiceIF_DeleteRecord_Call) Run(run func(ctx context.Context, recordId *string)) *MockRecordServiceIF_DeleteRecord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*string))
	})
	return _c
}

func (_c *MockRecordServiceIF_DeleteRecord_Call) Return(_a0 *models.Record, _a1 *error) *MockRecordServiceIF_DeleteRecord_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecordServiceIF_DeleteRecord_Call) RunAndReturn(run func(context.Context, *string) (*models.Record, *error)) *MockRecordServiceIF_DeleteRecord_Call {
	_c.Call.Return(run)
	return _c
}

// GetRecordList provides a mock function with given fields: ctx, options
func (_m *MockRecordServiceIF) GetRecordList(ctx context.Context, options *models.PaginationParams) (*models.PaginatedResponse[models.Record], *error) {
	ret := _m.Called(ctx, options)

	var r0 *models.PaginatedResponse[models.Record]
	var r1 *error
	if rf, ok := ret.Get(0).(func(context.Context, *models.PaginationParams) (*models.PaginatedResponse[models.Record], *error)); ok {
		return rf(ctx, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.PaginationParams) *models.PaginatedResponse[models.Record]); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PaginatedResponse[models.Record])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.PaginationParams) *error); ok {
		r1 = rf(ctx, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*error)
		}
	}

	return r0, r1
}

// MockRecordServiceIF_GetRecordList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRecordList'
type MockRecordServiceIF_GetRecordList_Call struct {
	*mock.Call
}

// GetRecordList is a helper method to define mock.On call
//   - ctx context.Context
//   - options *models.PaginationParams
func (_e *MockRecordServiceIF_Expecter) GetRecordList(ctx interface{}, options interface{}) *MockRecordServiceIF_GetRecordList_Call {
	return &MockRecordServiceIF_GetRecordList_Call{Call: _e.mock.On("GetRecordList", ctx, options)}
}

func (_c *MockRecordServiceIF_GetRecordList_Call) Run(run func(ctx context.Context, options *models.PaginationParams)) *MockRecordServiceIF_GetRecordList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.PaginationParams))
	})
	return _c
}

func (_c *MockRecordServiceIF_GetRecordList_Call) Return(_a0 *models.PaginatedResponse[models.Record], _a1 *error) *MockRecordServiceIF_GetRecordList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecordServiceIF_GetRecordList_Call) RunAndReturn(run func(context.Context, *models.PaginationParams) (*models.PaginatedResponse[models.Record], *error)) *MockRecordServiceIF_GetRecordList_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockRecordServiceIF interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRecordServiceIF creates a new instance of MockRecordServiceIF. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRecordServiceIF(t mockConstructorTestingTNewMockRecordServiceIF) *MockRecordServiceIF {
	mock := &MockRecordServiceIF{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
