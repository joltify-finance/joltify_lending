// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	proto "github.com/cosmos/gogoproto/proto"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// MsgRouter is an autogenerated mock type for the MsgRouter type
type MsgRouter struct {
	mock.Mock
}

// Handler provides a mock function with given fields: msg
func (_m *MsgRouter) Handler(msg proto.Message) func(types.Context, proto.Message) (*types.Result, error) {
	ret := _m.Called(msg)

	if len(ret) == 0 {
		panic("no return value specified for Handler")
	}

	var r0 func(types.Context, proto.Message) (*types.Result, error)
	if rf, ok := ret.Get(0).(func(proto.Message) func(types.Context, proto.Message) (*types.Result, error)); ok {
		r0 = rf(msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(types.Context, proto.Message) (*types.Result, error))
		}
	}

	return r0
}

// NewMsgRouter creates a new instance of MsgRouter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMsgRouter(t interface {
	mock.TestingT
	Cleanup(func())
},
) *MsgRouter {
	mock := &MsgRouter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
