// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	s3 "github.com/observiq/bindplane-agent/receiver/awss3rehydrationreceiver/internal/aws"
	mock "github.com/stretchr/testify/mock"
)

// MockS3Client is an autogenerated mock type for the S3Client type
type MockS3Client struct {
	mock.Mock
}

type MockS3Client_Expecter struct {
	mock *mock.Mock
}

func (_m *MockS3Client) EXPECT() *MockS3Client_Expecter {
	return &MockS3Client_Expecter{mock: &_m.Mock}
}

// DeleteObjects provides a mock function with given fields: ctx, bucket, keys
func (_m *MockS3Client) DeleteObjects(ctx context.Context, bucket string, keys []string) error {
	ret := _m.Called(ctx, bucket, keys)

	if len(ret) == 0 {
		panic("no return value specified for DeleteObjects")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) error); ok {
		r0 = rf(ctx, bucket, keys)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockS3Client_DeleteObjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteObjects'
type MockS3Client_DeleteObjects_Call struct {
	*mock.Call
}

// DeleteObjects is a helper method to define mock.On call
//   - ctx context.Context
//   - bucket string
//   - keys []string
func (_e *MockS3Client_Expecter) DeleteObjects(ctx interface{}, bucket interface{}, keys interface{}) *MockS3Client_DeleteObjects_Call {
	return &MockS3Client_DeleteObjects_Call{Call: _e.mock.On("DeleteObjects", ctx, bucket, keys)}
}

func (_c *MockS3Client_DeleteObjects_Call) Run(run func(ctx context.Context, bucket string, keys []string)) *MockS3Client_DeleteObjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *MockS3Client_DeleteObjects_Call) Return(_a0 error) *MockS3Client_DeleteObjects_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockS3Client_DeleteObjects_Call) RunAndReturn(run func(context.Context, string, []string) error) *MockS3Client_DeleteObjects_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadObject provides a mock function with given fields: ctx, bucket, key, buf
func (_m *MockS3Client) DownloadObject(ctx context.Context, bucket string, key string, buf []byte) (int64, error) {
	ret := _m.Called(ctx, bucket, key, buf)

	if len(ret) == 0 {
		panic("no return value specified for DownloadObject")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []byte) (int64, error)); ok {
		return rf(ctx, bucket, key, buf)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []byte) int64); ok {
		r0 = rf(ctx, bucket, key, buf)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, []byte) error); ok {
		r1 = rf(ctx, bucket, key, buf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockS3Client_DownloadObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadObject'
type MockS3Client_DownloadObject_Call struct {
	*mock.Call
}

// DownloadObject is a helper method to define mock.On call
//   - ctx context.Context
//   - bucket string
//   - key string
//   - buf []byte
func (_e *MockS3Client_Expecter) DownloadObject(ctx interface{}, bucket interface{}, key interface{}, buf interface{}) *MockS3Client_DownloadObject_Call {
	return &MockS3Client_DownloadObject_Call{Call: _e.mock.On("DownloadObject", ctx, bucket, key, buf)}
}

func (_c *MockS3Client_DownloadObject_Call) Run(run func(ctx context.Context, bucket string, key string, buf []byte)) *MockS3Client_DownloadObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].([]byte))
	})
	return _c
}

func (_c *MockS3Client_DownloadObject_Call) Return(_a0 int64, _a1 error) *MockS3Client_DownloadObject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockS3Client_DownloadObject_Call) RunAndReturn(run func(context.Context, string, string, []byte) (int64, error)) *MockS3Client_DownloadObject_Call {
	_c.Call.Return(run)
	return _c
}

// ListObjects provides a mock function with given fields: ctx, bucket, prefix, continuationToken
func (_m *MockS3Client) ListObjects(ctx context.Context, bucket string, prefix *string, continuationToken *string) ([]*s3.ObjectInfo, *string, error) {
	ret := _m.Called(ctx, bucket, prefix, continuationToken)

	if len(ret) == 0 {
		panic("no return value specified for ListObjects")
	}

	var r0 []*s3.ObjectInfo
	var r1 *string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *string, *string) ([]*s3.ObjectInfo, *string, error)); ok {
		return rf(ctx, bucket, prefix, continuationToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *string, *string) []*s3.ObjectInfo); ok {
		r0 = rf(ctx, bucket, prefix, continuationToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*s3.ObjectInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *string, *string) *string); ok {
		r1 = rf(ctx, bucket, prefix, continuationToken)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*string)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, *string, *string) error); ok {
		r2 = rf(ctx, bucket, prefix, continuationToken)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockS3Client_ListObjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListObjects'
type MockS3Client_ListObjects_Call struct {
	*mock.Call
}

// ListObjects is a helper method to define mock.On call
//   - ctx context.Context
//   - bucket string
//   - prefix *string
//   - continuationToken *string
func (_e *MockS3Client_Expecter) ListObjects(ctx interface{}, bucket interface{}, prefix interface{}, continuationToken interface{}) *MockS3Client_ListObjects_Call {
	return &MockS3Client_ListObjects_Call{Call: _e.mock.On("ListObjects", ctx, bucket, prefix, continuationToken)}
}

func (_c *MockS3Client_ListObjects_Call) Run(run func(ctx context.Context, bucket string, prefix *string, continuationToken *string)) *MockS3Client_ListObjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*string), args[3].(*string))
	})
	return _c
}

func (_c *MockS3Client_ListObjects_Call) Return(_a0 []*s3.ObjectInfo, _a1 *string, _a2 error) *MockS3Client_ListObjects_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockS3Client_ListObjects_Call) RunAndReturn(run func(context.Context, string, *string, *string) ([]*s3.ObjectInfo, *string, error)) *MockS3Client_ListObjects_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockS3Client creates a new instance of MockS3Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockS3Client(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockS3Client {
	mock := &MockS3Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
