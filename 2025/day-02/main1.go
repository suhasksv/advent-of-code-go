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

	// 2. Generate candidate "invalid" IDs (A + A pattern)
	var candidates []int

	maxDigits := len(strconv.Itoa(maxVal))
	maxHalfLen := maxDigits / 2

	for halfLen := 1; halfLen <= maxHalfLen; halfLen++ {
		startNum := 1
		for i := 1; i < halfLen; i++ {
			startNum *= 10
		}
		endNum := startNum*10 - 1

		for i := startNum; i <= endNum; i++ {
			s := strconv.Itoa(i)
			candidateStr := s + s
			candidateVal, _ := strconv.Atoi(candidateStr)

			if candidateVal > maxVal {
				break
			}

			candidates = append(candidates, candidateVal)
		}
	}

	fmt.Printf("Generated %d candidate 'invalid' IDs.\n", len(candidates))

	// 3. Sum invalid IDs that fall into any range
	sort.Ints(candidates)

	totalSum := 0
	foundCount := 0

	for _, candidate := range candidates {
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
