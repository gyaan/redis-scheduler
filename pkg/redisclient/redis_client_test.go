package redisclient

import (
	"gotest.tools/assert"
	"redis-scheduler/config"
	"testing"
)

func TestGetClient(t *testing.T) {
	conf := &config.Config{}
	conf.Redis.Addr = "localhost:6379"
	conf.Redis.Password = ""
	conf.Redis.DB = 0
	Init(conf)

	client := GetClient()
	assert.Assert(t, client != nil)
}

func TestPing(t *testing.T) {
	conf := &config.Config{}
	conf.Redis.Addr = "localhost:6379"
	conf.Redis.Password = ""
	conf.Redis.DB = 0
	Init(conf)

	ping := Ping()
	assert.Assert(t, ping == nil)
}