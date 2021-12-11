package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/sethlittleford/aoc-2021/utils"
)

func main() {
	input := utils.ReadStrings("day10/input.txt")

	fmt.Printf("Part 1 Answer: %d\n", part1(input))
	fmt.Printf("Part 2 Answer: %d\n", part2(input))
}

type stack struct {
	chars []string
}

func NewStack() *stack {
	return &stack{chars: make([]string, 0)}
}

func (s *stack) push(char string) {
	s.chars = append(s.chars, char)
}

func (s *stack) pop() (string, error) {
	if len(s.chars) == 0 {
		return "", errors.New("Cannot pop(), the stack is empty")
	}

	defer func() {
		s.chars[len(s.chars)-1] = ""
		s.chars = s.chars[:len(s.chars)-1]
	}()
	return s.chars[len(s.chars)-1], nil
}

type char struct {
	compliment string
	points     int
}

var pairs map[string]char = map[string]char{
	")": char{"(", 3},
	"]": char{"[", 57},
	"}": char{"{", 1197},
	">": char{"<", 25137},
}

func part1(input []string) int {
	// store map of accumulated syntax error score of each character
	errPoints := make(map[string]int)
	// find the corrupted lines
	for _, line := range input {
		s := NewStack()
		for _, char := range strings.Split(line, "") {
			// if char is an opener, push to stack
			if char == "(" || char == "[" || char == "{" || char == "<" {
				s.push(char)
			} else {
				// char is a closer, its complimentary opener should be at the top of the stack
				opener, err := s.pop()
				utils.CheckErr(err)

				if opener != pairs[char].compliment {
					// this line is corrupted
					// calculate the syntax error score for the corrupted line
					errPoints[char] += pairs[char].points
				}
			}
		}
	}

	// calculate the sum of all syntax error scores
	var totalScore int
	for _, score := range errPoints {
		totalScore += score
	}
	return totalScore
}

func part2(input []string) int {
	oPairs := map[string]char{
		"(": char{")", 1},
		"[": char{"]", 2},
		"{": char{"}", 3},
		"<": char{">", 4},
	}

	// find the incomplete lines
	lineScores := make([]int, 0)
	for _, line := range input {
		s := NewStack()
		var isCorrupted bool
		for _, char := range strings.Split(line, "") {
			// if char is an opener, push to stack
			if char == "(" || char == "[" || char == "{" || char == "<" {
				s.push(char)
				continue
			}
			// char is a closer, its complimentary opener should be at the top of the stack
			opener, err := s.pop()
			utils.CheckErr(err)

			if opener != pairs[char].compliment {
				// this line is corrupted, skip it
				isCorrupted = true
				break
			}
		}
		if !isCorrupted {
			// this line is incomplete
			// no need for completion string, just calculate line score by popping stack and
			// adding score for complimentary closers
			lineScore := 0
			for i := len(s.chars) - 1; i >= 0; i-- {
				lineScore *= 5
				lineScore += oPairs[s.chars[i]].points
			}
			lineScores = append(lineScores, lineScore)
		}
	}

	// sort the line scores, then take the middle score
	sort.Ints(lineScores)

	return lineScores[len(lineScores) / 2]
}
