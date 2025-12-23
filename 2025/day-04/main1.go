package main

import (
	"bufio"
	"fmt"
	"os"
)

func solvePrintingDeptPart1(filename string) {
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

	// Convert to 2D grid
	rows := len(lines)
	cols := len(lines[0])
	grid := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		grid[i] = []byte(lines[i])
	}

	accessibleCount := 0

	// 8-directional neighbor offsets
	neighborOffsets := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	fmt.Printf("Analyzing %dx%d grid...\n", rows, cols)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// Only care about paper rolls ('@')
			if grid[r][c] == '@' {
				neighborPaperCount := 0

				// Check all 8 neighbors
				for _, offset := range neighborOffsets {
					nr := r + offset[0]
					nc := c + offset[1]

					if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
						if grid[nr][nc] == '@' {
							neighborPaperCount++
						}
					}
				}

				// Rule: Accessible if fewer than 4 neighbors
				if neighborPaperCount < 4 {
					accessibleCount++
					// Optional visual marking:
					// grid[r][c] = 'x'
				}
			}
		}
	}

	fmt.Printf("Accessible paper rolls: %d\n", accessibleCount)
}

func main() {
	solvePrintingDeptPart1("input.txt")
}
