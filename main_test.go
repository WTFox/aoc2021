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

func TestCoordinates(t *testing.T) {
	inputs := []string{
		"0,9 -> 5,9",
		"0,9 -> 2,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
	expected := 12

	got := sumCoordinateOverlaps(inputs)

	if got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}

}
