package main

import (
	"fmt"

	"github.com/WTFox/aoc2021/submarine"
	"github.com/WTFox/aoc2021/util"
)

func main() {
	Day07Part1()
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
	bytes := util.ReadStringsFromFile("./inputs/day03.txt")
	d := submarine.NewDiagnosticReport(bytes)
	d.Process()
	fmt.Println(d.LifeSupportRating())
}

func Day04Part1() {
	inputs := util.ReadStringsFromFile("./inputs/day04.txt")
	bingoBoards := submarine.NewBingoMultiple(inputs)
	fmt.Printf("%d boards\n", len(bingoBoards))

	quickestWinner := &bingoBoards[0]
	for index := 0; index < len(bingoBoards); index++ {
		bingo := &bingoBoards[index]
		bingo.Play()
		if bingo.HasBingo() {
			score := bingo.Score()
			if bingo.NumTurns < quickestWinner.NumTurns {
				fmt.Printf("Board %d won in %d moves with score %d.\n", index, bingo.NumTurns, score)
				quickestWinner = bingo
				continue
			}
		}
	}

	fmt.Printf("Quickest board won in %d moves with score %d.\n", quickestWinner.NumTurns, quickestWinner.Score())
}

func Day04Part2() {
	inputs := util.ReadStringsFromFile("./inputs/day04.txt")
	bingoBoards := submarine.NewBingoMultiple(inputs)
	fmt.Printf("%d boards\n", len(bingoBoards))

	lastWinner := &bingoBoards[0]
	for index := 0; index < len(bingoBoards); index++ {
		bingo := &bingoBoards[index]
		bingo.Play()
		if bingo.HasBingo() {
			score := bingo.Score()
			if bingo.NumTurns > lastWinner.NumTurns {
				fmt.Printf("Board %d won in %d moves with score %d.\n", index, bingo.NumTurns, score)
				lastWinner = bingo
				continue
			}
		}
	}

	fmt.Printf("Last board won in %d moves with score %d.\n", lastWinner.NumTurns, lastWinner.Score())
}

func Day05() {
	inputs := util.ReadStringsFromFile("./inputs/day05.txt")
	fmt.Println(submarine.SumVentOverlaps(inputs))
}

func Day06() {
	fishesAsString := util.ReadStringsFromFile("./inputs/day06-lanternfish.txt")[0]
	fishies := util.StringToIntSlice(fishesAsString)

	fmt.Println(submarine.SimulateLanternFishLife(fishies, 256))
}

func Day07Part1() {
	inputString := util.ReadStringsFromFile("./inputs/day07-crabs-in-submarines.txt")[0]
	crabSubmarines := util.StringToIntSlice(inputString)

	fmt.Println(submarine.FindMostEfficientFuelUsage(crabSubmarines))

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
