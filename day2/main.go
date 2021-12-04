package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sethlittleford/aoc-2021/utils"
)

func main() {
	input := utils.ReadStrings("day2/input.txt")

	fmt.Printf("Part 1 Answer: %d\n", part1(input))
	fmt.Printf("Part 2 Answer: %d\n", part2(input))
}

func part1(input []string) int {
	var horizontalPos int
	var depth int
	for _, v := range input {
		direction, magnitude, err := submarineCommand(v)
		utils.CheckErr(err, "failed to parse submarine command")

		// calculate horizontal and depth coordinates
		switch direction {
		case "forward":
			horizontalPos += magnitude
		case "down":
			depth += magnitude
		case "up":
			depth -= magnitude
		}
	}
	return horizontalPos * depth
}

func part2(input []string) int {
	var horizontalPos int
	var depth int
	var aim int
	for _, v := range input {
		direction, magnitude, err := submarineCommand(v)
		utils.CheckErr(err, "failed to parse submarine command")

		// calculate horizontal pos, depth, and aim
		switch direction {
		case "forward":
			horizontalPos += magnitude
			depth += aim * magnitude
		case "down":
			aim += magnitude
		case "up":
			aim -= magnitude
		}
	}
	return horizontalPos * depth
}

// submarineCommand takes in the command from input line
// and returns the direction and magnitude
func submarineCommand(command string) (string, int, error) {
	line := strings.Split(command, " ")
	direction := line[0]
	magnitude, err := strconv.Atoi(line[1])
	if err != nil {
		return "", 0, err
	}

	return direction, magnitude, nil
}
