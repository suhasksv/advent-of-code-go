package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	row, col int
}

func parseMap(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, char := range line {
			row[i], _ = strconv.Atoi(string(char))
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return grid, nil
}

func findTrailheads(grid [][]int) []Point {
	var trailheads []Point
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 0 {
				trailheads = append(trailheads, Point{r, c})
			}
		}
	}
	return trailheads
}

func countReachableNines(grid [][]int, start Point) int {
	rows, cols := len(grid), len(grid[0])
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make(map[Point]bool)
	queue := []Point{start}
	reachableNines := make(map[Point]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}
		visited[current] = true

		// Check if this position is a reachable '9'
		if grid[current.row][current.col] == 9 {
			reachableNines[current] = true
			continue
		}

		// Explore neighbors with height increasing by 1
		for _, dir := range directions {
			nr, nc := current.row+dir.row, current.col+dir.col
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == grid[current.row][current.col]+1 {
				queue = append(queue, Point{nr, nc})
			}
		}
	}
	return len(reachableNines)
}

func calculateTrailheadScores(grid [][]int) int {
	trailheads := findTrailheads(grid)
	totalScore := 0
	for _, trailhead := range trailheads {
		totalScore += countReachableNines(grid, trailhead)
	}
	return totalScore
}

func main() {
	// Read input from "input.txt"
	filePath := "input.txt"
	grid, err := parseMap(filePath)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Calculate and print the total score
	totalScore := calculateTrailheadScores(grid)
	fmt.Println("Total Score:", totalScore)
}
