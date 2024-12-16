package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func readInputFile(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func bfs(grid [][]rune, visited [][]bool, startX, startY int) ([]Point, rune) {
	rows := len(grid)
	cols := len(grid[0])
	queue := []Point{{startX, startY}}
	regionCells := []Point{}
	plantType := grid[startX][startY]

	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	visited[startX][startY] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		regionCells = append(regionCells, current)

		for _, dir := range directions {
			nx, ny := current.x+dir.x, current.y+dir.y
			if isValid(nx, ny, rows, cols) && !visited[nx][ny] && grid[nx][ny] == plantType {
				queue = append(queue, Point{nx, ny})
				visited[nx][ny] = true
			}
		}
	}

	return regionCells, plantType
}

func calculateAreaAndSides(regionCells []Point, grid [][]rune) (int, int) {
	rows := len(grid)
	cols := len(grid[0])
	area := len(regionCells)
	sides := 0

	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, cell := range regionCells {
		for _, dir := range directions {
			nx, ny := cell.x+dir.x, cell.y+dir.y
			if !isValid(nx, ny, rows, cols) || grid[nx][ny] != grid[cell.x][cell.y] {
				sides++
			}
		}
	}

	return area, sides
}

func calculateFencingPrice(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	totalPrice := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				regionCells, _ := bfs(grid, visited, i, j)
				area, sides := calculateAreaAndSides(regionCells, grid)
				totalPrice += area * sides
			}
		}
	}

	return totalPrice
}

func main() {
	grid := readInputFile("input.txt")
	totalPrice := calculateFencingPrice(grid)
	fmt.Printf("Total price of fencing: %d\n", totalPrice)
}

