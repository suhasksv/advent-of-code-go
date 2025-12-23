package main

import (
	"bufio"
	"fmt"
	"os"
)

func solvePrintingDeptPart2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: Could not find file '%s'\n", filename)
		return
	}
	defer file.Close()

	// Read non-empty lines
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}

	if len(lines) == 0 {
		fmt.Println("No data to process.")
		return
	}

	// Convert to mutable 2D grid
	rows := len(lines)
	cols := len(lines[0])
	grid := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		grid[i] = []byte(lines[i])
	}

	// 8-direction neighbor offsets
	neighborOffsets := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	fmt.Printf("Analyzing %dx%d grid for recursive removal...\n", rows, cols)

	totalRemoved := 0
	iteration := 0

	for {
		iteration++
		var rollsToRemove [][2]int

		// Step 1: Scan grid for accessible rolls
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] == '@' {
					neighborPaperCount := 0

					for _, off := range neighborOffsets {
						nr := r + off[0]
						nc := c + off[1]

						if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
							if grid[nr][nc] == '@' {
								neighborPaperCount++
							}
						}
					}

					// Rule: accessible if fewer than 4 neighbors
					if neighborPaperCount < 4 {
						rollsToRemove = append(rollsToRemove, [2]int{r, c})
					}
				}
			}
		}

		// Step 2: stop if nothing to remove
		if len(rollsToRemove) == 0 {
			break
		}

		// Step 3: remove rolls
		totalRemoved += len(rollsToRemove)

		// fmt.Printf("Iteration %d: Removing %d rolls.\n", iteration, len(rollsToRemove))

		for _, pos := range rollsToRemove {
			grid[pos[0]][pos[1]] = '.'
		}
	}

	fmt.Println("------------------------------")
	fmt.Printf("Total rolls of paper removed: %d\n", totalRemoved)
}

func main() {
	solvePrintingDeptPart2("input.txt")
}
