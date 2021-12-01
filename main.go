package aoc2021

func countIncreases(inputs []int) int {
	result := 0
	for i := 0; i < len(inputs); i++ {
		if i == 0 {
			continue
		}

		if inputs[i-1] < inputs[i] {
			result++
		}
	}

	return result
}
