package submarine

import "testing"

func TestSubmarineResult(t *testing.T) {
	testCases := []struct {
		instructions []string
		expectedResult     int
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
		got := New(tt.instructions).Result
		if got != tt.expectedResult {
			t.Errorf("got %d expected %d", got, tt.expectedResult)
		}
	}
}
