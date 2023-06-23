// Code generated by mockery v2.28.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
	mock "github.com/stretchr/testify/mock"
)

// MockNewSquareRootServiceIF is an autogenerated mock type for the NewSquareRootServiceIF type
type MockNewSquareRootServiceIF struct {
	mock.Mock
}

type MockNewSquareRootServiceIF_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNewSquareRootServiceIF) EXPECT() *MockNewSquareRootServiceIF_Expecter {
	return &MockNewSquareRootServiceIF_Expecter{mock: &_m.Mock}
}

// SquareRoot provides a mock function with given fields: ctx, operation
func (_m *MockNewSquareRootServiceIF) SquareRoot(ctx context.Context, operation models.Operation) (*models.Record, *error) {
	ret := _m.Called(ctx, operation)

	var r0 *models.Record
	var r1 *error
	if rf, ok := ret.Get(0).(func(context.Context, models.Operation) (*models.Record, *error)); ok {
		return rf(ctx, operation)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Operation) *models.Record); ok {
		r0 = rf(ctx, operation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Record)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Operation) *error); ok {
		r1 = rf(ctx, operation)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*error)
		}
	}

	return r0, r1
}

// MockNewSquareRootServiceIF_SquareRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SquareRoot'
type MockNewSquareRootServiceIF_SquareRoot_Call struct {
	*mock.Call
}

// SquareRoot is a helper method to define mock.On call
//   - ctx context.Context
//   - operation models.Operation
func (_e *MockNewSquareRootServiceIF_Expecter) SquareRoot(ctx interface{}, operation interface{}) *MockNewSquareRootServiceIF_SquareRoot_Call {
	return &MockNewSquareRootServiceIF_SquareRoot_Call{Call: _e.mock.On("SquareRoot", ctx, operation)}
}

func (_c *MockNewSquareRootServiceIF_SquareRoot_Call) Run(run func(ctx context.Context, operation models.Operation)) *MockNewSquareRootServiceIF_SquareRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Operation))
	})
	return _c
}

func (_c *MockNewSquareRootServiceIF_SquareRoot_Call) Return(_a0 *models.Record, _a1 *error) *MockNewSquareRootServiceIF_SquareRoot_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNewSquareRootServiceIF_SquareRoot_Call) RunAndReturn(run func(context.Context, models.Operation) (*models.Record, *error)) *MockNewSquareRootServiceIF_SquareRoot_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockNewSquareRootServiceIF interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockNewSquareRootServiceIF creates a new instance of MockNewSquareRootServiceIF. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockNewSquareRootServiceIF(t mockConstructorTestingTNewMockNewSquareRootServiceIF) *MockNewSquareRootServiceIF {
	mock := &MockNewSquareRootServiceIF{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}