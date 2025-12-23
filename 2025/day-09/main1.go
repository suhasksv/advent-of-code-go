package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func parsePoints(input string) []Point {
	var points []Point
	re := regexp.MustCompile(`-?\d+`)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		nums := re.FindAllString(line, -1)
		if len(nums) >= 2 {
			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])
			points = append(points, Point{x, y})
		}
	}
	return points
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solveMovieTheater(input string) int {
	points := parsePoints(input)
	n := len(points)
	if n < 2 {
		fmt.Printf("Warning: Only found %d valid points. Check input format.\n", n)
		return 0
	}

	fmt.Printf("Successfully parsed %d red tiles.\n", n)

	maxArea := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			width := abs(points[i].x-points[j].x) + 1
			height := abs(points[i].y-points[j].y) + 1
			area := width * height
			maxArea = max(maxArea, area)
		}
	}
	return maxArea
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	exampleInput := `
7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`
	fmt.Println("--- Example ---")
	exResult := solveMovieTheater(exampleInput)
	fmt.Printf("Example Result: %d (Expected: 50)\n", exResult)

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("input.txt not found. Run with local file for real answer.")
		return
	}

	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf("Real Input Result: %d\n", solveMovieTheater(string(data)))
}
