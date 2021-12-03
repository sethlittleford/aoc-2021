package main

import (
	"fmt"

	"github.com/sethlittleford/aoc-2021/utils"
)

func main() {
	report := utils.ReadInts("day1/input.txt")

	fmt.Printf("Part 1 Answer: %d\n", part1(report))
	fmt.Printf("Part 2 Answer: %d\n", part2(report))
}

func part1(input []int) int {

	var answer int
	var previousDepth int
	for _, depthMeasurement := range input {
		if depthMeasurement > previousDepth {
			answer++
		}

		previousDepth = depthMeasurement
	}

	answer -= 1

	return answer
}

func part2(measurements []int) int {
	windowSums := make([]int, 0)
	for i := 1; i < len(measurements)-1; i++ {
		currentSum := measurements[i-1] + measurements[i] + measurements[i+1]
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
