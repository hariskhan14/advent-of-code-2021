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
	prevSlideIncrease := -1
	slidingIncrease := 0
	for i, val := range numbers {
		if len(numbers[i:]) <= 2 {
			return slidingIncrease
		}

		slideMeasurement := val + numbers[i+1] + numbers[i+2]

		if prevSlideIncrease == -1 {
			prevSlideIncrease = slideMeasurement
		}

		if slideMeasurement > prevSlideIncrease {
			slidingIncrease++
		}
	}

	return slidingIncrease
}
