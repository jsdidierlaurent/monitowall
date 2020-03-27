// Code generated by mockery v1.0.0. DO NOT EDIT.

// If you want to rebuild this file, make mock

package mocks

import (
	context "context"

	http "net/http"

	mock "github.com/stretchr/testify/mock"

	travis "github.com/shuheiktgw/go-travis"
)

// TravisCI is an autogenerated mock type for the TravisCI type
type TravisCI struct {
	mock.Mock
}

// ListByRepoSlug provides a mock function with given fields: ctx, repoSlug, opt
func (_m *TravisCI) ListByRepoSlug(ctx context.Context, repoSlug string, opt *travis.BuildsByRepoOption) ([]*travis.Build, *http.Response, error) {
	ret := _m.Called(ctx, repoSlug, opt)

	var r0 []*travis.Build
	if rf, ok := ret.Get(0).(func(context.Context, string, *travis.BuildsByRepoOption) []*travis.Build); ok {
		r0 = rf(ctx, repoSlug, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*travis.Build)
		}
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, *travis.BuildsByRepoOption) *http.Response); ok {
		r1 = rf(ctx, repoSlug, opt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, *travis.BuildsByRepoOption) error); ok {
		r2 = rf(ctx, repoSlug, opt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
