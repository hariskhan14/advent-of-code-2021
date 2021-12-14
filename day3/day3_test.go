package day3_test

import (
	"github.com/hariskhan14/advent-of-code-2021/day3"
	"testing"
)

func Test_DecodeBinaryCodes(t *testing.T) {
	testSuites := []struct{
		name string
		report []string
		expected int
	} {
		{name: "base case", report: []string{"00100", "11110", "10111"}, expected: 198}, //gamma-rate: 10110(22), epsilon-rate: 01001(9)
		{name: "adding another code", report: []string{"00100", "11110", "01110", "01011"}, expected: 238}, //gamma-rate: 01110(14), epsilon-rate: 10000(17)
		{name: "sample input", report: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}, expected: 198},
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
