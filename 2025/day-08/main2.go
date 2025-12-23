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
	x, y, z int
}

type EdgeP2 struct {
	dist int
	u, v int
}

type UnionFindP2 struct {
	parent        []int
	numComponents int
}

func NewUnionFindP2(n int) *UnionFindP2 {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFindP2{parent, n}
}

func (uf *UnionFindP2) Find(i int) int {
	if uf.parent[i] != i {
		uf.parent[i] = uf.Find(uf.parent[i])
	}
	return uf.parent[i]
}

func (uf *UnionFindP2) Union(i, j int) bool {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)
	if rootI != rootJ {
		uf.parent[rootI] = rootJ
		uf.numComponents--
		return true
	}
	return false
}

func distSqP2(p1, p2 PointP2) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return dx*dx + dy*dy + dz*dz
}

func parseInput(filename string) ([]PointP2, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var points []PointP2
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, PointP2{x, y, z})
	}
	return points, nil
}

func main() {
	points, err := parseInput("input.txt")
	if err != nil {
		fmt.Println("Error: 'input.txt' not found.")
		return
	}

	n := len(points)
	var edges []EdgeP2
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, EdgeP2{distSqP2(points[i], points[j]), i, j})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	uf := NewUnionFindP2(n)
	lastU, lastV := -1, -1

	for _, e := range edges {
		if uf.Union(e.u, e.v) {
			if uf.numComponents == 1 {
				lastU = e.u
				lastV = e.v
				break
			}
		}
	}

	if lastU != -1 {
		ans := points[lastU].x * points[lastV].x
		fmt.Printf("Part 2 Answer: %d\n", ans)
	} else {
		fmt.Println("Error: Could not connect all points.")
	}
}
