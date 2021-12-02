package main

import "testing"

func TestCountNumberOfPositiveChangesInDepth(t *testing.T) {
	testCases := []struct {
		depths          []int
		expectedChanges int
	}{
		{[]int{1, 1, 1}, 0},
		{[]int{1, 1, 1, 2}, 1},
		{[]int{1, 1, 1, 2, 2}, 2},
		{[]int{1, 1, 1, 2, 2, 2}, 3},
		{[]int{1, 1, 1, 0, 0, 0}, 0},
	}

	for _, tt := range testCases {
		got := countNumberOfPositiveChangesInDepth(tt.depths)
		if got != tt.expectedChanges {
			t.Errorf("got %d want %d", got, tt.expectedChanges)
		}
	}
}
