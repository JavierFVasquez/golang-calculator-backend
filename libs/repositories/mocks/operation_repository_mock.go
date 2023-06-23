// Code generated by mockery v2.28.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
	mock "github.com/stretchr/testify/mock"
)

// MockOperationRepositoryIF is an autogenerated mock type for the OperationRepositoryIF type
type MockOperationRepositoryIF struct {
	mock.Mock
}

type MockOperationRepositoryIF_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOperationRepositoryIF) EXPECT() *MockOperationRepositoryIF_Expecter {
	return &MockOperationRepositoryIF_Expecter{mock: &_m.Mock}
}

// GetOperation provides a mock function with given fields: ctx, id
func (_m *MockOperationRepositoryIF) GetOperation(ctx context.Context, id models.Operator) (*models.Operation, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Operator) (*models.Operation, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Operator) *models.Operation); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Operator) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOperationRepositoryIF_GetOperation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOperation'
type MockOperationRepositoryIF_GetOperation_Call struct {
	*mock.Call
}

// GetOperation is a helper method to define mock.On call
//   - ctx context.Context
//   - id models.Operator
func (_e *MockOperationRepositoryIF_Expecter) GetOperation(ctx interface{}, id interface{}) *MockOperationRepositoryIF_GetOperation_Call {
	return &MockOperationRepositoryIF_GetOperation_Call{Call: _e.mock.On("GetOperation", ctx, id)}
}

func (_c *MockOperationRepositoryIF_GetOperation_Call) Run(run func(ctx context.Context, id models.Operator)) *MockOperationRepositoryIF_GetOperation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Operator))
	})
	return _c
}

func (_c *MockOperationRepositoryIF_GetOperation_Call) Return(_a0 *models.Operation, _a1 error) *MockOperationRepositoryIF_GetOperation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOperationRepositoryIF_GetOperation_Call) RunAndReturn(run func(context.Context, models.Operator) (*models.Operation, error)) *MockOperationRepositoryIF_GetOperation_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockOperationRepositoryIF interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockOperationRepositoryIF creates a new instance of MockOperationRepositoryIF. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockOperationRepositoryIF(t mockConstructorTestingTNewMockOperationRepositoryIF) *MockOperationRepositoryIF {
	mock := &MockOperationRepositoryIF{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}