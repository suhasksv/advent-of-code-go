package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// sumPartNumbers processes the schematic and calculates the sum of all part numbers
func sumPartNumbers(schematic []string) int {
	totalSum := 0
	re := regexp.MustCompile(`\d+`)

	for rowIdx, row := range schematic {
		matches := re.FindAllStringIndex(row, -1)
		for _, match := range matches {
			numberStart, numberEnd := match[0], match[1]
			numberStr := row[numberStart:numberEnd]
			isPartNumber := false

			// Check surrounding characters
			for y := max(0, rowIdx-1); y <= min(len(schematic)-1, rowIdx+1); y++ {
				for x := max(0, numberStart-1); x <= min(len(schematic[y]), numberEnd+1); x++ {
					if x < len(schematic[y]) {
						char := schematic[y][x]
						if char < '0' || char > '9' && char != '.' {
							isPartNumber = true
							break
						}
					}
				}
				if isPartNumber {
					break
				}
			}

			if isPartNumber {
				number, err := strconv.Atoi(numberStr)
				if err != nil {
					fmt.Printf("Error converting number: %v\n", err) // Debug print
					continue
				}
				fmt.Printf("Identified part number: %d\n", number) // Debug print
				totalSum += number
			}
		}
	}

	return totalSum
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

func main() {
	filePath := "input.txt" // Replace with the actual file path

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the schematic from the file
	var schematic []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		schematic = append(schematic, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Calculate the sum of part numbers
	partNumberSum := sumPartNumbers(schematic)

	fmt.Println("The sum of all part numbers in the engine schematic is:", partNumberSum)
}
