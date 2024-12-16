package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func calculateFencingPrice(mapInput [][]rune) int {
	rows := len(mapInput)
	cols := len(mapInput[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	isValid := func(x, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}

	var bfs func(x, y int) []Point
	bfs = func(x, y int) []Point {
		queue := []Point{{x, y}}
		regionCells := []Point{}
		visited[x][y] = true
		plantType := mapInput[x][y]

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			regionCells = append(regionCells, current)

			directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
			for _, d := range directions {
				nx, ny := current.x+d.x, current.y+d.y
				if isValid(nx, ny) && !visited[nx][ny] && mapInput[nx][ny] == plantType {
					visited[nx][ny] = true
					queue = append(queue, Point{nx, ny})
				}
			}
		}
		return regionCells
	}

	calculateAreaAndPerimeter := func(regionCells []Point) (int, int) {
		area := len(regionCells)
		perimeter := 0

		for _, cell := range regionCells {
			directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
			for _, d := range directions {
				nx, ny := cell.x+d.x, cell.y+d.y
				if !isValid(nx, ny) || mapInput[nx][ny] != mapInput[cell.x][cell.y] {
					perimeter++
				}
			}
		}
		return area, perimeter
	}

	totalPrice := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				regionCells := bfs(i, j)
				area, perimeter := calculateAreaAndPerimeter(regionCells)
				totalPrice += area * perimeter
			}
		}
	}
	return totalPrice
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mapInput := [][]rune{}
	for scanner.Scan() {
		mapInput = append(mapInput, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := calculateFencingPrice(mapInput)
	fmt.Printf("Total price of fencing: %d\n", result)
}


