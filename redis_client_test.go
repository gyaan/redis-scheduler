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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
