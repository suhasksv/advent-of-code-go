package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func solveGiftShop(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error: Could not find file '%s'\n", filename)
		return
	}

	rawData := strings.TrimSpace(string(data))

	// 1. Parse the ranges
	var ranges []Range
	maxVal := 0

	parts := strings.Split(rawData, ",")
	for _, part := range parts {
		bounds := strings.Split(part, "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])

		ranges = append(ranges, Range{start, end})
		if end > maxVal {
			maxVal = end
		}
	}

	fmt.Printf("Processing %d ranges. Maximum ID to check: %d\n", len(ranges), maxVal)

	// 2. Generate invalid IDs (repeated sequence at least twice)
	candidates := make(map[int]struct{}) // set

	maxDigits := len(strconv.Itoa(maxVal))
	maxBaseLen := maxDigits / 2

	for baseLen := 1; baseLen <= maxBaseLen; baseLen++ {
		startNum := 1
		for i := 1; i < baseLen; i++ {
			startNum *= 10
		}
		endNum := startNum*10 - 1

		for i := startNum; i <= endNum; i++ {
			baseStr := strconv.Itoa(i)
			repeats := 2

			for {
				candidateStr := strings.Repeat(baseStr, repeats)

				if len(candidateStr) > maxDigits {
					break
				}

				candidateVal, _ := strconv.Atoi(candidateStr)
				if candidateVal > maxVal {
					break
				}

				candidates[candidateVal] = struct{}{}
				repeats++
			}
		}
	}

	fmt.Printf("Generated %d unique candidate 'invalid' IDs.\n", len(candidates))

	// 3. Sum invalid IDs in ranges
	var sortedCandidates []int
	for v := range candidates {
		sortedCandidates = append(sortedCandidates, v)
	}
	sort.Ints(sortedCandidates)

	totalSum := 0
	foundCount := 0

	for _, candidate := range sortedCandidates {
		for _, r := range ranges {
			if candidate >= r.start && candidate <= r.end {
				totalSum += candidate
				foundCount++
				break
			}
		}
	}

	fmt.Printf("Found %d invalid IDs in the specified ranges.\n", foundCount)
	fmt.Printf("Total Sum: %d\n", totalSum)
}

func main() {
	solveGiftShop("input.txt")
}
