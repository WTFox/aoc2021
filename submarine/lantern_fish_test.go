package submarine

import (
	"container/list"
	"testing"
)

func TestLanternFish(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	fishies := list.New()
	for _, num := range input {
		fishies.PushBack(num)
	}

	expected := 5934
	got := SimulateLanternFishLife(fishies, 80)
	if got != expected {
		t.Errorf("Got %d total fish, expected %d", got, expected)
	}
}

func BenchmarkLanternFish(b *testing.B) {
	input := []uint8{3, 4, 3, 1, 2}
	l := list.New()
	for _, num := range input {
		l.PushBack(num)
	}

	for n := 0; n < b.N; n++ {
		SimulateLanternFishLife(l, 80)
	}
}
