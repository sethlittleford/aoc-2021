package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
)

// ReadInts reads the file at the relative filePath into memory as a slice of ints
func ReadInts(filePath string) []int {
	absPath, err := filepath.Abs(filePath)
	CheckErr(err, "failed to return absolute file path representation")

	file, err := os.Open(absPath)
	CheckErr(err, "could not open file at: ", absPath)
	defer file.Close()

	result := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		CheckErr(err, "could not convert string to int while reading file")
		result = append(result, i)
	}

	return result
}

// ReadStrings reads the file at the relative filePath into memory as a slice of strings
func ReadStrings(filePath string) []string {
	absPath, err := filepath.Abs(filePath)
	CheckErr(err, "failed to return absolute file path representation")

	file, err := os.Open(absPath)
	CheckErr(err, "could not open file at: ", absPath)
	defer file.Close()

	result := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}
