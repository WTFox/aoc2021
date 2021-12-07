package submarine

type LanternFish struct {
	daysLeftBeforeSpawn int
}

func SimulateLanternFishLife(fishies []LanternFish, numDays int) []LanternFish {
	if numDays == 0 {
		return fishies
	}
	for i := 0; i < len(fishies); i++ {
		fish := &fishies[i]
		fish.daysLeftBeforeSpawn--
		if fish.daysLeftBeforeSpawn < 0 {
			fish.daysLeftBeforeSpawn = 6
			fishies = append(fishies, LanternFish{daysLeftBeforeSpawn: 9})
		}
	}
	return SimulateLanternFishLife(fishies, numDays-1)
}

func BuildFishies(timers []int) (fishies []LanternFish) {
	for _, timer := range timers {
		fishies = append(fishies, LanternFish{daysLeftBeforeSpawn: timer})
	}
	return
}
