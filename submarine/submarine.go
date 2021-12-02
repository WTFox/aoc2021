package submarine

import (
	"strconv"
	"strings"
)

type Submarine struct {
	positionX,
	// positionY, -- might be needed one day?
	positionZ,
	aim int
}

func FromInstructions(input []string) Submarine {
	s := Submarine{}
	s.processInputs(input)
	return s
}

func (s Submarine) Result() int {
	return s.positionX * s.positionZ
}

func (s Submarine) parseDirectionAndIncrementFromSingleInstruction(input string) (direction string, increment int) {
	input = strings.TrimSpace(input)
	results := strings.Split(input, " ")
	direction = results[0]
	increment, _ = strconv.Atoi(results[1])
	return
}

func (s *Submarine) processInputs(instructions []string) {
	for i := 0; i < len(instructions); i++ {
		direction, increment := s.parseDirectionAndIncrementFromSingleInstruction(instructions[i])

		switch direction {
		case "forward":
			s.positionX += increment
			s.positionZ += s.aim * increment
		case "up":
			s.aim -= increment
		case "down":
			s.aim += increment
		}
	}
}
