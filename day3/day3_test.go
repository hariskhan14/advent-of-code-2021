package day3_test

import (
	"github.com/hariskhan14/advent-of-code-2021/day3"
	"testing"
)

func Test_DecodeBinaryCodes(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		report := []string{"00100", "11110", "10111"} //gamma-rate: 10110(22), epsilon-rate: 01001(9)
		expected := 198
		got := day3.CalculatePowerConsumption(report)

		if got != expected {
			t.Errorf("got: %d, expected: %d", got, expected)
		}
	})

	t.Run("adding another code", func(t *testing.T) {
		report := []string{"00100", "11110", "01110", "01011"} //gamma-rate: 01110(14), epsilon-rate: 10000(17)
		expected := 238
		got := day3.CalculatePowerConsumption(report)

		if got != expected {
			t.Errorf("got: %d, expected: %d", got, expected)
		}
	})
}
