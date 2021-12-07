package submarine

type LanternFish struct {
	daysLeftBeforeSpawn int
	canSpawn            bool
}

func (l *LanternFish) Tick() bool {
	l.daysLeftBeforeSpawn--
	shouldSpawnNew := false
	if l.daysLeftBeforeSpawn < 0 {
		l.daysLeftBeforeSpawn = 6
		shouldSpawnNew = true
	}
	l.canSpawn = true
	return shouldSpawnNew
}

func BuildFishies(timers []int) (fishies []LanternFish) {
	for _, timer := range timers {
		fishies = append(fishies, LanternFish{daysLeftBeforeSpawn: timer, canSpawn: false})
	}
	return
}

func SimulateLanternFishLife(fishies []LanternFish, numDays int) []LanternFish {
	if numDays == 0 {
		return fishies
	}
	for i := 0; i < len(fishies); i++ {
		fish := &fishies[i]
		if fish.Tick() {
			fishies = append(fishies, LanternFish{daysLeftBeforeSpawn: 9, canSpawn: false})
		}
	}
	return SimulateLanternFishLife(fishies, numDays-1)
}
