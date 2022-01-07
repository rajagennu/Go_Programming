package main

import "testing"

func Test_init_func(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := init_func(); got != tt.want {
				t.Errorf("init_func() = %v, want %v", got, tt.want)
			}
		})
	}
}
