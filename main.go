package main

import (
	"fmt"

	"github.com/WTFox/aoc2021/utils"
)

func main() {
	inputs := utils.ReadIntegersFromFile("./inputs/day01/inputs.txt")
	fmt.Println(countIncreases(inputs))
}

func countIncreases(inputs []int) (result int) {
	for index := 1; index < len(inputs); index++ {
		if index+2 > len(inputs)-1 {
			break
		}

		lastWindowSum := inputs[index-1] + inputs[index] + inputs[index+1]
		currentWindowSum := inputs[index] + inputs[index+1] + inputs[index+2]
		if lastWindowSum < currentWindowSum {
			result++
		}
	}
	return result
}
