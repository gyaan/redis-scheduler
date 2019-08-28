package redisclient

import (
	"gotest.tools/assert"
	"reflect"
	"testing"

	"github.com/go-redis/redis"
)

//test get client
func TestGetClient(t *testing.T) {
	tests := []struct {
		name string
		want *redis.Client
	}{
		{"test1", GetClient()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// no singleton implemented so compare string of redis client obj
			if got := GetClient(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("GetClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

//test ping only for successful connection
func TestPing(t *testing.T) {
	ping := Ping()
	assert.Assert(t, ping, nil)
}
