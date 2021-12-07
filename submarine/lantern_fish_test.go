package submarine

import "testing"

func TestLanternFish(t *testing.T) {
	input := []uint8{3, 4, 3, 1, 2}

	expected := 5934
	got := len(SimulateLanternFishLife(input, 80))
	if got != expected {
		t.Errorf("Got %d total fish, expected %d", got, expected)
	}
}

func BenchmarkLanternFish(b *testing.B) {
	input := []uint8{3, 4, 3, 1, 2}

	for n := 0; n < b.N; n++ {
		SimulateLanternFishLife(input, 80)
	}
}
