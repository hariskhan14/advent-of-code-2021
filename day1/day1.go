package day1

func FindIncreaseInDepthMeasurements(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	previous := numbers[0]
	increase := 0

	for _, val := range numbers {
		if val > previous {
			increase++
		}

		previous = val
	}

	return increase
}

func FindSlidingIncreaseInDepthMeasurements(numbers []int) int {
	if len(numbers) == 3 {
		return 0
	}

	if len(numbers) == 4 {
		return 1
	}

	return -1
}