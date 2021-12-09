package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/sethlittleford/aoc-2021/utils"
)

func main() {
	input := utils.ReadStrings("day7/input.txt")
	positions := getPositions(input)

	fmt.Printf("Part 1 Answer: %d\n", part1(positions))
	fmt.Printf("Part 2 Answer: %d\n", part2(positions))
}

func part1(positions []float64) int {
	// align all crabs on same horizontal position
	// using the least fuel possible (using median)

	median := median(positions)

	// calculate least fuel cost to move to median
	var fuel int
	for _, position := range positions {
		fuel += int(math.Abs(position - median))
	}

	return fuel
}

// median calculates the median of the finite data set
func median(data []float64) float64 {
	// sort the values in data set in increasing order
	sort.Float64s(data)

	// calculate the median
	n := len(data)
	if n % 2 != 0 {
		return data[n / 2]
	}
	return (data[n / 2] + data[(n + 1) / 2]) / 2	
}

func part2(positions []float64) int {
	// align all crabs on same horizontal position
	// using the least fuel possible (using mean)

	mean := mean(positions)

	meanC := math.Ceil(mean)
	meanF := math.Floor(mean)

	// calculate least fuel cost to move to mean (ceiling vs floor)
	var fuelC, fuelF int
	for _, position := range positions {
		// mean ceiling
		for i := 1; i <= int(math.Abs(position - meanC)); i++ {
			fuelC += i
		}
		// mean floor
		for i := 1; i <= int(math.Abs(position - meanF)); i++ {
			fuelF += i
		}		
	}

	if fuelC < fuelF {
		return fuelC
	}

	return fuelF
}

// mean calculates the arithmetic mean of the finite data set 
func mean(data []float64) float64 {
	var sum, count float64
	for _, v := range data {
		count++
		sum += v
	}
	return sum / count
}

// getPositions parses the input string into crab positions
// as a slice of float64
func getPositions(input []string) []float64 {
	positions := make([]float64, 0)
	for _, pos := range strings.Split(input[0], ",") {
		pos, err := strconv.ParseFloat(pos, 64)
		utils.CheckErr(err)
		positions = append(positions, pos)
	}
	return positions
}
