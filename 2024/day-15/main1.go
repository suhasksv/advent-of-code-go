package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

var DIRS = map[rune]Point{
	'^': {-1, 0}, // up
	'v': {1, 0},  // down
	'<': {0, -1}, // left
	'>': {0, 1},  // right
}

func parseInput(filename string) ([][]rune, []rune, Point) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]rune{}
	var moves []rune
	var robotPos Point

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "@") {
			for y, char := range line {
				if char == '@' {
					robotPos = Point{x: len(grid), y: y}
				}
			}
		}
		if strings.Contains(line, "#") || strings.Contains(line, "O") {
			grid = append(grid, []rune(line))
		} else {
			moves = append(moves, []rune(line)...)
		}
	}
	return grid, moves, robotPos
}

func isValid(grid [][]rune, p Point) bool {
	return p.x >= 0 && p.x < len(grid) && p.y >= 0 && p.y < len(grid[0]) && grid[p.x][p.y] != '#'
}

func simulate(grid [][]rune, moves []rune, robot Point) int {
	for _, move := range moves {
		dir := DIRS[move]
		next := Point{x: robot.x + dir.x, y: robot.y + dir.y}
		if isValid(grid, next) {
			if grid[next.x][next.y] == 'O' {
				boxNext := Point{x: next.x + dir.x, y: next.y + dir.y}
				if isValid(grid, boxNext) && grid[boxNext.x][boxNext.y] == '.' {
					grid[boxNext.x][boxNext.y] = 'O'
					grid[next.x][next.y] = '@'
					grid[robot.x][robot.y] = '.'
					robot = next
				}
			} else {
				grid[next.x][next.y] = '@'
				grid[robot.x][robot.y] = '.'
				robot = next
			}
		}
	}

	gpsSum := 0
	for x, row := range grid {
		for y, cell := range row {
			if cell == 'O' {
				gpsSum += 100*x + y
			}
		}
	}
	return gpsSum
}

func main() {
	infile := "input.txt"
	if len(os.Args) > 1 {
		infile = os.Args[1]
	}

	grid, moves, robot := parseInput(infile)
	gpsSum := simulate(grid, moves, robot)
	fmt.Println("Sum of GPS coordinates:", gpsSum)
}

