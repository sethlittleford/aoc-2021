package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sethlittleford/aoc-2021/utils"
)

func main() {
	// digit -> segments
	// 0 -> 6
	// 1 -> 2
	// 2 -> 5
	// 3 -> 5
	// 4 -> 4
	// 5 -> 5
	// 6 -> 6
	// 7 -> 3
	// 8 -> 7
	// 9 -> 6

	input := utils.ReadStrings("day8/input.txt")

	fmt.Printf("Part 1 Answer: %d\n", part1(input))
	fmt.Printf("Part 2 Answer: %d\n", part2(input))
}

func part1(input []string) int {
	outputVals := make([]string, 0)
	for _, line := range input {
		outputVals = append(outputVals, outputValues(line)...)
	}
	
	var easyDigits int
	for _, val := range outputVals {
		switch len(val) {
		case 2:
			// 1
			easyDigits++
		case 3:
			// 7
			easyDigits++
		case 4:
			// 4
			easyDigits++
		case 7:
			// 8
			easyDigits++
		default:
			continue
		}
	}

	return easyDigits
}

func part2(input []string) int {
	var answer int
	for _, line := range input {
		answer += output(line)
	}
	return answer
}

// output decodes the four-digit value
// that should be output on the seven-segment display 
func output(line string) int {
	in := inputSignal(line)

	// map of digits to signal patterns
	m := make(map[int][]string, 10)

	// find simplest cases 1, 4, 7, 8
	for _, val := range in {
		switch len(val) {
		case 2:
			// 1
			m[1] = strings.Split(val, "")
		case 3:
			// 7
			m[7] = strings.Split(val, "")
		case 4:
			// 4
			m[4] = strings.Split(val, "")
		case 7:
			// 8
			m[8] = strings.Split(val, "")
		default:
			continue
		}
	}

	// deductive cases, in order of dependency

	// deduce 3 and 6
	for _, val := range in {
		// 3 has same segments that 1 has
		if len(val) == 5 && contains(val, m[1]) {
			m[3] = strings.Split(val, "")
		}
		// 6 does not have same segments as 1
		if len(val) == 6 && !contains(val, m[1]) {
			m[6] = strings.Split(val, "")
		}
	}

	// deduce 2 and 5 
	for _, val := range in {
		if len(val) == 5 {
			// check for a 5
			if isAFive(val, m[6]) {
				m[5] = strings.Split(val, "")
			} else if contains(val, m[3]) {
				// this is a 3
				continue
			} else {
				// it's a 2
				m[2] = strings.Split(val, "")
			}
		}
	}
	
	// deduce 0 and 9
	for _, val := range in {
		if len(val) == 6 {
			// 9 has same segments as 5 and 1
			if contains(val, m[5]) && contains(val, m[1]) {
				m[9] = strings.Split(val, "")
			} else if contains(val, m[6]) {
				// this is a 6
				continue
			} else {
				// it's a 0
				m[0] = strings.Split(val, "")
			}
		}
	}

	var outputDigit string
	for _, outVal := range outputValues(line) {
			for digit, signal := range m {
				if equals(outVal, signal) {
					outputDigit += strconv.Itoa(digit)
				}
			}	
	}

	out, err := strconv.Atoi(outputDigit)
	utils.CheckErr(err, "could not convert string to int")
	
	return out
}

// outputValues parses an input line and returns all ouput
// to the right of the "|" delimiter
func outputValues(line string) []string {
	outputVals := make([]string, 0)
	for _, v := range strings.Split(strings.Split(line, "|")[1], " ") {
		if len(v) == 0 {
			continue
		}
		outputVals = append(outputVals, v)	
	}
	return outputVals
}

// inputSignal parses an input line and returns all signals
// to the left of the "|" delimiter
func inputSignal(line string) []string {
	input := make([]string, 0)
	for _, v := range strings.Split(strings.Split(line, "|")[0], " ") {
		if len(v) == 0 {
			continue
		}
		input = append(input, v)	
	}
	return input
}

// isAFive takes a signal string and the segments of
// a known 6 digit, and checks if the input signal
// is a 5 digit
func isAFive(s string, sixSegments []string) bool {
	// 5 has all but one segments same as 6
	var segInCommon int
	for _, segment := range sixSegments {
		if strings.Contains(s, segment) {
			segInCommon++
		}
	}
	return segInCommon == len(sixSegments) - 1 
}

// contains checks if the input signal string has all segments
// of a given digit
func contains(s string, segments []string) bool {
	for _, segment := range segments {
		if !strings.Contains(s, segment) {
			return false
		}
	}
	return true
}

// equals checks whether the input signal string is
// exactly equal to a digit's segments
func equals(s string, segments []string) bool {
	if len(s) != len(segments) {
		return false
	}
	if !contains(s, segments) {
		return false
	}
	return true
}
