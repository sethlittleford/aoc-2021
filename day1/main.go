package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	absPath, err := filepath.Abs("day1/input.txt")
	check(err)

	fmt.Printf("Part 1 Answer: %d\n", part1(absPath))
	fmt.Printf("Part 2 Answer: %d\n", part2(absPath))
}

func part1(filePath string) int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var answer int
	var previousDepth int
	for scanner.Scan() {
		depthMeasurement, err := strconv.Atoi(scanner.Text())
		check(err)

		if depthMeasurement > previousDepth {
			answer++
		}

		previousDepth = depthMeasurement
	}

	answer -= 1

	return answer
}

func part2(filePath string) int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	measurements := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depthMeasurement, err := strconv.Atoi(scanner.Text())
		check(err)

		measurements = append(measurements, depthMeasurement)
	}

	windowSums := make([]int, 0)
	for i := 1; i < len(measurements) - 1; i++ {
		currentSum := measurements[i - 1] + measurements[i] + measurements[i + 1]
		windowSums = append(windowSums, currentSum)
	}

	var previousSum int
	var answer int
	for _, currentSum := range windowSums {
		if currentSum > previousSum {
			answer++
		} 
		previousSum = currentSum
	}
	
	answer -= 1

	return answer
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
