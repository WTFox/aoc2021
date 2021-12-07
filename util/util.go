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

		inputs = append(inputs, MustConvertToInt(value))
	}

	return inputs
}

func MustConvertToInt(input string) int {
	intValue, err := strconv.Atoi(input)
	if err != nil {
		panic("couldn't convert to int")
	}
	return intValue
}

func ByteStringToInt(byteString string) int {
	val, _ := strconv.ParseInt(byteString, 2, 64)
	return int(val)
}

func IntToByteString(val int) string {
	return strconv.FormatInt(int64(val), 2)
}

func StringToIntSlice(input string) (output []int) {
	for _, t := range strings.Split(input, ",") {
		val, _ := strconv.Atoi(t)
		output = append(output, val)
	}
	return
}
