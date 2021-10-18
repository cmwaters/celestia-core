// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	conn "github.com/tendermint/tendermint/internal/p2p/conn"

	mock "github.com/stretchr/testify/mock"

	p2p "github.com/tendermint/tendermint/internal/p2p"
)

// Transport is an autogenerated mock type for the Transport type
type Transport struct {
	mock.Mock
}

// Accept provides a mock function with given fields:
func (_m *Transport) Accept() (p2p.Connection, error) {
	ret := _m.Called()

	var r0 p2p.Connection
	if rf, ok := ret.Get(0).(func() p2p.Connection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.Connection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddChannelDescriptors provides a mock function with given fields: _a0
func (_m *Transport) AddChannelDescriptors(_a0 []*conn.ChannelDescriptor) {
	_m.Called(_a0)
}

// Close provides a mock function with given fields:
func (_m *Transport) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dial provides a mock function with given fields: _a0, _a1
func (_m *Transport) Dial(_a0 context.Context, _a1 p2p.Endpoint) (p2p.Connection, error) {
	ret := _m.Called(_a0, _a1)

	var r0 p2p.Connection
	if rf, ok := ret.Get(0).(func(context.Context, p2p.Endpoint) p2p.Connection); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.Connection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, p2p.Endpoint) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Endpoints provides a mock function with given fields:
func (_m *Transport) Endpoints() []p2p.Endpoint {
	ret := _m.Called()

	var r0 []p2p.Endpoint
	if rf, ok := ret.Get(0).(func() []p2p.Endpoint); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]p2p.Endpoint)
		}
	}

	return r0
}

// Protocols provides a mock function with given fields:
func (_m *Transport) Protocols() []p2p.Protocol {
	ret := _m.Called()

	var r0 []p2p.Protocol
	if rf, ok := ret.Get(0).(func() []p2p.Protocol); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]p2p.Protocol)
		}
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Transport) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
