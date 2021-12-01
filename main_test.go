package aoc2021

import "testing"

func TestCountIncreases(t *testing.T) {
	testCases := []struct {
		inputs   []int
		expected int
	}{
		{[]int{12, 2}, 0},
		{[]int{12, 13}, 1},
		{[]int{12, 13, 14}, 2},
		{[]int{12, 13, 12}, 1},
	}

	for _, tt := range testCases {
		got := countIncreases(tt.inputs)
		if got != tt.expected {
			t.Errorf("got %d want %d", got, tt.expected)
		}
	}
}
