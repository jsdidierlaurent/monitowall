// Code generated by mockery v1.0.0. DO NOT EDIT.

// If you want to rebuild this file, make mock

package mocks

import (
	pingdom "github.com/jsdidierlaurent/go-pingdom/pingdom"
	mock "github.com/stretchr/testify/mock"
)

// PingdomCheckAPI is an autogenerated mock type for the PingdomCheckAPI type
type PingdomCheckAPI struct {
	mock.Mock
}

// List provides a mock function with given fields: params
func (_m *PingdomCheckAPI) List(params ...map[string]string) ([]pingdom.CheckResponse, error) {
	_va := make([]interface{}, len(params))
	for _i := range params {
		_va[_i] = params[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []pingdom.CheckResponse
	if rf, ok := ret.Get(0).(func(...map[string]string) []pingdom.CheckResponse); ok {
		r0 = rf(params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pingdom.CheckResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...map[string]string) error); ok {
		r1 = rf(params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields: id
func (_m *PingdomCheckAPI) Read(id int) (*pingdom.CheckResponse, error) {
	ret := _m.Called(id)

	var r0 *pingdom.CheckResponse
	if rf, ok := ret.Get(0).(func(int) *pingdom.CheckResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pingdom.CheckResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
