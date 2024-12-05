package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input file
	infile := "input.txt"
	if len(os.Args) >= 2 {
		infile = os.Args[1]
	}

	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	// Parse the grid
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
	p1, p2 := 0, 0

	// Process the grid for patterns
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			// Check for "XMAS" horizontally, vertically, and diagonally
			if c+3 < C && grid[r][c] == 'X' && grid[r][c+1] == 'M' && grid[r][c+2] == 'A' && grid[r][c+3] == 'S' {
				p1++
			}
			if r+3 < R && grid[r][c] == 'X' && grid[r+1][c] == 'M' && grid[r+2][c] == 'A' && grid[r+3][c] == 'S' {
				p1++
			}
			if r+3 < R && c+3 < C && grid[r][c] == 'X' && grid[r+1][c+1] == 'M' && grid[r+2][c+2] == 'A' && grid[r+3][c+3] == 'S' {
				p1++
			}

			// Check for "SAMX" horizontally, vertically, and diagonally
			if c+3 < C && grid[r][c] == 'S' && grid[r][c+1] == 'A' && grid[r][c+2] == 'M' && grid[r][c+3] == 'X' {
				p1++
			}
			if r+3 < R && grid[r][c] == 'S' && grid[r+1][c] == 'A' && grid[r+2][c] == 'M' && grid[r+3][c] == 'X' {
				p1++
			}
			if r+3 < R && c+3 < C && grid[r][c] == 'S' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'M' && grid[r+3][c+3] == 'X' {
				p1++
			}
			if r-3 >= 0 && c+3 < C && grid[r][c] == 'S' && grid[r-1][c+1] == 'A' && grid[r-2][c+2] == 'M' && grid[r-3][c+3] == 'X' {
				p1++
			}
			if r-3 >= 0 && c+3 < C && grid[r][c] == 'X' && grid[r-1][c+1] == 'M' && grid[r-2][c+2] == 'A' && grid[r-3][c+3] == 'S' {
				p1++
			}

			// Check for "MAS" patterns surrounded by "M/S"
			if r+2 < R && c+2 < C && grid[r][c] == 'M' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'S' && grid[r+2][c] == 'M' && grid[r][c+2] == 'S' {
				p2++
			}
			if r+2 < R && c+2 < C && grid[r][c] == 'M' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'S' && grid[r+2][c] == 'S' && grid[r][c+2] == 'M' {
				p2++
			}
			if r+2 < R && c+2 < C && grid[r][c] == 'S' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'M' && grid[r+2][c] == 'M' && grid[r][c+2] == 'S' {
				p2++
			}
			if r+2 < R && c+2 < C && grid[r][c] == 'S' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'M' && grid[r+2][c] == 'S' && grid[r][c+2] == 'M' {
				p2++
			}
		}
	}

	// Output results
	fmt.Println(p1)
	fmt.Println(p2)
}
