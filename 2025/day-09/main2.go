package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type PointP2 struct {
	x, y int
}

func minP2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxP2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func readInput(filename string) ([]PointP2, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	points := []PointP2{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, PointP2{x, y})
	}
	return points, nil
}

func compressCoordinates(points []PointP2) (map[int]int, map[int]int, int, int) {
	xSet := map[int]struct{}{}
	ySet := map[int]struct{}{}
	for _, p := range points {
		xSet[p.x] = struct{}{}
		ySet[p.y] = struct{}{}
	}

	uniqueX := []int{}
	uniqueY := []int{}
	for k := range xSet {
		uniqueX = append(uniqueX, k)
	}
	for k := range ySet {
		uniqueY = append(uniqueY, k)
	}
	sort.Ints(uniqueX)
	sort.Ints(uniqueY)

	mapX := map[int]int{}
	mapY := map[int]int{}
	idx := 1
	for i, val := range uniqueX {
		mapX[val] = idx
		if i < len(uniqueX)-1 && uniqueX[i+1] > val+1 {
			idx += 2
		} else {
			idx++
		}
	}
	W := idx + 2

	idx = 1
	for i, val := range uniqueY {
		mapY[val] = idx
		if i < len(uniqueY)-1 && uniqueY[i+1] > val+1 {
			idx += 2
		} else {
			idx++
		}
	}
	H := idx + 2

	return mapX, mapY, W, H
}

func floodFill(grid [][]int, W, H int) {
	type Coord struct{ x, y int }
	queue := []Coord{{0, 0}}
	grid[0][0] = 2

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		dirs := []Coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, d := range dirs {
			nx, ny := curr.x+d.x, curr.y+d.y
			if nx >= 0 && nx < W && ny >= 0 && ny < H && grid[ny][nx] == 0 {
				grid[ny][nx] = 2
				queue = append(queue, Coord{nx, ny})
			}
		}
	}
}

func buildPrefixSum(grid [][]int, W, H int) [][]int {
	prefix := make([][]int, H)
	for y := 0; y < H; y++ {
		prefix[y] = make([]int, W)
		for x := 0; x < W; x++ {
			val := 0
			if grid[y][x] == 2 {
				val = 1
			}
			top := 0
			left := 0
			diag := 0
			if y > 0 {
				top = prefix[y-1][x]
			}
			if x > 0 {
				left = prefix[y][x-1]
			}
			if x > 0 && y > 0 {
				diag = prefix[y-1][x-1]
			}
			prefix[y][x] = val + top + left - diag
		}
	}
	return prefix
}

func countBad(prefix [][]int, x1, y1, x2, y2 int) int {
	val := func(x, y int) int {
		if x < 0 || y < 0 {
			return 0
		}
		return prefix[y][x]
	}
	return val(x2, y2) - val(x2, y1-1) - val(x1-1, y2) + val(x1-1, y1-1)
}

func solve(filename string) {
	points, err := readInput(filename)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	if len(points) == 0 {
		fmt.Println("No points found.")
		return
	}

	mapX, mapY, W, H := compressCoordinates(points)

	// Create grid
	grid := make([][]int, H)
	for i := 0; i < H; i++ {
		grid[i] = make([]int, W)
	}

	// Draw polygon
	n := len(points)
	for i := 0; i < n; i++ {
		p1 := points[i]
		p2 := points[(i+1)%n]
		cx1, cy1 := mapX[p1.x], mapY[p1.y]
		cx2, cy2 := mapX[p2.x], mapY[p2.y]

		if cx1 == cx2 {
			for y := minP2(cy1, cy2); y <= maxP2(cy1, cy2); y++ {
				grid[y][cx1] = 1
			}
		} else {
			for x := minP2(cx1, cx2); x <= maxP2(cx1, cx2); x++ {
				grid[cy1][x] = 1
			}
		}
	}

	// Flood fill outside
	floodFill(grid, W, H)

	// Build prefix sum for outside ("bad") cells
	prefix := buildPrefixSum(grid, W, H)

	// Check every pair of points
	maxArea := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1 := points[i]
			p2 := points[j]
			width := absP2(p1.x-p2.x) + 1
			height := absP2(p1.y-p2.y) + 1
			area := width * height
			if area <= maxArea {
				continue
			}

			cx1, cx2 := minP2(mapX[p1.x], mapX[p2.x]), maxP2(mapX[p1.x], mapX[p2.x])
			cy1, cy2 := minP2(mapY[p1.y], mapY[p2.y]), maxP2(mapY[p1.y], mapY[p2.y])
			if countBad(prefix, cx1, cy1, cx2, cy2) == 0 {
				maxArea = area
			}
		}
	}

	fmt.Printf("Largest valid area: %d\n", maxArea)
}

func absP2(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	solve("input.txt")
}
