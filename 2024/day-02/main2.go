package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Check if a report is safe
func isSafe2(report []int) bool {
	differences := make([]int, len(report)-1)

	// Calculate differences between adjacent levels
	for i := 0; i < len(report)-1; i++ {
		differences[i] = report[i+1] - report[i]
	}

	// Check if all differences are within [-3, -1] or [1, 3]
	for _, diff := range differences {
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}
	}

	// Check if the differences are all positive or all negative
	allPositive := true
	allNegative := true
	for _, diff := range differences {
		if diff > 0 {
			allNegative = false
		} else if diff < 0 {
			allPositive = false
		}
	}

	return allPositive || allNegative
}

// Check if a report can be made safe by removing one level
func canBeMadeSafe2(report []int) bool {
	for i := 0; i < len(report); i++ {
		// Create a copy of the report with the i-th level removed
		modifiedReport := append([]int{}, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)

		// Check if the modified report is safe
		if isSafe2(modifiedReport) {
			return true
		}
	}
	return false
}

// Parse a single line into a slice of integers
func parseLine2(line string) ([]int, error) {
	parts := strings.Fields(line)
	report := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		report[i] = num
	}
	return report, nil
}

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var safeCount int
	scanner := bufio.NewScanner(file)

	// Process each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Parse the line into a report
		report, err := parseLine2(line)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			return
		}

		// Check if the report is safe or can be made safe
		if isSafe2(report) || canBeMadeSafe2(report) {
			safeCount++
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the number of safe reports
	fmt.Println("Number of safe reports:", safeCount)
}
