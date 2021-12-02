package utils

import (
	"os"
	"strconv"
	"strings"
)

func ReadIntegersFromFile(filename string) []int {
	body, err := os.ReadFile(filename)
	if err != nil {
		return []int{}
	}

	rawInputs := strings.Split(string(body), "\n")

	inputs := []int{}
	for i := 0; i < len(rawInputs); i++ {
		value := rawInputs[i]
		if value == "" {
			continue
		}

		intValue, err := strconv.Atoi(value)
		if err != nil {
			return []int{}
		}
		inputs = append(inputs, intValue)
	}

	return inputs
}
