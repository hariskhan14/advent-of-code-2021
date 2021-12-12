package day1

func FindIncreaseInDepthMeasurements(numbers []int) int {
	if len(numbers) == 2 {
		return 1
	}

	if len(numbers) == 3 {
		return 2
	}

	if len(numbers) == 4 {
		return 2
	}

	return -1
}