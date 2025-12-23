package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type RangeP2 struct {
	start int
	end   int
}

func solveCafeteriaPart2(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error: Could not find file '%s'\n", filename)
		return
	}

	content := strings.TrimSpace(string(data))
	parts := strings.Split(content, "\n\n")
	if len(parts) < 2 {
		fmt.Println("Error: Input format invalid.")
		return
	}

	rangeLines := strings.Split(strings.TrimSpace(parts[0]), "\n")
	idLines := strings.Split(strings.TrimSpace(parts[1]), "\n")

	// --- PARSING ---
	var ranges []RangeP2
	for _, line := range rangeLines {
		bounds := strings.Split(strings.TrimSpace(line), "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, RangeP2{start, end})
	}

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

	// --- PART 1: Check specific IDs ---
	freshCountP1 := 0
	for _, ingredientID := range idsToCheck {
		isFresh := false
		for _, r := range ranges {
			if ingredientID >= r.start && ingredientID <= r.end {
				isFresh = true
				break
			}
		}
		if isFresh {
			freshCountP1++
		}
	}

	fmt.Printf("Part 1 - Specific Fresh Ingredients: %d\n", freshCountP1)

	// --- PART 2: Count total unique integers in ranges ---
	// Sort ranges by start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	var merged []RangeP2
	for _, r := range ranges {
		if len(merged) == 0 {
			merged = append(merged, r)
			continue
		}

		last := &merged[len(merged)-1]
		if r.start <= last.end {
			if r.end > last.end {
				last.end = r.end
			}
		} else {
			merged = append(merged, r)
		}
	}

	totalFreshIDs := 0
	for _, r := range merged {
		totalFreshIDs += (r.end - r.start + 1)
	}

	fmt.Printf("Part 2 - Total Unique Fresh IDs: %d\n", totalFreshIDs)
}

func main() {
	solveCafeteriaPart2("input.txt")
}
