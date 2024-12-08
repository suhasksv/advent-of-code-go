package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	// Read the input file
	infile := "input.txt"
	if len(os.Args) >= 2 {
		infile = os.Args[1]
	}

	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	R := len(grid)
	C := len(grid[0])

	// Store antenna positions grouped by their frequencies
	P := make(map[rune][][2]int)
	for r, row := range grid {
		for c, char := range row {
			if char != '.' {
				P[char] = append(P[char], [2]int{r, c})
			}
		}
	}

	A1 := make(map[[2]int]bool)
	A2 := make(map[[2]int]bool)

	// Iterate through each position and calculate antinodes
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			for _, positions := range P {
				for i, p1 := range positions {
					for j, p2 := range positions {
						if i == j {
							continue
						}

						d1 := manhattanDistance(r, c, p1[0], p1[1])
						d2 := manhattanDistance(r, c, p2[0], p2[1])

						dr1, dc1 := r-p1[0], c-p1[1]
						dr2, dc2 := r-p2[0], c-p2[1]

						// Check if collinear
						if d1 == 2*d2 || d1*2 == d2 {
							if 0 <= r && r < R && 0 <= c && c < C && dr1*dc2 == dc1*dr2 {
								A1[[2]int{r, c}] = true
							}
						}
						if 0 <= r && r < R && 0 <= c && c < C && dr1*dc2 == dc1*dr2 {
							A2[[2]int{r, c}] = true
						}
					}
				}
			}
		}
	}

	// Output results
	fmt.Println(len(A1))
	fmt.Println(len(A2))
}

// Helper function to calculate Manhattan distance
func manhattanDistance(r1, c1, r2, c2 int) int {
	return int(math.Abs(float64(r1-r2)) + math.Abs(float64(c1-c2)))
}
