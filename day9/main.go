package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/sethlittleford/aoc-2021/utils"
)

func main() {
	rows := utils.ReadStrings("day9/input.txt")

	fmt.Printf("Part 1 Answer: %d\n", part1(rows))
	fmt.Printf("Part 2 Answer: %d\n", part2(rows))

}

func part1(rows []string) int {
	heightMap := make([][]int, 0)
	for _, row := range rows {
		r := make([]int, 0)
		for _, col := range strings.Split(row, "") {
			height, err := strconv.Atoi(col)
			utils.CheckErr(err, "could not convert string to int")
			r = append(r, height)
		}
		heightMap = append(heightMap, r)
	}

	// find the low points on the heightmap

	var up, down, left, right int
	lowPoints := make([]int, 0)
	// iterate through heightmap
	for y, row := range heightMap {
		for x, current := range row {
			// find adjacent locations in the matrix

			// left & right
			if x == 0 {
				// left-most column
				left = 10
				right = row[x+1]
			} else if x == len(row)-1 {
				// right-most column
				right = 10
				left = row[x-1]
			} else {
				// has left & right
				left = row[x-1]
				right = row[x+1]
			}
			// up & down
			if y == 0 {
				// top-most row
				up = 10
				down = heightMap[y+1][x]
			} else if y == len(heightMap)-1 {
				// bottom-most row
				down = 10
				up = heightMap[y-1][x]
			} else {
				// has up & down
				up = heightMap[y-1][x]
				down = heightMap[y+1][x]
			}

			// fmt.Printf(" %d \n%d %d %d\n %d \n\n", up, left, current, right, down)

			// low points are lower than any adjacent height
			if current < up && current < down && current < left && current < right {
				lowPoints = append(lowPoints, current)
			}
		}
	}
	// calculate risk level
	var riskLevel int
	for _, point := range lowPoints {
		riskLevel += point + 1
	}
	return riskLevel
}

type point struct {
	height   int
	visited  bool
	basinNum int
}

type heightMap struct {
	vals [][]*point
}

func NewHeightMap() *heightMap {
	return &heightMap{vals: make([][]*point, 0)}
}

func part2(rows []string) int {
	heightMap := NewHeightMap()
	for _, row := range rows {
		r := make([]*point, 0)
		for _, col := range strings.Split(row, "") {
			height, err := strconv.Atoi(col)
			utils.CheckErr(err, "could not convert string to int")
			r = append(r, &point{height, false, -1})
		}
		heightMap.vals = append(heightMap.vals, r)
	}

	// calculate basins in heightMap, assign each basin an arbitrary but unique basin number
	basinNumber := 0
	for y, row := range heightMap.vals {
		for x := range row {
			if !heightMap.vals[y][x].visited {
				basinNumber++
			}
			calcBasin(heightMap, x, y, basinNumber)
		}
	}

	// calculate the size of each basin
	basinSizes := make(map[int]int)
	for _, row := range heightMap.vals {
		for _, point := range row {
			if point.basinNum != -1 {
				basinSizes[point.basinNum]++
			}
		}
	}

	// find the three largest basins and multiply their sizes together
	sizes := make([]int, 0)
	for _, size := range basinSizes {
		sizes = append(sizes, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	return sizes[0] * sizes[1] * sizes[2]
}

// calcBasin starts at a point, then recursively traverses outward in all
// four directions, assigning each eligible point the same basin number
func calcBasin(heightMap *heightMap, x, y, currentBasinNum int) {
	// base case
	if indexOutOfBounds(heightMap, x, y) || isNotEligible(heightMap.vals[y][x]) {
		return
	}

	// mark the current point as visited
	heightMap.vals[y][x].visited = true

	// assign the current point the current basin number
	heightMap.vals[y][x].basinNum = currentBasinNum

	// traverse up, assign all eligible points the current basin number
	calcBasin(heightMap, x, y-1, currentBasinNum)
	// traverse down, assign all eligible points the current basin number
	calcBasin(heightMap, x, y+1, currentBasinNum)
	// traverse left, assign all eligible points the current basin number
	calcBasin(heightMap, x-1, y, currentBasinNum)
	// traverse right, assign all eligible points the current basin number
	calcBasin(heightMap, x+1, y, currentBasinNum)
}

// indexOutOfBounds returns true if the current coordinate lies outside
// the bounds of the given heightMap
func indexOutOfBounds(heightMap *heightMap, x, y int) bool {
	// top, bottom, left, right
	if y == -1 || y > len(heightMap.vals)-1 || x == -1 || x > len(heightMap.vals[0])-1 {
		return true
	}
	return false
}

// isNotEligible returns true if the point's height is 9 or if
// the point has already been visited in basin calculation
func isNotEligible(p *point) bool {
	if p.height == 9 || p.visited {
		return true
	}
	return false
}
