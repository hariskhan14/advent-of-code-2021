package day2

import (
	"strconv"
	"strings"
)

func LetsDive(commands []string) int {
	forward := 0
	depth := 0

	for _, val := range commands {
		split := strings.Split(val, " ")
		command := split[0]
		commandValue, err := strconv.Atoi(split[1])
		if err != nil {
			return -1
		}

		switch command {
		case "forward":
			forward += commandValue
		case "down":
			depth += commandValue
		case "up":
			depth -= commandValue
		default:
			return -1
		}
	}

	return forward * depth
}

func LetsDiveWithAim(commands []string) int {
	forward := 0
	depth := 0
	aim := 0

	for _, val := range commands {
		split := strings.Split(val, " ")
		command := split[0]
		commandValue, err := strconv.Atoi(split[1])
		if err != nil {
			return -1
		}

		switch command {
		case "forward":
			forward += commandValue
			depth += commandValue * aim
		case "down":
			aim += commandValue
		case "up":
			aim -= commandValue
		default:
			return -1
		}
	}

	return forward * depth
}