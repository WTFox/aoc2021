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
	aim int
}

// New returns a new instance of Submarine
func New() Submarine {
	return Submarine{}
}

// Result returns the product of positionX and positionZ
func (s *Submarine) Result() int {
	return s.positionX * s.positionZ
}

// ProcessInstructions runs through the instructionset provided and updates Submarine's location/aim.
func (s *Submarine) ProcessInstructions(instructions []string) {
	for _, instruction := range instructions {
		direction, increment := parseInstruction(instruction)

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

func parseInstruction(instruction string) (direction string, increment int) {
	instruction = strings.TrimSpace(instruction)
	results := strings.Split(instruction, " ")
	direction = results[0]
	increment, _ = strconv.Atoi(results[1])
	return direction, increment
}
