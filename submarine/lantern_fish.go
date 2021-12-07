package submarine

func SimulateLanternFishLife(fishies []int, numDays int) int {
	fishStates := []int{}
	for i := 0; i < 9; i++ {
		fishStates = append(fishStates, 0)
	}

	for _, input := range fishies {
		fishStates[input] += 1
	}

	for i := 0; i < numDays; i++ {
		// pop 0th element, then append it
		fishStates = append(fishStates[1:], fishStates[:1]...)
		// add that count to 6
		fishStates[6] += fishStates[8]
	}

	sum := 0
	for _, v := range fishStates {
		sum += v
	}

	return sum
}
