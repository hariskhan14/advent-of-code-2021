package day2

import (
	"strconv"
	"strings"
)

func LetsDive(commands []string) int {
	result := 0

	for _, val := range commands {
		split := strings.Split(val, " ")
		commandValue, err := strconv.Atoi(split[1])
		if err != nil {
			return -1
		}

		if result == 0 {
			result = commandValue
			continue
		}

		result *= commandValue
	}
	return result
}
