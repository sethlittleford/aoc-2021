package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sethlittleford/aoc-2021/utils"
)

type point struct {
	x int
	y int
}

func main() {
	input := utils.ReadStrings("day5/input.txt")

	fmt.Printf("Part 1 Answer: %d\n", part1(input))
	fmt.Printf("Part 2 Answer: %d\n", part2(input))
}

func part1(input []string) int {
	var coveredPoints []point

	for _, line := range input {
		x1, y1, x2, y2 := getXYCoordinates(line)

		coveredPoints = append(coveredPoints, coveredCoordinates(x1, x2, y1, y2, false)...)
	}

	return countCovered(coveredPoints)
}

func part2(input []string) int {
	var coveredPoints []point

	for _, line := range input {
		x1, y1, x2, y2 := getXYCoordinates(line)

		coveredPoints = append(coveredPoints, coveredCoordinates(x1, x2, y1, y2, true)...)
	}

	return countCovered(coveredPoints)
}

// getXYCoordinates parses an input line into
// x1, y1, x2, and y2 coordinates
func getXYCoordinates(line string) (int, int, int, int) {
	coordinates := strings.Split(line, " -> ")
	// x1 and y1
	begin := strings.Split(coordinates[0], ",")
	x1, err := strconv.Atoi(begin[0])
	utils.CheckErr(err)
	y1, err := strconv.Atoi(begin[1])
	utils.CheckErr(err)
	// x2 and y2
	end := strings.Split(coordinates[1], ",")
	x2, err := strconv.Atoi(end[0])
	utils.CheckErr(err)
	y2, err := strconv.Atoi(end[1])
	utils.CheckErr(err)

	return x1, y1, x2, y2
}

// coveredCoordinates calculates and returns all points
// over which the line covers
func coveredCoordinates(x1, x2, y1, y2 int, calculateDiagonal bool) []point {
	var coveredPoints []point

	if x1 == x2 {
		// get coordinates of points the line covers
		if y1 < y2 {
			for i := y1; i <= y2; i++ {
				coveredPoints = append(coveredPoints, point{x1, i})
			}
		} else if y1 > y2 {
			for i := y1; i >= y2; i-- {
				coveredPoints = append(coveredPoints, point{x1, i})
			}
		} else {
			// x1 = x2 and y1 = y2 so line covers one point
			coveredPoints = append(coveredPoints, point{x1, y1})
		}
	} else if y1 == y2 {
		if x1 < x2 {
			for i := x1; i <= x2; i++ {
				coveredPoints = append(coveredPoints, point{i, y1})
			}
		} else if x1 > x2 {
			for i := x1; i >= x2; i-- {
				coveredPoints = append(coveredPoints, point{i, y1})
			}
		} else {
			// y1 = y2 and x1 = x2 so line covers one point
			coveredPoints = append(coveredPoints, point{x1, y1})
		}
	} else {
		// line is neither horizontal, nor vertical but diagonal at 45 degrees
		if !calculateDiagonal {
			return nil
		}

		// calculate x coordinates
		var xCords []int
		if x1 < x2 {
			for i := x1; i <= x2; i++ {
				xCords = append(xCords, i)
			}
		} else {
			for i := x1; i >= x2; i-- {
				xCords = append(xCords, i)
			}
		}
		// calculate y coordinates
		var yCords []int
		if y1 < y2 {
			for i := y1; i <= y2; i++ {
				yCords = append(yCords, i)
			}
		} else {
			for i := y1; i >= y2; i-- {
				yCords = append(yCords, i)
			}
		}

		// create points 45 deg line covers
		for i := 0; i < len(xCords); i++ {
			coveredPoints = append(coveredPoints, point{xCords[i], yCords[i]})
		}
	}
	return coveredPoints
}

// countCovered tracks the number of ocurrances of
// each point being covered by a line and returns
// the number of points where at least two lines overlap
func countCovered(coveredPoints []point) int {
	// two-dimensional matrix to store count of all points
	count := make([][]int, 1000)
	for i := range count {
		count[i] = make([]int, 1000) // cols
	}

	// populate count of covered points matrix
	for _, point := range coveredPoints {
		count[point.y][point.x]++
	}

	// return number of points where at least two lines overlap
	var result int
	for _, row := range count {
		for _, col := range row {
			if col >= 2 {
				result++
			}
		}
	}
	return result
}
