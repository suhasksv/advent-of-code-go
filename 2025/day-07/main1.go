package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solveTachyonManifold() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error: 'input.txt' not found. Please save your puzzle input to this file.")
		return
	}
	defer file.Close()

	grid := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			grid = append(grid, line)
		}
	}

	if len(grid) == 0 {
		fmt.Println("Error: Input grid is empty.")
		return
	}

	rows := len(grid)
	cols := len(grid[0])

	// Find starting position 'S'
	activeBeams := map[[2]int]struct{}{}
foundStart:
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'S' {
				activeBeams[[2]int{r, c}] = struct{}{}
				break foundStart
			}
		}
	}

	if len(activeBeams) == 0 {
		fmt.Println("Error: Could not find start point 'S'.")
		return
	}

	totalSplits := 0

	// Simulate beams row by row
	for len(activeBeams) > 0 {
		nextBeams := map[[2]int]struct{}{}

		for beam := range activeBeams {
			r, c := beam[0], beam[1]
			nextR := r + 1

			// Beam goes out of the bottom
			if nextR >= rows {
				continue
			}

			cell := grid[nextR][c]

			if cell == '^' {
				// Splitter
				totalSplits++
				// Left beam
				if c-1 >= 0 {
					nextBeams[[2]int{nextR, c - 1}] = struct{}{}
				}
				// Right beam
				if c+1 < cols {
					nextBeams[[2]int{nextR, c + 1}] = struct{}{}
				}
			} else {
				// Continue straight down
				nextBeams[[2]int{nextR, c}] = struct{}{}
			}
		}

		activeBeams = nextBeams
	}

	fmt.Printf("Total Splits: %d\n", totalSplits)
}

func main() {
	solveTachyonManifold()
}
