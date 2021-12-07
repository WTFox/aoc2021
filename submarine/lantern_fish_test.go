package submarine

import (
	"testing"
)

func TestLanternFish(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	expected := 5934
	got := SimulateLanternFishLife(input, 80)
	if got != expected {
		t.Errorf("Got %d total fish, expected %d", got, expected)
	}
}

func BenchmarkLanternFish(b *testing.B) {
	input := []int{3, 4, 3, 1, 2}

	for n := 0; n < b.N; n++ {
		SimulateLanternFishLife(input, 80)
	}
}
