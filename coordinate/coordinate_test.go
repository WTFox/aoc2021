package coordinate

import "testing"

func TestCoordinate(t *testing.T) {
	testCases := []struct {
		inputs []string
		wanted int
	}{
		{[]string{}, 0},
		{[]string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}, 150},
	}

	for _, tt := range testCases {
		got := NewFromInputs(tt.inputs).Result()
		if got != tt.wanted {
			t.Errorf("got %d want %d", got, tt.wanted)
		}
	}
}
