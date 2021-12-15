package day3_test

import (
	"github.com/hariskhan14/advent-of-code-2021/day3"
	"github.com/hariskhan14/advent-of-code-2021/utils"
	"testing"
)

func Test_DecodeBinaryCodes(t *testing.T) {
	testSuites := []struct {
		name     string
		report   []string
		expected int
	}{
		{name: "base case", report: []string{"00100", "11110", "10111"}, expected: 198},                    //gamma-rate: 10110(22), epsilon-rate: 01001(9)
		{name: "adding another code", report: []string{"00100", "11110", "01110", "01011"}, expected: 238}, //gamma-rate: 01110(14), epsilon-rate: 10000(17)
		{name: "sample input", report: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}, expected: 198},
		{name: "puzzle input", report: utils.ReadFileAsStrings("puzzle_input.txt"), expected: 3813416},
	}

	for _, testSuite := range testSuites {
		t.Run(testSuite.name, func(t *testing.T) {
			report := testSuite.report
			expected := testSuite.expected
			got := day3.CalculatePowerConsumption(report)

			if got != expected {
				t.Errorf("got: %d, expected: %d", got, expected)
			}
		})
	}
}

func Test_DecodeBinaryCodes_PartTwo(t *testing.T) {
	t.Run("calculate oxygen rating", func(t *testing.T) {
		report := []string{"00100", "11110", "10111"} // 1. (1) in majority => {11110, 10111} 2. equal so take 1 => oxygen = {11110} = 30
		expected := 30

		got := day3.CalculateOxygenGeneratorRating(report)
		if got != expected {
			t.Errorf("got %d, expected: %d", got, expected)
		}
	})
	t.Run("calculate oxygen rating with more details", func(t *testing.T) {
		report := []string{"00100", "11110", "10110", "10111", "10101"}
			// 1. (1) in majority => {11110, 10110, 10111, 10101}
			// 2. (0) in majority => {10110, 10111, 10101}
			// 3. (1) in majority => {10110, 10111, 10101}
			// 4. (1) in majority => {10110, 10111}
			// 5. (1) in majority => {10111}  => 22
		expected := 23

		got := day3.CalculateOxygenGeneratorRating(report)
		if got != expected {
			t.Errorf("got %d, expected: %d", got, expected)
		}
	})

}
