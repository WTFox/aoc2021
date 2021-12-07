package submarine

import (
	"fmt"
	"math"
)

func CalculateCrabMove(crabs []int, pos int) int {
	var totalFuelUsed int
	for _, crab := range crabs {
		steps := int(math.Abs(float64(crab - pos)))
		for i := 0; i < steps; i++ {
			totalFuelUsed += i
		}
		totalFuelUsed += steps
	}
	return totalFuelUsed
}

func FindMostEfficientFuelUsage(crabs []int) int {
	lowestFuelUsage := 0

	for i := 0; i < len(crabs); i++ {
		fuelUsed := CalculateCrabMove(crabs, i)
		if lowestFuelUsage == 0 {
			lowestFuelUsage = fuelUsed
		}
		if fuelUsed < lowestFuelUsage {
			lowestFuelUsage = fuelUsed
		}
		fmt.Println(i, fuelUsed)
	}

	return lowestFuelUsage
}
