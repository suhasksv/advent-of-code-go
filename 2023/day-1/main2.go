package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func findCalibrationValuePart2(line string) int {
	// Create a slice to store the indices and values of all digits found
	type digit struct {
		value int
		index int
	}
	digits := []digit{}

	for i := 0; i < len(line); i++ {
		// Check for numeric digits
		if charDigit := line[i] - '0'; charDigit >= 0 && charDigit <= 9 {
			digits = append(digits, digit{int(charDigit), i})
		} else {
			// Check for spelled-out digits
			for word, value := range digitMap {
				if strings.HasPrefix(line[i:], word) {
					digits = append(digits, digit{value, i})
				}
			}
		}
	}

	// Sort the digits by their index to get the first and last
	sort.Slice(digits, func(i, j int) bool {
		return digits[i].index < digits[j].index
	})

	// If any digits were found, return the combination of the first and last
	if len(digits) > 0 {
		return digits[0].value*10 + digits[len(digits)-1].value
	}

	return 0 // If no digits are found
}

func calculateTotalPart2(filename string) (int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(data), "\n")
	total := 0
	for _, line := range lines {
		if line != "" {
			total += findCalibrationValuePart2(line)
		}
	}
	return total, nil
}

func main() {
	filename := "input.txt" // Replace with your actual input file

	total, err := calculateTotalPart2(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Sum of all calibration values:", total)
}
