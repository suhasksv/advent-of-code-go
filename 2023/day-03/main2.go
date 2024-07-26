package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func sumGearRatios(schematic []string) (int, error) {
	totalSum := 0
	symbolRegex := regexp.MustCompile(`[^0-9.]`) // Matches any non-digit and non-dot character

	for rowIdx, row := range schematic {
		for colIdx, char := range row {
			if symbolRegex.MatchString(string(char)) { // Found a symbol
				adjacentNumbers := []int{}

				for y := max(0, rowIdx-1); y <= min(len(schematic)-1, rowIdx+1); y++ {
					for x := max(0, colIdx-1); x <= min(len(row)-1, colIdx+1); x++ {
						if y == rowIdx && x == colIdx {
							continue // Skip the symbol itself
						}
						numStr := regexp.MustCompile(`\d+`).FindString(schematic[y][x:])
						if numStr != "" {
							num, err := strconv.Atoi(numStr)
							if err != nil {
								return 0, fmt.Errorf("error converting number '%s' to integer: %v", numStr, err)
							}
							adjacentNumbers = append(adjacentNumbers, num)
						}
					}
				}

				if len(adjacentNumbers) == 2 { // Gear must have exactly two adjacent part numbers
					gearRatio := adjacentNumbers[0] * adjacentNumbers[1]
					fmt.Printf("Found gear at position (%d, %d) with ratio: %d\n", rowIdx, colIdx, gearRatio)
					totalSum += gearRatio
				}
			}
		}
	}
	return totalSum, nil
}

func main() {
	// Read the engine schematic from a file
	file, err := os.Open("input.txt") // Replace with your file path
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var schematic []string
	for scanner.Scan() {
		schematic = append(schematic, scanner.Text())
	}

	gearRatioSum, err := sumGearRatios(schematic)
	if err != nil {
		fmt.Println("Error calculating gear ratio sum:", err)
		return
	}
	fmt.Println("The sum of all gear ratios in the engine schematic is:", gearRatioSum)
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
