package coordinate

import (
	"strconv"
	"strings"
)

type Coordinate struct {
	X int // forwards & backwards
	Y int // left & right
	Z int // up & down
}

func NewFromInputs(input []string) Coordinate {
	c := Coordinate{}
	c.calculateDestination(input)
	return c
}

func (c Coordinate) Result() int {
	return c.X * c.Z
}

func (c Coordinate) parseDirectionAndIncrement(input string) (direction string, increment int) {
	input = strings.TrimSpace(input)
	results := strings.Split(input, " ")
	direction = results[0]
	increment, _ = strconv.Atoi(results[1])
	return
}

func (c *Coordinate) calculateDestination(coordInputs []string) {
	for i := 0; i < len(coordInputs); i++ {
		direction, increment := c.parseDirectionAndIncrement(coordInputs[i])
		switch direction {
		case "forward":
			c.X += increment
		case "up":
			c.Z -= increment
		case "down":
			c.Z += increment
		}
	}
}
