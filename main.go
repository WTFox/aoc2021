package main

import (
	"fmt"

	"github.com/WTFox/aoc2021/coordinate"
	"github.com/WTFox/aoc2021/util"
)

func main() {
	// day 1
	countInputs := util.ReadIntegersFromFile("./inputs/day01.txt")
	fmt.Println(countIncreases(countInputs))

	// day 2
	coordInputs := util.ReadStringsFromFile("./inputs/day02.txt")
	fmt.Println(coordinate.NewFromInputs(coordInputs).Result())
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
