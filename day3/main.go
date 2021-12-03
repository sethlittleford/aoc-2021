package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	absPath, err := filepath.Abs("day3/input.txt")
	check(err)

	// examplePath, err := filepath.Abs("day3/example.txt")
	// check(err)

	fmt.Printf("Part 1 Answer: %d\n", part1(absPath))
	fmt.Printf("Part 2 Answer: %d\n", part2(absPath))
}

func part1(filePath string) int64 {

	gamma := calculateGammaRate(filePath)
	epsilon := calculateEpsilonRate(gamma)

	// convert gamma and epsilon rates to decimal numbers
	gammaInt, err := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	check(err)
	epsilonInt, err := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)
	check(err)

	return gammaInt * epsilonInt
}

func calculateGammaRate(filePath string) []string {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	countOf1s := make([]int, 12)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
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

func part2(filePath string) int64 {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	inputMatrix := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		inputMatrix = append(inputMatrix, line)
	}

	ox, _ := oxygenGenRate(inputMatrix, 0)
	c2, _ := co2ScrubberRate(inputMatrix, 0)

	// convert to decimal
	oxygen, err := strconv.ParseInt(strings.Join(ox[0], ""), 2, 64)
	check(err)
	co2, err := strconv.ParseInt(strings.Join(c2[0], ""), 2, 64)
	check(err)

	return oxygen * co2
}

func oxygenGenRate(input [][]string, index int) ([][]string, int) {
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
			countOf1s ++
		} else {
			countOf0s ++
		}
	}
	if countOf1s >= countOf0s {
		// keep lines with the most common bit, 1
		filtered = filterLines(input, "1", index)
	} else {
		// keep lines with the most common bit, 0
		filtered = filterLines(input, "0", index)
	}

	return oxygenGenRate(filtered, index + 1)
}

func co2ScrubberRate(input [][]string, index int) ([][]string, int) {
	// base case
	if len(input) == 1 {
		return input, index
	}
	var filtered [][]string
	// find least common bit
	var countOf1s int
	var countOf0s int
	for _, line := range input {
		if line[index] == "1" {
			countOf1s ++
		} else {
			countOf0s ++
		}
	}
	if countOf0s <= countOf1s {
		// keep lines with the least common bit, 0
		filtered = filterLines(input, "0", index)
	} else {
		// keep lines with the least common bit, 1
		filtered = filterLines(input, "1", index)
	}

	return co2ScrubberRate(filtered, index + 1)
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
