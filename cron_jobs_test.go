package main

import (
	"fmt"
	"testing"
)

func Test_createInitialList(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"1 2 3 4 5 6 7 8 9",
		},
	}
	client := GetClient()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createInitialList()
			list := client.LRange("items", 0,9)

			fmt.Println(list)

		})
	}
}

func Test_task(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task()
		})
	}
}

func Test_runJobs(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runJobs()
		})
	}
}
