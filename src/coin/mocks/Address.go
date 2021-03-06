// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import core "github.com/fibercrypto/fibercryptowallet/src/core"
import mock "github.com/stretchr/testify/mock"

// Address is an autogenerated mock type for the Address type
type Address struct {
	mock.Mock
}

// Bytes provides a mock function with given fields:
func (_m *Address) Bytes() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// Checksum provides a mock function with given fields:
func (_m *Address) Checksum() core.Checksum {
	ret := _m.Called()

	var r0 core.Checksum
	if rf, ok := ret.Get(0).(func() core.Checksum); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Checksum)
		}
	}

	return r0
}

// GetCryptoAccount provides a mock function with given fields:
func (_m *Address) GetCryptoAccount() core.CryptoAccount {
	ret := _m.Called()

	var r0 core.CryptoAccount
	if rf, ok := ret.Get(0).(func() core.CryptoAccount); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.CryptoAccount)
		}
	}

	return r0
}

// IsBip32 provides a mock function with given fields:
func (_m *Address) IsBip32() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Null provides a mock function with given fields:
func (_m *Address) Null() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Address) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Verify provides a mock function with given fields: _a0
func (_m *Address) Verify(_a0 core.PubKey) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.PubKey) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
