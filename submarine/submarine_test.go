package submarine

import "testing"

func TestSubmarineCoordinates(t *testing.T) {
	testCases := []struct {
		instructions []string
		expected     int
	}{
		{[]string{}, 0},
		{[]string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}, 900},
	}

	for _, tt := range testCases {
		got := FromInstructions(tt.instructions).Result()
		if got != tt.expected {
			t.Errorf("got %d expected %d", got, tt.expected)
		}
	}
}
