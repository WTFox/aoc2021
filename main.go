package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/WTFox/aoc2021/submarine"
	"github.com/WTFox/aoc2021/util"
)

func main() {
	Day05Part1()
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

type Coordinates struct{ x, y int }

func NewFromInputString(input string) (start, end Coordinates) {
	coordinates := strings.Split(input, " -> ")
	startCoordinateX, _ := strconv.Atoi(strings.Split(coordinates[0], ",")[0])
	startCoordinateY, _ := strconv.Atoi(strings.Split(coordinates[0], ",")[1])
	endCoordinateX, _ := strconv.Atoi(strings.Split(coordinates[1], ",")[0])
	endCoordinateY, _ := strconv.Atoi(strings.Split(coordinates[1], ",")[1])

	return Coordinates{startCoordinateX, startCoordinateY}, Coordinates{endCoordinateX, endCoordinateY}
}

func Day05Part1() {
	// inputs := util.ReadStringsFromFile("./inputs/day05.txt")
	// fmt.Println(sumCoordinateOverlaps(inputs))
	fmt.Println(rangeCoordinatesDiagonal(Coordinates{4, 0}, Coordinates{0, 4}))
}

func sumCoordinateOverlaps(inputs []string) int {
	counter := map[Coordinates]int{}

	for _, input := range inputs {
		start, end := NewFromInputString(input)

		var coordinates []Coordinates
		if start.x != end.x && start.y != end.y {
			coordinates = rangeCoordinatesDiagonal(start, end)
		} else {
			coordinates = rangeCoordinatesStraightLine(start, end)
		}

		for _, coordinate := range coordinates {
			counter[coordinate]++
		}
	}

	tally := 0
	for _, v := range counter {
		if v > 1 {
			tally++
		}
	}
	return tally
}

func rangeCoordinatesDiagonal(start, end Coordinates) (output []Coordinates) {
	steps := int(math.Abs(math.Max(float64(end.x-start.x), float64(end.y-start.y)) + 1))

	for i := 0; i < steps; i++ {
		output = append(output, Coordinates{start.x, start.y})
		if start.x < end.x {
			start.x++
		} else {
			start.x--
		}
		if start.y < end.y {
			start.y++
		} else {
			start.y--
		}
	}

	return
}

func rangeCoordinatesStraightLine(start, end Coordinates) (output []Coordinates) {
	startX := int(math.Min(float64(start.x), float64(end.x)))
	startY := int(math.Min(float64(start.y), float64(end.y)))

	endX := int(math.Max(float64(start.x), float64(end.x)))
	endY := int(math.Max(float64(start.y), float64(end.y)))

	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			output = append(output, Coordinates{x, y})
		}
	}
	return
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
