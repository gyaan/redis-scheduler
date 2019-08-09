package main

import (
	"reflect"
	"testing"

	"github.com/go-redis/redis"
)

func TestGetClient(t *testing.T) {
	tests := []struct {
		name string
		want *redis.Client
	}{
		{"test1",GetClient()},
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
