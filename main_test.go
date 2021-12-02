package main

import "testing"

func TestCountIncreases(t *testing.T) {
	testCases := []struct {
		depths []int
		wanted int
	}{
		{[]int{1, 1, 1}, 0},
		{[]int{1, 1, 1, 2}, 1},
		{[]int{1, 1, 1, 2, 2}, 2},
		{[]int{1, 1, 1, 2, 2, 2}, 3},
		{[]int{1, 1, 1, 0, 0, 0}, 0},
	}

	for _, tt := range testCases {
		got := countNumberOfPositiveChangesInDepth(tt.depths)
		if got != tt.wanted {
			t.Errorf("got %d want %d", got, tt.wanted)
		}
	}
}
