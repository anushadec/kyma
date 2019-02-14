// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"
import mock "github.com/stretchr/testify/mock"

// Cleaner is an autogenerated mock type for the Cleaner type
type Cleaner struct {
	mock.Mock
}

// Clean provides a mock function with given fields: ctx, bucket, objectPrefix
func (_m *Cleaner) Clean(ctx context.Context, bucket string, objectPrefix string) error {
	ret := _m.Called(ctx, bucket, objectPrefix)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, bucket, objectPrefix)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}