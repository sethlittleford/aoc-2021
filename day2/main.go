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
	absPath, err := filepath.Abs("day2/input.txt")
	check(err)

	fmt.Printf("Part 1 Answer: %d\n", part1(absPath))
	fmt.Printf("Part 2 Answer: %d\n", part2(absPath))
}

func part1(filePath string) int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var horizontalPos int
	var depth int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		command := line[0]
		magnitude, err := strconv.Atoi(line[1])
		check(err)

		// calculate horizontal and depth coordinates
		switch command {
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

func part2(filePath string) int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var horizontalPos int
	var depth int
	var aim int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		command := line[0]
		magnitude, err := strconv.Atoi(line[1])
		check(err)

		// calculate horizontal pos, depth, and aim
		switch command {
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


func check(err error) {
	if err != nil {
		panic(err)
	}
}
