package submarine

import (
	"fmt"
	"sort"
	"strings"
)

var lengthToNum = map[int]int{
	2: 1,
	3: 7,
	4: 4,
	8: 7,
}

func DisplaySearch(inputs []string) int {
	num := 0
	for _, v := range inputs {
		values := strings.Split(v, "|")
		inputValue := values[0]
		outputValue := values[1]
		inputValues := strings.Split(inputValue, " ")
		outputValues := strings.Split(outputValue, " ")

		for _, v := range inputValues {
			fmt.Println(SortStringByCharacter(v))
		}

		for _, v := range outputValues {
			if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
				num++
			}
		}
	}

	return num
}

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}
