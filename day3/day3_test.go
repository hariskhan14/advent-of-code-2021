package day3_test

import (
	"github.com/hariskhan14/advent-of-code-2021/day3"
	"testing"
)

func Test_DecodeBinaryCodes(t *testing.T) {
	report := []string{"00100", "11110", "10110"} //gamma-rate: 10110(22), epsilon-rate: 01100(12)
	expected := 264
	got := day3.CalculatePowerConsumption(report)

	if got != expected {
		t.Errorf("got: %d, expected: %d", got, expected)
	}
}
