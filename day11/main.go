package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sethlittleford/aoc-2021/utils"
)

func main() {
	rows := utils.ReadStrings("day11/input.txt")
	
	grid1 := NewGrid()
	grid1.populateGrid(rows)

	fmt.Printf("Part 1 Answer: %d\n", part1(grid1))

	grid2 := NewGrid()
	grid2.populateGrid(rows)

	fmt.Printf("Part 2 Answer: %d\n", part2(grid2))
}

type grid struct {
	g [][]*octopus
}

func NewGrid() *grid {
	return &grid{g: make([][]*octopus, 10)}
}

// populateGrid adds octopuses with energy levels given
// by the input file to the 10x10 grid
func (g *grid) populateGrid(rows []string) {
	for y, row := range rows {
		r := make([]*octopus, 10)
		for x, col := range strings.Split(row, "") {
			energy, err := strconv.Atoi(col)
			utils.CheckErr(err)
			r[x] = &octopus{energy, false}
		}
		g.g[y] = r
	}
}

// printGrid is a helper method that prints the grid
// of octopus energy levels in human-readable format
func (g *grid) printGrid() {
	for _, row := range g.g {
		for _, octopus := range row {
			fmt.Printf("%d", octopus.energy)
		}
		fmt.Println()
	}
}

type octopus struct {
	energy     int
	hasFlashed bool
}

// flash marks the octpus as flashed, then increments the
// energy levels of all neighboring octopuses within bounds
func (o *octopus) flash(grid *grid, row, col int) {
	// mark as flashed
	o.hasFlashed = true
	// increment energy levels of neighboring octopuses
	if row > 0 {
		grid.g[row-1][col].energy++
		if col > 0 {
			grid.g[row-1][col-1].energy++
		}
		if col < len(grid.g[0])-1 {
			grid.g[row-1][col+1].energy++
		}
	}
	if row < len(grid.g)-1 {
		grid.g[row+1][col].energy++
		if col < len(grid.g[0])-1 {
			grid.g[row+1][col+1].energy++
		}
		if col > 0 {
			grid.g[row+1][col-1].energy++
		}
	}
	if col < len(grid.g[0])-1 {
		grid.g[row][col+1].energy++
	}
	if col > 0 {
		grid.g[row][col-1].energy++
	}
}

func part1(grid *grid) int {
	// model energy levels and light in steps
	var flashes int
	for i := 0; i < 100; i++ {
		flashes += step(grid)
	}
	// grid.printGrid()

	return flashes
}

func part2(grid *grid) int {
	// find step where all octopuses flash simultaneously
	// note that this is guaranteed to occur given this input
	var steps int
	for {
		steps++
		if step(grid) == 100 {
			return steps
		}
	}
}

// step performs a single step forward in the octopus model,
// executing the algorithm while keeping track of octopus flashes
func step(grid *grid) int {
	// keep track of the number of flashes that occurred in this step
	var flashes int
	
	// energy level of each octopus increases by 1
	for _, row := range grid.g {
		for _, octopus := range row {
			octopus.energy++
		}
	}
	// any octopus with an energy level greater than 9 flashes
	var restart, stepComplete bool
	for !stepComplete {
		for y, row := range grid.g {
			for x, octopus := range row {
				if octopus.energy > 9 && !octopus.hasFlashed {
					// flash
					octopus.flash(grid, y, x)
					flashes++
					restart = true
					break
				}
			}
			if restart {
				restart = false
				break
			}
			if y == len(grid.g)-1 {
				stepComplete = true
			}
		}
	}
	// any octopus that flashed, reset hasFlashed & set energy level to 0
	for _, row := range grid.g {
		for _, octopus := range row {
			if octopus.hasFlashed {
				octopus.hasFlashed = false
				octopus.energy = 0
			}
		}
	}
	return flashes
}
