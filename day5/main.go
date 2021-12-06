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
	// parse input, only consider horizontal and vertical lines
	input := utils.ReadStrings("day5/input.txt")

	fmt.Printf("Part 1 Answer: %d\n", part1(input))
	fmt.Printf("Part 2 Answer: %d\n", part2(input))
}

func part1(input []string) int {
	var coveredPoints []point
	// convert input lines to x and y coordinates
	for _, line := range input {
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
			continue
		}
	}

	// two-dimensional matrix, count of all points
	count := make([][]int, 1000)
	for i := range count {
		count[i] = make([]int, 1000) // cols
	}

	// populate count of points matrix
	for _, point := range coveredPoints {
		count[point.y][point.x] ++
	}

	// return number of points where at least two lines overlap
	var result int
	for _, row := range count {
		for _, col := range row {
			if col >= 2 {
				result ++
			}
		}
	}

	return result
}

func part2(input []string) int {
	var coveredPoints []point
	// convert input lines to x and y coordinates
	for _, line := range input {
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
			// diagonal line at 45 degrees

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
			var yCords []int
			// calculate y coordinates
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
	}

	// two-dimensional matrix, count of all points
	count := make([][]int, 1000)
	for i := range count {
		count[i] = make([]int, 1000) // cols
	}

	// populate count of points matrix
	for _, point := range coveredPoints {
		count[point.y][point.x] ++
	}

	// return number of points where at least two lines overlap
	var result int
	for _, row := range count {
		for _, col := range row {
			if col >= 2 {
				result ++
			}
		}
	}

	return result
}