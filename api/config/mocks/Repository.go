// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/monitoror/monitoror/api/config/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetConfigFromPath provides a mock function with given fields: baseDir, filePath
func (_m *Repository) GetConfigFromPath(baseDir string, filePath string) (*models.Config, error) {
	ret := _m.Called(baseDir, filePath)

	var r0 *models.Config
	if rf, ok := ret.Get(0).(func(string, string) *models.Config); ok {
		r0 = rf(baseDir, filePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Config)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(baseDir, filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfigFromURL provides a mock function with given fields: url
func (_m *Repository) GetConfigFromURL(url string) (*models.Config, error) {
	ret := _m.Called(url)

	var r0 *models.Config
	if rf, ok := ret.Get(0).(func(string) *models.Config); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Config)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
