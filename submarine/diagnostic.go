package submarine

import (
	"fmt"
	"strconv"

	"github.com/WTFox/aoc2021/util"
)

type Diagnostic struct {
	Gamma     int
	Epsilon   int
	Report    []string
	bitCounts []int
}

func NewDiagnosticReport(report []string) Diagnostic {
	return Diagnostic{Report: report}
}

func (d *Diagnostic) Process() {
	d.populateBitCounts()
	d.populateGammaAndEpsilon()
}

func (d *Diagnostic) PowerConsumption() int {
	return d.Gamma * d.Epsilon
}

func (d *Diagnostic) populateGammaAndEpsilon() {
	gamma := ""
	epsilon := ""

	for _, v := range d.bitCounts {
		if v < len(d.Report)/2 {
			gamma = fmt.Sprintf("%s%d", gamma, 0)
			epsilon = fmt.Sprintf("%s%d", epsilon, 1)
		} else {
			gamma = fmt.Sprintf("%s%d", gamma, 1)
			epsilon = fmt.Sprintf("%s%d", epsilon, 0)
		}
	}
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)

	d.Gamma = int(gammaInt)
	d.Epsilon = int(epsilonInt)
}

func (d *Diagnostic) populateBitCounts() {
	length := len(d.Report[0])
	counts := []int{}

	for i := 0; i < length; i++ {
		counts = append(counts, 0)
	}

	for _, byteString := range d.Report {
		for i := 0; i < length; i++ {
			value := util.MustConvertToInt(string(byteString[i]))
			if value == 1 {
				counts[i]++
			}
		}
	}

	d.bitCounts = counts
}
