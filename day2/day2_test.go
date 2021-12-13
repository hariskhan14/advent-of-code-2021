package day2_test

import (
	"github.com/hariskhan14/advent-of-code-2021/day2"
	"github.com/hariskhan14/advent-of-code-2021/utils"
	"testing"
)

func Test_LetsDive(t *testing.T) {
	testCases := []struct {
		name     string
		commands []string
		expected int
	}{
		{name: "easy forward and down command", commands: []string{"forward 1", "down 1"}, expected: 1},
		{name: "increase depth in same command", commands: []string{"forward 1", "down 2"}, expected: 2},
		{name: "increasing both forward and depth", commands: []string{"forward 2", "down 3"}, expected: 6},
		{name: "introducing up command", commands: []string{"forward 2", "down 3", "up 1"}, expected: 4},
		{name: "sample input", commands: []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}, expected: 150},
		{name: "puzzle input", commands: utils.ReadFileAsStrings("sample_input.txt"), expected: 1924923},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			commands := testCase.commands
			expected := testCase.expected

			got := day2.LetsDive(commands)

			if got != expected {
				t.Errorf("got %d, expected: %d, input: %v", got, expected, commands)
			}
		})
	}
}

func Test_LetsDiveWithAim(t *testing.T) {
	testCases := []struct {
		name     string
		commands []string
		expected int
	}{
		{name: "base case", commands: []string{"forward 1", "down 1", "forward 3"}, expected: 12},
		{name: "added aim case", commands: []string{"forward 2", "down 2", "forward 3"}, expected: 30},
		{name: "sample case", commands: []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}, expected: 900},
		{name: "puzzle input", commands: utils.ReadFileAsStrings("sample_input.txt"), expected: 1982495697},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			commands := testCase.commands
			expected := testCase.expected

			got := day2.LetsDiveWithAim(commands)

			if got != expected {
				t.Errorf("got %d, expected: %d, input: %v", got, expected, commands)
			}
		})
	}

}