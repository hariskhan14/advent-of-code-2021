package day2_test

import (
	"github.com/hariskhan14/advent-of-code-2021/day2"
	"testing"
)

func Test_Dive(t *testing.T) {
	t.Run("1 forward, 1 down", func(t *testing.T) {
		commands := []string{"forward 1", "down 1"}
		expected := 1

		got := day2.LetsDive(commands)

		if got != expected {
			t.Errorf("got %d, expected: %d, input: %v", got, expected, commands)
		}
	})

	t.Run("1 forward, 2 down", func(t *testing.T) {
		commands := []string{"forward 1", "down 2"}
		expected := 2

		got := day2.LetsDive(commands)

		if got != expected {
			t.Errorf("got %d, expected: %d, input: %v", got, expected, commands)
		}
	})
}
