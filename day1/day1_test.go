package day1_test

import (
	"github.com/hariskhan14/advent-of-code-2021/day1"
	"testing"
)

func TestIncreaseInDepthMeasurements(t *testing.T) {
	t.Run("increase with 2 measurements", func(t *testing.T) {
		numbers := []int{1, 2}
		expected := 1

		got := day1.FindIncreaseInDepthMeasurements(numbers)

		if got != expected {
			t.Errorf("got %d, expected: %d, input: %v", got, expected, numbers)
		}
	})

	t.Run("increase with 3 measurements", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		expected := 2

		got := day1.FindIncreaseInDepthMeasurements(numbers)

		if got != expected {
			t.Errorf("got %d, expected: %d, input: %v", got, expected, numbers)
		}
	})

	t.Run("increase with 4 measurements", func(t *testing.T) {
		numbers := []int{1, 2, 3, 2}
		expected := 2

		got := day1.FindIncreaseInDepthMeasurements(numbers)

		if got != expected {
			t.Errorf("got %d, expected: %d, input: %v", got, expected, numbers)
		}
	})
}



