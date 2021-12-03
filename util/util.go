package util

import (
	"os"
	"strconv"
	"strings"
)

// ReadStringsFromFile reads a file and produces an array of strings representing each line.
func ReadStringsFromFile(filename string) []string {
	body, err := os.ReadFile(filename)
	if err != nil {
		return []string{}
	}

	rawInputs := strings.Split(string(body), "\n")

	inputs := []string{}
	for i := 0; i < len(rawInputs); i++ {
		value := rawInputs[i]
		if value == "" {
			continue
		}

		if err != nil {
			return []string{}
		}
		inputs = append(inputs, value)
	}

	return inputs
}

// ReadIntegersFromFile reads a file and produces an array of integers.
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
