package day

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Day4 struct{}

func (d *Day4) Part1(input string) string {
	puzzleInput := strings.Split(input, "\n")
	orders := strings.Split(puzzleInput[0], ",")
	boards := getBingoBoards(input)

	var marked map[int][]string
	var lastDrawn string
	for _, order := range orders {
		marked, lastDrawn = markBoards(boards, order)

		if marked != nil {
			break
		}
	}

	sumOfUnmarkedNumbers := sumOfUnmarkedNumbers(marked)
	numLastDrawn, err := strconv.Atoi(lastDrawn)

	if err != nil {
		fmt.Printf("Could not convert string [%v] to int\n", lastDrawn)
	}

	return fmt.Sprintf("%d", (sumOfUnmarkedNumbers * numLastDrawn))
}

func (d *Day4) Part2(input string) string {
	return input
}

func getBingoBoards(input string) map[int]map[int][]string {
	re := regexp.MustCompile(`(\d+)`)
	puzzleRows := strings.Split(input, "\n")
	boards := make(map[int]map[int][]string, 0)

	boardIndex := 2
	i := 0

	for boardIndex < len(puzzleRows) {
		tempRow := make(map[int][]string)
		for x := 0; x < 5; x++ {
			tempRow[x] = re.FindAllString(puzzleRows[boardIndex+x], -1)
		}
		boards[i] = tempRow
		i++

		boardIndex += 6
	}

	return boards
}

func markBoards(boards map[int]map[int][]string, draw string) (map[int][]string, string) {
	for x := 0; x < len(boards); x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < 5; z++ {
				elem := boards[x][y][z]

				if elem == draw && !strings.Contains(elem, "*") {
					boards[x][y][z] = elem + "*"
				}
			}

			if checkRowMatch(boards[x][y]) {
				fmt.Printf("Column match: %v\n", boards[x][y])
				return boards[x], draw
			}
		}

		if checkColumnMatch(boards[x]) {
			return boards[x], draw
		}
	}

	return nil, draw
}

func checkRowMatch(row []string) bool {
	re := regexp.MustCompile(`(\*)`)
	result := re.FindAllString(strings.Join(row, " "), -1)

	return len(result) == 5
}

func checkColumnMatch(board map[int][]string) bool {
	for i := 0; i < 5; i++ {
		tempColumn := make([]string, 0)

		for j := 0; j < 5; j++ {
			tempColumn = append(tempColumn, board[j][i])
		}

		if checkRowMatch(tempColumn) {
			fmt.Printf("Column match: %v\n", tempColumn)
			return true
		}
	}

	return false
}

func sumOfUnmarkedNumbers(board map[int][]string) int {
	sum := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			elem := board[i][j]
			if !strings.Contains(elem, "*") {
				numElem, err := strconv.Atoi(elem)

				if err != nil {
					fmt.Printf("Could not convert string [%s] to number\n", elem)
					os.Exit(2)
				}
				sum += numElem
			}
		}
	}

	return sum
}
