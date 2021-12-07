package submarine

import (
	"math"
)

func CalculateCrabMove(crabs []int, pos int) int {
	var totalFuelUsed int
	for _, crab := range crabs {
		steps := int(math.Abs(float64(crab - pos)))
		totalFuelUsed += steps + steps*(steps-1)/2
	}
	return totalFuelUsed
}

func FindMostEfficientFuelUsage(crabsInSubs []int) int {
	lowestFuelUsage := CalculateCrabMove(crabsInSubs, 0)
	for i := 0; i < len(crabsInSubs); i++ {
		lowestFuelUsage = int(math.Min(float64(lowestFuelUsage), float64(CalculateCrabMove(crabsInSubs, i))))
	}
	return lowestFuelUsage
}
