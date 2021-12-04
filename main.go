package main

import (
	"fmt"
	"strconv"

	"github.com/WTFox/aoc2021/submarine"
	"github.com/WTFox/aoc2021/util"
)

func main() {
	Day03()
}

func Day01() {
	depths := util.ReadIntegersFromFile("./inputs/day01.txt")
	fmt.Println(countNumberOfPositiveChangesInDepth(depths))
}

func Day02() {
	instructions := util.ReadStringsFromFile("./inputs/day02.txt")
	s := submarine.New()
	s.ProcessInstructions(instructions)
	fmt.Println(s.Result())
}

func Day03() {
	outputs := util.ReadStringsFromFile("./inputs/day03.txt")
	fmt.Println(calculatePowerConsumption(outputs))
}

func tallyBits(outputs []string, length int) []int {
	tally := []int{}
	for i := 0; i < length; i++ {
		tally = append(tally, 0)
	}

	for _, output := range outputs {
		for i := 0; i < length; i++ {
			value := util.MustConvertToInt(string(output[i]))
			if value == 1 {
				tally[i]++
			}
		}
	}
	return tally
}

func calculateGammaAndEpsilon(talliedBits []int, threshold int) (int, int) {
	gamma := ""
	epsilon := ""
	for _, v := range talliedBits {
		if v < threshold {
			gamma = fmt.Sprintf("%s%d", gamma, 0)
			epsilon = fmt.Sprintf("%s%d", epsilon, 1)
		} else {
			gamma = fmt.Sprintf("%s%d", gamma, 1)
			epsilon = fmt.Sprintf("%s%d", epsilon, 0)
		}
	}
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	return int(gammaInt), int(epsilonInt)
}

func calculatePowerConsumption(outputs []string) int {
	length := len(outputs[0])
	tally := tallyBits(outputs, length)
	threshold := len(outputs) / 2

	gamma, epsilon := calculateGammaAndEpsilon(tally, threshold)

	return gamma * epsilon
}

func countNumberOfPositiveChangesInDepth(inputs []int) (result int) {
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
