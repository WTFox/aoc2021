package submarine

import "math"

func CalculateCrabMove(crabs []int, pos int) (totalFuelUsed int) {
	for _, crab := range crabs {
		totalFuelUsed += int(math.Abs(float64(crab - pos)))
	}
	return
}

func FindMostEfficientFuelUsage(crabs []int) (fuel int) {

	for i := 0; i < len(crabs); i++ {
		fuelUsed := CalculateCrabMove(crabs, i)
		if fuel == 0 {
			fuel = fuelUsed
		}
		if fuelUsed < fuel {
			fuel = fuelUsed
		}
	}
	return
}
