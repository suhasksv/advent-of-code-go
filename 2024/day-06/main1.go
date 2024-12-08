package main

import (
	"bufio"
	"fmt"
	"os"
)

// Directions mapping: '^' (Up), '>' (Right), 'v' (Down), '<' (Left)
var directions = []struct {
	dr, dc int
}{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

// Parse the input grid from a file
func parseInput(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid
}

// Find the guard's starting position and direction
func findGuard(grid [][]rune) (int, int, int) {
	for r, row := range grid {
		for c, cell := range row {
			switch cell {
			case '^':
				return r, c, 0 // Facing Up
			case '>':
				return r, c, 1 // Facing Right
			case 'v':
				return r, c, 2 // Facing Down
			case '<':
				return r, c, 3 // Facing Left
			}
		}
	}
	panic("Guard not found in the grid")
}

// Simulate the guard's patrol
func simulatePatrol(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	r, c, dir := findGuard(grid)

	visited := make(map[string]bool)

	for r >= 0 && r < rows && c >= 0 && c < cols {
		// Mark the current position as visited
		visited[fmt.Sprintf("%d,%d", r, c)] = true

		// Calculate next position
		dr, dc := directions[dir].dr, directions[dir].dc
		nr, nc := r+dr, c+dc

		// Check for obstacles or boundaries
		if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '#' {
			// Turn Right
			dir = (dir + 1) % 4
		} else {
			// Move Forward
			r, c = nr, nc
		}
	}

	return len(visited)
}

func main() {
	inputFile := "input.txt"
	grid := parseInput(inputFile)
	result := simulatePatrol(grid)
	fmt.Printf("Distinct positions visited: %d\n", result)
}
