package main

import (
	"testing"
)

func TestKnownHappies(t *testing.T) {
	testCases := []struct {
		x       int
		isHappy bool
	}{
		{44, true},
		{1, true},
		{45, false},
		{0, false},
	}
	for _, tc := range testCases {
		got, _ := isHappy(tc.x)
		if got != tc.isHappy {
			t.Errorf("For test-case (%v), got %v, want %v", tc, got, tc.isHappy)
		}
	}
}
