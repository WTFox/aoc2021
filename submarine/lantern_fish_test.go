package submarine

import "testing"

func TestLanternFish(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	fishies := BuildFishies(input)
	expected := 5934
	got := len(SimulateLanternFishLife(fishies, 80))
	if got != expected {
		t.Errorf("Got %d total fish, expected %d", got, expected)
	}
}

func BenchmarkLanternFish(b *testing.B) {
	input := []int{3, 4, 3, 1, 2}
	fishies := BuildFishies(input)

	for n := 0; n < b.N; n++ {
		SimulateLanternFishLife(fishies, 80)
	}
}
