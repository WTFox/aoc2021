package submarine

import (
	"strconv"
	"strings"
)

// Submarine is the entrypoint struct for this package
type Submarine struct {
	positionX,
	// positionY, -- might be needed one day?
	positionZ,
	aim,
	Result int
}

// New is the constructor for Submarine
func New(input []string) Submarine {
	s := Submarine{}
	s.processInstructions(input)
	s.Result = s.positionX * s.positionZ
	return s
}

func (s Submarine) parseDirectionAndIncrementFromSingleInstruction(input string) (direction string, increment int) {
	input = strings.TrimSpace(input)
	results := strings.Split(input, " ")
	direction = results[0]
	increment, _ = strconv.Atoi(results[1])
	return
}

func (s *Submarine) processInstructions(instructions []string) {
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
