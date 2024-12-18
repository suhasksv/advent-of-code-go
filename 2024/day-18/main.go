package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func parseInput(fileName string) ([]Point, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var points []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		if len(coords) != 2 {
			return nil, fmt.Errorf("invalid input format: %s", line)
		}
		x, err1 := strconv.Atoi(coords[0])
		y, err2 := strconv.Atoi(coords[1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid coordinates: %s", line)
		}
		points = append(points, Point{X: x, Y: y})
	}
	return points, nil
}

func buildGrid(size int) [][]bool {
	grid := make([][]bool, size)
	for i := range grid {
		grid[i] = make([]bool, size)
	}
	return grid
}

func bfs(grid [][]bool, start, end Point) bool {
	size := len(grid)
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := []Point{start}
	visited := make(map[Point]bool)
	visited[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == end {
			return true
		}

		for _, dir := range directions {
			next := Point{X: curr.X + dir.X, Y: curr.Y + dir.Y}
			if next.X >= 0 && next.X < size && next.Y >= 0 && next.Y < size && !grid[next.Y][next.X] && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return false
}

func findShortestPath(grid [][]bool, start, end Point) int {
	size := len(grid)
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := []struct {
		Point Point
		Steps int
	}{{Point: start, Steps: 0}}
	visited := make(map[Point]bool)
	visited[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Point == end {
			return curr.Steps
		}

		for _, dir := range directions {
			next := Point{X: curr.Point.X + dir.X, Y: curr.Point.Y + dir.Y}
			if next.X >= 0 && next.X < size && next.Y >= 0 && next.Y < size && !grid[next.Y][next.X] && !visited[next] {
				visited[next] = true
				queue = append(queue, struct {
					Point Point
					Steps int
				}{Point: next, Steps: curr.Steps + 1})
			}
		}
	}

	return -1
}

func main() {
	inputFile := "input.txt"
	points, err := parseInput(inputFile)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	size := 71
	grid := buildGrid(size)
	start := Point{X: 0, Y: 0}
	end := Point{X: size - 1, Y: size - 1}

	// Part 1: Find the shortest path after 1024 bytes
	for i := 0; i < 1024 && i < len(points); i++ {
		point := points[i]
		grid[point.Y][point.X] = true
	}
	shortestPath := findShortestPath(grid, start, end)
	if shortestPath == -1 {
		fmt.Println("No valid path to the exit after 1024 bytes.")
	} else {
		fmt.Printf("Shortest Path after 1024 bytes: %d\n", shortestPath)
	}

	// Reset the grid for Part 2
	grid = buildGrid(size)

	// Part 2: Find the first byte that blocks the path
	for _, point := range points {
		grid[point.Y][point.X] = true
		if !bfs(grid, start, end) {
			fmt.Printf("First blocking byte: %d,%d\n", point.X, point.Y)
			break
		}
	}
}
