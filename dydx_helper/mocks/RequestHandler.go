// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RequestHandler is an autogenerated mock type for the RequestHandler type
type RequestHandler struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, url
func (_m *RequestHandler) Get(ctx context.Context, url string) (*http.Response, error) {
	ret := _m.Called(ctx, url)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*http.Response, error)); ok {
		return rf(ctx, url)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *http.Response); ok {
		r0 = rf(ctx, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRequestHandler creates a new instance of RequestHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequestHandler(t interface {
	mock.TestingT
	Cleanup(func())
},
) *RequestHandler {
	mock := &RequestHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
