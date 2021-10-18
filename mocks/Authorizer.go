// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Authorizer is an autogenerated mock type for the Authorizer type
type Authorizer struct {
	mock.Mock
}

// IsAuthorized provides a mock function with given fields: r
func (_m *Authorizer) IsAuthorized(r *http.Request) bool {
	ret := _m.Called(r)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*http.Request) bool); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}