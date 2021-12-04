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
		"10111",
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
}
