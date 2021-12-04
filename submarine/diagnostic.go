package submarine

import (
	"fmt"

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

func (d *Diagnostic) OxygenGeneratorRating() int {
	length := len(d.Report[0])
	majorityThreshold := len(d.Report) / 2

	report := d.Report
	mask := ""
	for i := 0; i < length; i++ {

		bitCounts := bitCounts(report, 1)
		bitValueWeCareAbout := 0
		if bitCounts[i] >= majorityThreshold {
			bitValueWeCareAbout = 1
		}

		mask = mask + util.IntToByteString(bitValueWeCareAbout)

		var paddedMask string
		if len(mask) == length-1 {
			paddedMask = mask + fmt.Sprint(bitValueWeCareAbout)
			report := filterReportByMask(report, paddedMask)
			return util.ByteStringToInt(report[0])
		} else {
			paddedMask = util.PadRight(mask, length, '0')
			report := filterReportByMask(report, paddedMask)
			if len(report) <= 2 {
				return util.ByteStringToInt(report[0])
			}
		}

	}
	return 0
}

func (d *Diagnostic) Co2ScrubberRating() int {
	mask := ""
	report := d.Report
	for i := 0; i < len(report[0]); i++ {
		if len(report) == 1 {
			return util.ByteStringToInt(report[0])
		}

		bitCounts := bitCounts(report, 1)
		numOfOnes := bitCounts[i]
		numOfZeros := len(report) - bitCounts[i]

		bitValueWeCareAbout := 0
		if numOfOnes < numOfZeros {
			bitValueWeCareAbout = 1
		}

		mask = mask + util.IntToByteString(bitValueWeCareAbout)
		report = filterReportByMask(report, mask)
	}

	return -1
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
	d.Gamma = util.ByteStringToInt(gamma)
	d.Epsilon = util.ByteStringToInt(epsilon)
}

func (d *Diagnostic) populateBitCounts() {
	d.bitCounts = bitCounts(d.Report, 1)
}

func filterReportByMask(report []string, mask string) (out []string) {
	for _, code := range report {
		if mask == code[:len(mask)] {
			out = append(out, code)
		}
	}
	return
}

func bitCounts(report []string, lookingFor int) []int {
	length := len(report[0])
	counts := []int{}

	for i := 0; i < length; i++ {
		counts = append(counts, 0)
	}

	for _, byteString := range report {
		for i := 0; i < length; i++ {
			value := util.MustConvertToInt(string(byteString[i]))
			if value == lookingFor {
				counts[i]++
			}
		}
	}
	return counts
}
