package submarine

import (
	"strings"
)

func DisplaySearch(inputs []string) int {
	num := 0
	for _, v := range inputs {
		values := strings.Split(v, "|")
		// inputValue := values[0]
		outputValue := values[1]
		outputValues := strings.Split(outputValue, " ")

		for _, v := range outputValues {
			if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
				num++
			}
		}
	}

	return num
}
