package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	solveQuantumManifold("input.txt")
}

func solveQuantumManifold(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: '%s' not found.\n", filename)
		return
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			grid = append(grid, line)
		}
	}

	if len(grid) == 0 {
		fmt.Println("Empty grid.")
		return
	}

	rows := len(grid)
	cols := len(grid[0])

	// 1. Find the starting position 'S'
	startR, startC := -1, -1
	for r := 0; r < rows; r++ {
		if idx := strings.IndexRune(grid[r], 'S'); idx != -1 {
			startR = r
			startC = idx
			break
		}
	}

	if startR == -1 {
		fmt.Println("Error: Could not find start point 'S'.")
		return
	}

	// 2. Track timelines
	currentTimelines := make(map[int]int)
	currentTimelines[startC] = 1

	for r := startR; r < rows-1; r++ {
		nextTimelines := make(map[int]int)

		for c, count := range currentTimelines {
			nextCell := rune(grid[r+1][c])

			if nextCell == '^' {
				// Split timeline
				if c-1 >= 0 {
					nextTimelines[c-1] += count
				}
				if c+1 < cols {
					nextTimelines[c+1] += count
				}
			} else {
				// Continue straight
				nextTimelines[c] += count
			}
		}

		currentTimelines = nextTimelines
	}

	// 3. Sum total timelines
	totalTimelines := 0
	for _, count := range currentTimelines {
		totalTimelines += count
	}

	fmt.Printf("Total Active Timelines: %d\n", totalTimelines)
}
