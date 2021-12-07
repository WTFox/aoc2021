package submarine

import (
	"math"
	"strconv"
	"strings"
)

type Coordinates struct{ x, y int }

func NewFromInputString(input string) (start, end Coordinates) {
	coordinates := strings.Split(input, " -> ")
	startCoordinateX, _ := strconv.Atoi(strings.Split(coordinates[0], ",")[0])
	startCoordinateY, _ := strconv.Atoi(strings.Split(coordinates[0], ",")[1])
	endCoordinateX, _ := strconv.Atoi(strings.Split(coordinates[1], ",")[0])
	endCoordinateY, _ := strconv.Atoi(strings.Split(coordinates[1], ",")[1])

	return Coordinates{startCoordinateX, startCoordinateY}, Coordinates{endCoordinateX, endCoordinateY}
}

func SumVentOverlaps(inputs []string) int {
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
	startX := start.x
	startY := start.y

	steps := int(math.Abs(math.Max(float64(end.x-start.x), float64(end.y-start.y))))

	output = append(output, Coordinates{startX, startY})
	for i := 0; i < steps; i++ {

		if startX < end.x {
			startX++
		} else {
			startX--
		}
		if startY < end.y {
			startY++
		} else {
			startY--
		}
		output = append(output, Coordinates{startX, startY})
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
