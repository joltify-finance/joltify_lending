// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/joltify-finance/joltify_lending/daemons/bridge/api"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// BridgeServiceClient is an autogenerated mock type for the BridgeServiceClient type
type BridgeServiceClient struct {
	mock.Mock
}

// AddBridgeEvents provides a mock function with given fields: ctx, in, opts
func (_m *BridgeServiceClient) AddBridgeEvents(ctx context.Context, in *api.AddBridgeEventsRequest, opts ...grpc.CallOption) (*api.AddBridgeEventsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddBridgeEvents")
	}

	var r0 *api.AddBridgeEventsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.AddBridgeEventsRequest, ...grpc.CallOption) (*api.AddBridgeEventsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.AddBridgeEventsRequest, ...grpc.CallOption) *api.AddBridgeEventsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.AddBridgeEventsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.AddBridgeEventsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBridgeServiceClient creates a new instance of BridgeServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBridgeServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
},
) *BridgeServiceClient {
	mock := &BridgeServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}