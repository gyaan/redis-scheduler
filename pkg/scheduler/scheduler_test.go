package scheduler

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"redis-scheduler/pkg/redisclient"
)

// MockRedisClient is a mock implementation of the redis client
type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) RPush(key string, args ...interface{}) *redis.IntCmd {
	// This is a bit of a hack to get the arguments to the mock function
	// to be of the correct type.
	allArgs := append([]interface{}{key}, args...)
	ret := m.Called(allArgs...)
	return ret.Get(0).(*redis.IntCmd)
}

func (m *MockRedisClient) LRange(key string, start, stop int64) *redis.StringSliceCmd {
	ret := m.Called(key, start, stop)
	return ret.Get(0).(*redis.StringSliceCmd)
}

func (m *MockRedisClient) Del(keys ...string) *redis.IntCmd {
	// This is a bit of a hack to get the arguments to the mock function
	// to be of the correct type.
	var allArgs []interface{}
	for _, key := range keys {
		allArgs = append(allArgs, key)
	}
	ret := m.Called(allArgs...)
	return ret.Get(0).(*redis.IntCmd)
}

func (m *MockRedisClient) LPop(key string) *redis.StringCmd {
	ret := m.Called(key)
	return ret.Get(0).(*redis.StringCmd)
}

func (m *MockRedisClient) Ping() *redis.StatusCmd {
	ret := m.Called()
	return ret.Get(0).(*redis.StatusCmd)
}

func (m *MockRedisClient) Close() error {
	ret := m.Called()
	return ret.Error(0)
}

func TestCreateInitialList(t *testing.T) {
	client := new(MockRedisClient)
	redisclient.SetClient(client)

	for i := 1; i <= 10; i++ {
		client.On("RPush", "test_list", i).Return(redis.NewIntResult(int64(i), nil))
	}
	client.On("Close").Return(nil)

	err := createInitialList("test_list")
	assert.NoError(t, err)

	client.AssertExpectations(t)
}

func TestTask(t *testing.T) {
	client := new(MockRedisClient)
	redisclient.SetClient(client)

	client.On("LRange", "circular_list_for_update", int64(0), int64(-1)).Return(redis.NewStringSliceResult([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, nil))
	client.On("LRange", "circular_list_for_update", int64(0), int64(2)).Return(redis.NewStringSliceResult([]string{"1", "2", "3"}, nil))
	client.On("Del", "current_list_for_update").Return(redis.NewIntResult(1, nil))
	client.On("RPush", "current_list_for_update", "1").Return(redis.NewIntResult(1, nil))
	client.On("RPush", "current_list_for_update", "2").Return(redis.NewIntResult(2, nil))
	client.On("RPush", "current_list_for_update", "3").Return(redis.NewIntResult(3, nil))
	client.On("LRange", "current_list_for_update", int64(0), int64(2)).Return(redis.NewStringSliceResult([]string{"1", "2", "3"}, nil))
	client.On("LPop", "circular_list_for_update").Return(redis.NewStringResult("1", nil))
	client.On("LPop", "circular_list_for_update").Return(redis.NewStringResult("2", nil))
	client.On("LPop", "circular_list_for_update").Return(redis.NewStringResult("3", nil))
	client.On("RPush", "circular_list_for_update", "1").Return(redis.NewIntResult(1, nil))
	client.On("RPush", "circular_list_for_update", "2").Return(redis.NewIntResult(2, nil))
	client.On("RPush", "circular_list_for_update", "3").Return(redis.NewIntResult(3, nil))
	client.On("Close").Return(nil)

	err := task()
	assert.NoError(t, err)

	client.AssertExpectations(t)
}
