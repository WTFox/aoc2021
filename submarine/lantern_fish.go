package submarine

func SimulateLanternFishLife(fishies []uint8, numDays uint16) []uint8 {
	if numDays == 0 {
		return fishies
	}
	for i := 0; i < len(fishies); i++ {
		if fishies[i] == 0 {
			fishies[i] = 6
			fishies = append(fishies, 9)
		} else {
			fishies[i]--
		}
	}
	return SimulateLanternFishLife(fishies, numDays-1)
}
