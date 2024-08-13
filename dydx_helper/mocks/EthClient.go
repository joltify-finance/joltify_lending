// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	coretypes "github.com/ethereum/go-ethereum/core/types"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"
)

// EthClient is an autogenerated mock type for the EthClient type
type EthClient struct {
	mock.Mock
}

// ChainID provides a mock function with given fields: ctx
func (_m *EthClient) ChainID(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ChainID")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterLogs provides a mock function with given fields: ctx, q
func (_m *EthClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]coretypes.Log, error) {
	ret := _m.Called(ctx, q)

	if len(ret) == 0 {
		panic("no return value specified for FilterLogs")
	}

	var r0 []coretypes.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) ([]coretypes.Log, error)); ok {
		return rf(ctx, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) []coretypes.Log); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coretypes.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEthClient creates a new instance of EthClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEthClient(t interface {
	mock.TestingT
	Cleanup(func())
},
) *EthClient {
	mock := &EthClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
