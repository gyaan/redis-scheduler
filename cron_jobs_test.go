package main

import (
	"gotest.tools/assert"
	"strings"
	"testing"
)

func Test_createInitialList(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"1 2 3 4 5 6 7 8 9 10",
		},
	}
	client := GetClient()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createInitialList("testItems")
			list := client.LRange("testItems", 0, 9)
			testList, _ := list.Result()
			assert.Equal(t, tt.name, strings.Join(testList, " "))
		})
	}
}

