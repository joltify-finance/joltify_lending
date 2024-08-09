// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	pricestypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// PreparePricesKeeper is an autogenerated mock type for the PreparePricesKeeper type
type PreparePricesKeeper struct {
	mock.Mock
}

// GetValidMarketPriceUpdates provides a mock function with given fields: ctx
func (_m *PreparePricesKeeper) GetValidMarketPriceUpdates(ctx types.Context) *pricestypes.MsgUpdateMarketPrices {
	ret := _m.Called(ctx)

	var r0 *pricestypes.MsgUpdateMarketPrices
	if rf, ok := ret.Get(0).(func(types.Context) *pricestypes.MsgUpdateMarketPrices); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pricestypes.MsgUpdateMarketPrices)
		}
	}

	return r0
}

type mockConstructorTestingTNewPreparePricesKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewPreparePricesKeeper creates a new instance of PreparePricesKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPreparePricesKeeper(t mockConstructorTestingTNewPreparePricesKeeper) *PreparePricesKeeper {
	mock := &PreparePricesKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
