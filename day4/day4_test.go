package day4

import (
	"fmt"
	"github.com/hariskhan14/advent-of-code-2021/utils"
	"strconv"
	"strings"
	"testing"
)

type Board [][]int

func TestGiantSquid(t *testing.T) {
	drawnNumbers, boards := getSquidInput("sample_input.txt")
	fmt.Printf("%v", drawnNumbers)
	fmt.Printf("%v", boards)
}

func getSquidInput(filename string) ([]int, []Board) {
	lines := utils.ReadFileAsStrings(filename)
	if len(lines) == 0 {
		return nil,  nil
	}

	drawnNumbers := getDrawnNumbers(lines)
	boards := getBoards(lines)

	return drawnNumbers, boards
}

func getBoards(lines []string) []Board{
	boards := []Board{make(Board, 5)}
	boardIdx := 0
	lineIndex := 0
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}

		nums := strings.Split(line, " ")
		i := 0
		for _, num := range nums {
			if num == "" {
				continue
			}

			if len(boards)-1 < boardIdx {
				boards = append(boards, make(Board, 5))
			}

			currentIdx := lineIndex % 5
			if boards[boardIdx][currentIdx] == nil {
				boards[boardIdx][currentIdx] = make([]int, 5)
			}

			boards[boardIdx][currentIdx][i], _ = strconv.Atoi(num)

			if currentIdx == 4 && i == 4 {
				boardIdx++

			}
			i++
		}

		lineIndex++
	}

	return boards
}

func getDrawnNumbers(lines []string) []int {
	drawnNumbersLine := strings.Split(lines[0], ",")
	var drawnNumbers []int
	for _, val := range drawnNumbersLine {
		number, _ := strconv.Atoi(val)
		drawnNumbers = append(drawnNumbers, number)
	}
	return drawnNumbers
}