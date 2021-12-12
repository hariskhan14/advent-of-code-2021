package day1_test

import (
	"github.com/hariskhan14/advent-of-code-2021/day1"
	"github.com/hariskhan14/advent-of-code-2021/utils"
	"testing"
)

func TestIncreaseInDepthMeasurements(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "increase with 2 measurements", input: []int{1, 2}, expected: 1},
		{name: "increase with 3 measurements", input: []int{1, 2, 3}, expected: 2},
		{name: "increase with 4 measurements", input: []int{1, 2, 3, 2}, expected: 2},
		{name: "increase with 7 measurements (sample input)", input: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, expected: 7},
		{name: "increase with 2k measurements (puzzle input)", input: utils.ReadFile("sample_input.txt"), expected: 1791},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			numbers := testCase.input
			expected := testCase.expected

			got := day1.FindIncreaseInDepthMeasurements(numbers)

			if got != expected {
				t.Errorf("got %d, expected: %d, input: %v", got, expected, numbers)
			}
		})
	}
}

func TestFindSlidingIncreaseInDepthMeasurements(t *testing.T) {
	t.Run("sliding increase with 1 window", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		expected := 0

		got := day1.FindSlidingIncreaseInDepthMeasurements(numbers)

		if got != expected {
			t.Errorf("got %d, expected: %d, input: %v", got, expected, numbers)
		}
	})

	t.Run("sliding increase with 2 windows", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4} // 1+2+3=6, 2+3+4=9
		expected := 1

		got := day1.FindSlidingIncreaseInDepthMeasurements(numbers)

		if got != expected {
			t.Errorf("got %d, expected: %d, input: %v", got, expected, numbers)
		}
	})

	t.Run("sliding increase with 3 windows", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5} // 1+2+3=6, 2+3+4=9, 3+4+5=12
		expected := 2

		got := day1.FindSlidingIncreaseInDepthMeasurements(numbers)

		if got != expected {
			t.Errorf("got %d, expected: %d, input: %v", got, expected, numbers)
		}
	})

}
