package main

import "testing"

func TestCountIncreases(t *testing.T) {
	testCases := []struct {
		inputs []int
		wanted int
	}{
		{[]int{1, 1, 1}, 0},
		{[]int{1, 1, 1, 2}, 1},
		{[]int{1, 1, 1, 2, 2}, 2},
		{[]int{1, 1, 1, 2, 2, 2}, 3},
		{[]int{1, 1, 1, 0, 0, 0}, 0},
	}

	for _, tt := range testCases {
		got := countIncreases(tt.inputs)
		if got != tt.wanted {
			t.Errorf("got %d want %d", got, tt.wanted)
		}
	}
}
