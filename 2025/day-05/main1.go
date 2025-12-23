package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangeP1 struct {
	start int
	end   int
}

func solveCafeteriaPart1(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error: Could not find file '%s'\n", filename)
		return
	}

	content := strings.TrimSpace(string(data))

	// Split into two sections by blank line
	// This handles both \n\n and Windows-style newlines after trimming
	parts := strings.Split(content, "\n\n")
	if len(parts) < 2 {
		fmt.Println("Error: Input format invalid. Expected ranges and IDs separated by a blank line.")
		return
	}

	rangeLines := strings.Split(strings.TrimSpace(parts[0]), "\n")
	idLines := strings.Split(strings.TrimSpace(parts[1]), "\n")

	// 1. Parse the fresh ranges
	var ranges []RangeP1
	for _, line := range rangeLines {
		bounds := strings.Split(strings.TrimSpace(line), "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, RangeP1{start: start, end: end})
	}

	// 2. Parse the IDs to check
	var idsToCheck []int
	for _, line := range idLines {
		line = strings.TrimSpace(line)
		if line != "" {
			val, _ := strconv.Atoi(line)
			idsToCheck = append(idsToCheck, val)
		}
	}

	fmt.Printf("Loaded %d ranges and %d IDs to check.\n",
		len(ranges), len(idsToCheck))

	// 3. Check freshness
	freshCount := 0

	for _, ingredientID := range idsToCheck {
		isFresh := false

		for _, r := range ranges {
			if ingredientID >= r.start && ingredientID <= r.end {
				isFresh = true
				break
			}
		}

		if isFresh {
			freshCount++
		}
	}

	fmt.Println("------------------------------")
	fmt.Printf("Total Fresh Ingredients: %d\n", freshCount)
}

func main() {
	solveCafeteriaPart1("input.txt")
}
