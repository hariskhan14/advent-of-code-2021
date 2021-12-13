package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFileAsInts(filename string) []int {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []int{}
	}

	lines := strings.Split(string(b), "\n")
	numbers := make([]int, len(lines))

	for i, l := range lines {
		if len(l) == 0 {
			continue
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			return []int{}
		}
		numbers[i] = n
	}

	return numbers
}

func ReadFileAsStrings(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}
	}

	lines := strings.Split(string(b), "\n")

	return lines
}