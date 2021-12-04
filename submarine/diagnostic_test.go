package submarine

import (
	"reflect"
	"testing"
)

func TestDiagnostic(t *testing.T) {
	rawInput := []string{
		"00100",
		"11110",
		"10110",
		"10111", // <--
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	t.Run("Test Gamma", func(t *testing.T) {
		expectedGamma := 22

		d := NewDiagnosticReport(rawInput)
		d.Process()

		got := d.Gamma
		if got != expectedGamma {
			t.Errorf("got %d want %d", got, expectedGamma)
		}
	})

	t.Run("Test Epsilon", func(t *testing.T) {
		expectedEpsilon := 9

		d := NewDiagnosticReport(rawInput)
		d.Process()

		got := d.Epsilon
		if got != expectedEpsilon {
			t.Errorf("got %d want %d", got, expectedEpsilon)
		}
	})

	t.Run("Test PowerConsumption", func(t *testing.T) {
		expectedPowerConsumption := 198

		d := NewDiagnosticReport(rawInput)
		d.Process()

		got := d.PowerConsumption()
		if got != expectedPowerConsumption {
			t.Errorf("got %d want %d", got, expectedPowerConsumption)
		}
	})

	t.Run("Test bitCounts", func(t *testing.T) {
		expected := []int{7, 5, 8, 7, 5}

		d := NewDiagnosticReport(rawInput)
		d.Process()

		got := d.bitCounts

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("got %d want %d", got, expected)
		}
	})

	t.Run("Test OxygenGeneratorRating", func(t *testing.T) {
		d := NewDiagnosticReport(rawInput)
		d.Process()

		got := d.OxygenGeneratorRating()
		if got != 23 {
			t.Errorf("got %d want %d", got, 23)
		}
	})

	t.Run("Test Co2ScrubberRating", func(t *testing.T) {
		d := NewDiagnosticReport(rawInput)
		d.Process()

		got := d.Co2ScrubberRating()
		if got != 10 {
			t.Errorf("got %d want %d", got, 10)
		}
	})

	t.Run("Test LifeSupportRating", func(t *testing.T) {
		d := NewDiagnosticReport(rawInput)
		d.Process()

		got := d.LifeSupportRating()
		if got != 230 {
			t.Errorf("got %d want %d", got, 230)
		}
	})
}

func TestFilterReportByMask(t *testing.T) {
	rawInput := []string{
		"00100",
		"11110",
		"10110",
		"10111", // <--
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	got := filterReportByMask(rawInput, "11100")
	expected := []string{
		"11100",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v want %v", got, expected)
	}
}
