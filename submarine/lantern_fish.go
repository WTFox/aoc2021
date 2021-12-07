package submarine

import (
	"container/list"
)

func SimulateLanternFishLife(fishies *list.List, numDays uint16) int {
	if numDays == 0 {
		return fishies.Len()
	}
	for fishy := fishies.Front(); fishy != nil; fishy = fishy.Next() {
		if fishy.Value == 0 {
			fishy.Value = 6
			fishies.PushBack(9)
		} else {
			if v, ok := fishy.Value.(int); ok {
				fishy.Value = v - 1
			}
		}
	}
	return SimulateLanternFishLife(fishies, numDays-1)
}
