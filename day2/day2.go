package day2

import (
	"strconv"
	"strings"
)

func LetsDive(commands []string) int {
	result := 1

	for _, val := range commands {
		split := strings.Split(val, " ")
		commandValue, err := strconv.Atoi(split[1])
		if err != nil {
			return -1
		}

		result *= commandValue
	}
	return result
}
