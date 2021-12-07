package submarine

import "testing"

func Test_CalculateCrabMove(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	testCases := []struct {
		pos,
		cost int
	}{
		{2, 206},
	}

	for _, tt := range testCases {
		got := CalculateCrabMove(input, tt.pos)
		if got != tt.cost {
			t.Errorf("Expected %d got %d", tt.cost, got)
		}
	}
}

func Test_FindMostEfficientFuelUsage(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	got := FindMostEfficientFuelUsage(input)
	expected := 168
	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}
}
