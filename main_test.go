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

func TestCalculatePowerConsumption(t *testing.T) {
	testCases := []struct {
		outputs          []string
		powerConsumption int
	}{
		{[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}, 198},
	}

	t.Run("CalculatePowerConsumption", func(t *testing.T) {
		for _, tt := range testCases {
			got := calculatePowerConsumption(tt.outputs)
			if got != tt.powerConsumption {
				t.Errorf("got %d want %d", got, tt.powerConsumption)
			}
		}
	})
}
