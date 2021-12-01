package main

import "fmt"

func main() {
	fmt.Println(countIncreases(ReadIntegersFromFile("./inputs.txt")))
}

func countIncreases(inputs []int) int {
	result := 0

	for index := 0; index < len(inputs); index++ {
		if index == 0 {
			continue
		}

		if index+2 > len(inputs)-1 {
			break
		}

		previousWindowSum := inputs[index-1] + inputs[index] + inputs[index+1]
		currentWindowSum := inputs[index] + inputs[index+1] + inputs[index+2]
		if previousWindowSum < currentWindowSum {
			result++
		}
	}

	return result
}
