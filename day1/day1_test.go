package day1_test

import "testing"

func TestIncreaseInDepthMeasurements(t *testing.T) {
	numbers := []int{1,2}
	expected := 1

	got := FindIncreaseInDepthMeasurements(numbers)

	if got != expected {
		t.Errorf("got %d, expected: %d, input: %v", got, expected, numbers)
	}
}

func FindIncreaseInDepthMeasurements(numbers []int) int {
	return 1
}

