package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sethlittleford/aoc-2021/utils"
)

func main() {
	input := utils.ReadStrings("day6/input.txt")

	fmt.Printf("Part 1 Answer: %d\n", part1(input))
	fmt.Printf("Part 2 Answer: %d\n", part2(input))
}

func part1(input []string) int {
	fish := parseFishTimers(input)

	// Model exponential fish growth rate over 80 days (naive implementation, too much memory for part 2!)
	for day := 1; day <= 80; day++ {
		// iterate over all fish
		for i, timer := range fish {
			// if timer > 0, decrement by 1 day
			if timer > 0 {
				fish[i] --
				continue
			}
			// if timer == 0, reset to 6 days
			fish[i] = 6
			// append a new fish age of 8 days
			fish = append(fish, 8)
		}
	}

	// total lanternfish after 80 days of population growth
	return len(fish)
}

func part2(input []string) int {
	fish := parseFishTimers(input)

	// The count array stores how many fish there are in each internal timer group (0 - 8 days)
	count := make([]int, 9) // timers from 0 to 8 days

	// Populate the initial count of fish in each timer group
	for i := 0; i < 9; i++ {
		var numFish int
		for _, timer := range fish {
			if timer == i {
				numFish ++
			}
		}
		count[i] = numFish
	}

	// Model exponential fish growth rate over 256 days (memory efficient implementation!)
	for day := 1; day <= 256; day++ {
		// the nextGen array stores the fish count of the new generation
		nextGen := make([]int, 0)
		// timer > 0 decrements by moving left one index
		nextGen = append(nextGen, count[1:]...)
		// time == 0 creates new fish with timer of 8
		nextGen = append(nextGen, count[0]) 
		// time == 0 also gets reset with timer of 6
		nextGen[6] += count[0]
		// update count array after each generation
		count = nextGen
	}

	// total lanternfish after 256 days of population growth
	var totalFish int
	for _, timer := range count {
		totalFish += timer
	}

	return totalFish
}

// parseFishTimers creates the slice of all fish timers as ints
func parseFishTimers(input []string) []int {
	timerStrs := strings.Split(input[0], ",")

	timers := make([]int, 0)
	for _, timerStr := range timerStrs {
		timer, err := strconv.Atoi(timerStr)
		utils.CheckErr(err, "failed to convert string to int")
		timers = append(timers, timer)
	}
	return timers
}