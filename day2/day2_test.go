package day2_test

import (
	"github.com/hariskhan14/advent-of-code-2021/day2"
	"testing"
)

func Test_Dive(t *testing.T) {
	testCases := []struct {
		name     string
		commands []string
		expected int
	}{
		{name: "easy forward and down command", commands: []string{"forward 1", "down 1"}, expected: 1},
		{name: "increase depth in same command", commands: []string{"forward 1", "down 2"}, expected: 2},
		{name: "increasing both forward and depth", commands: []string{"forward 2", "down 3"}, expected: 6},
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
