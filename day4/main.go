package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sethlittleford/aoc-2021/utils"
)

func main() {
	input := utils.ReadStrings("day4/input.txt")
	// input := utils.ReadStrings("day4/example.txt")

	selections, boards := parseSelectionsAndBoards(input)

	fmt.Printf("Part 1 Answer: %d\n", part1(selections, boards))
	fmt.Printf("Part 2 Answer: %d\n", part2(selections, boards))

}

type board struct {
	matrix   [][]int
	isWinner bool
}

// mark takes the selected number and "marks"
// the number on the board, if present, by setting
// the space to 100
func (b board) mark(selection int) {
	for _, row := range b.matrix {
		for x, col := range row {
			if selection == col {
				row[x] = 100 // marked = 100
			}
		}
	}
}

// hasBingo checks if the board has 5 "marked"
// spaces in a row or a column and returns true if so
func (b board) hasBingo() bool {
	// check rows
	for _, row := range b.matrix {
		if sum(row) == 5*100 {
			// this board is a winner!
			return true
		}
	}
	// check columns
	for i := 0; i < 5; i++ {
		var colSum int
		for _, row := range b.matrix {
			colSum += row[i]
		}
		if colSum == 5*100 {
			// this board is a winnder!
			return true
		}
	}
	return false
}

// unmarkedSum calculates the sum of all "unmarked"
// spaces on the board
func (b board) unmarkedSum() int {
	var sum int
	for _, row := range b.matrix {
		for _, num := range row {
			if num != 100 {
				sum += num
			}
		}
	}
	return sum
}

func sum(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

// parseSelectionsAndBoards creates the bingo selection numbers
// as a slice of int and the boards as a slice of matrices
func parseSelectionsAndBoards(input []string) ([]int, []board) {
	var selectionsS string
	var sboard board
	var boards []board
	for i, v := range input {
		if i == 0 {
			selectionsS = v
		} else if len(v) == 0 {
			continue
		} else {
			row := make([]int, 0)
			for _, v := range strings.Split(v, " ") {
				if len(v) == 0 {
					continue
				}
				i, err := strconv.Atoi(v)
				utils.CheckErr(err, "failed to convert string to int")
				row = append(row, i)
			}
			sboard.matrix = append(sboard.matrix, row)
			if len(sboard.matrix) == 5 {
				boards = append(boards, sboard)
				// reset the board matrix
				sboard.matrix = nil
			}
		}
	}

	selections := make([]int, 0)
	for _, v := range strings.Split(selectionsS, ",") {
		selection, err := strconv.Atoi(v)
		utils.CheckErr(err, "failed to convert string to int")
		selections = append(selections, selection)
	}

	return selections, boards
}

func part1(selections []int, boards []board) int {
	// mark boards while checking for bingo
	for i, n := range selections {
		// mark boards
		for _, board := range boards {
			board.mark(n)
		}

		// check for bingo if >= 5 nums drawn
		if i >= 5 {
			for _, board := range boards {
				if board.hasBingo() {
					return board.unmarkedSum() * n
				}
			}
		}
	}

	return 0
}


func part2(selections []int, boards []board) int {
	var lastBoard board // losing board
	// mark boards while checking for bingo
	for i, n := range selections {
		// mark boards
		for _, board := range boards {
			board.mark(n)
		}

		// check for bingo if >= 5 nums drawn
		if i >= 5 {
			for m, board := range boards {
				if board.hasBingo() {
					// mark the board as a winner
					boards[m].isWinner = true
				}
			}
		}

		// check for last board
		var numWinB int
		for _, board := range boards {
			if board.isWinner {
				numWinB ++
			}
		}
		if numWinB == len(boards) - 1 {
			// find losing board
			for _, board := range boards {
				if !board.isWinner {
					lastBoard = board
				}
			}
		}
		if numWinB == len(boards) {
			return lastBoard.unmarkedSum() * n
		}
	}

	return 0
}
