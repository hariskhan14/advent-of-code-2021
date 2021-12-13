package day2_test

import (
	"github.com/hariskhan14/advent-of-code-2021/day2"
	"testing"
)

func Test_Dive(t *testing.T) {
	t.Run("easy forward and down command", func(t *testing.T) {
		commands := []string{"forward 1", "down 1"}
		expected := 1

		got := day2.LetsDive(commands)

		if got != expected {
			t.Errorf("got %d, expected: %d, input: %v", got, expected, commands)
		}
	})

	t.Run("increase depth in same command", func(t *testing.T) {
		commands := []string{"forward 1", "down 2"}
		expected := 2

		got := day2.LetsDive(commands)

		if got != expected {
			t.Errorf("got %d, expected: %d, input: %v", got, expected, commands)
		}
	})

	t.Run("increasing both forward and depth", func(t *testing.T) {
		commands := []string{"forward 2", "down 3"}
		expected := 6

		got := day2.LetsDive(commands)

		if got != expected {
			t.Errorf("got %d, expected: %d, input: %v", got, expected, commands)
		}
	})
}
