package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sethlittleford/aoc-2021/utils"
)

func main() {
	input := utils.ReadStrings("day3/input.txt")

	// exampleInput := utils.ReadStrings("day3/input.txt")

	fmt.Printf("Part 1 Answer: %d\n", part1(input))
	fmt.Printf("Part 2 Answer: %d\n", part2(input))
}

func part1(input []string) int64 {
	gamma := calculateGammaRate(input)
	epsilon := calculateEpsilonRate(gamma)

	// convert gamma and epsilon rates to decimal numbers
	gammaInt, err := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	utils.CheckErr(err)
	epsilonInt, err := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)
	utils.CheckErr(err)

	return gammaInt * epsilonInt
}

func calculateGammaRate(input []string) []string {
	countOf1s := make([]int, 12)
	for _, v := range input {
		line := strings.Split(v, "")
		for i, b := range line {
			if b == "1" {
				countOf1s[i] += 1
			}
		}
	}

	gammaRate := make([]string, 12)
	for i, b := range countOf1s {
		if b > 500 {
			// the 1 bit is the most common
			gammaRate[i] = "1"
		} else {
			// the 0 bit is the most common
			gammaRate[i] = "0"
		}
	}

	return gammaRate
}

func calculateEpsilonRate(gammaRate []string) []string {
	epsilonRate := make([]string, 12)
	for i, b := range gammaRate {
		if b == "0" {
			epsilonRate[i] = "1"
		} else {
			epsilonRate[i] = "0"
		}
	}
	return epsilonRate
}

func part2(input []string) int64 {
	inputMatrix := make([][]string, 0)
	for _, v := range input {
		line := strings.Split(v, "")
		inputMatrix = append(inputMatrix, line)
	}

	ox, _ := lifeSupport(inputMatrix, 0, "oxygen")
	c2, _ := lifeSupport(inputMatrix, 0, "co2")

	// convert to decimal
	oxygen, err := strconv.ParseInt(strings.Join(ox[0], ""), 2, 64)
	utils.CheckErr(err)
	co2, err := strconv.ParseInt(strings.Join(c2[0], ""), 2, 64)
	utils.CheckErr(err)

	return oxygen * co2
}

func lifeSupport(input [][]string, index int, param string) ([][]string, int) {
	// base case
	if len(input) == 1 {
		return input, index
	}
	var filtered [][]string

	// find most common bit
	var countOf1s int
	var countOf0s int
	for _, line := range input {
		if line[index] == "1" {
			countOf1s++
		} else {
			countOf0s++
		}
	}
	switch param {
	case "oxygen":
		if countOf1s >= countOf0s {
			// keep lines with the most common bit, 1
			filtered = filterLines(input, "1", index)
		} else {
			// keep lines with the most common bit, 0
			filtered = filterLines(input, "0", index)
		}
	case "co2":
		if countOf0s <= countOf1s {
			// keep lines with the least common bit, 0
			filtered = filterLines(input, "0", index)
		} else {
			// keep lines with the least common bit, 1
			filtered = filterLines(input, "1", index)
		}
	}

	return lifeSupport(filtered, index+1, param)
}

func filterLines(input [][]string, common string, index int) [][]string {
	filtered := make([][]string, 0)
	for _, line := range input {
		if line[index] == common {
			filtered = append(filtered, line)
		}
	}
	return filtered
}
