package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int
}

type Edge struct {
	dist int
	u, v int
}

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent, size}
}

func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] != i {
		uf.parent[i] = uf.Find(uf.parent[i])
	}
	return uf.parent[i]
}

func (uf *UnionFind) Union(i, j int) bool {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)
	if rootI != rootJ {
		uf.parent[rootI] = rootJ
		uf.size[rootJ] += uf.size[rootI]
		return true
	}
	return false
}

func distSq(p1, p2 Point) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return dx*dx + dy*dy + dz*dz
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error: 'input.txt' not found.")
		return
	}
	defer file.Close()

	var points []Point
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
		points = append(points, Point{x, y, z})
	}

	n := len(points)
	var edges []Edge
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, Edge{distSq(points[i], points[j]), i, j})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	uf := NewUnionFind(n)
	limit := 1000
	if len(edges) < limit {
		limit = len(edges)
	}
	for k := 0; k < limit; k++ {
		uf.Union(edges[k].u, edges[k].v)
	}

	componentSizes := make([]int, 0)
	seen := make(map[int]bool)
	for i := 0; i < n; i++ {
		root := uf.Find(i)
		if !seen[root] {
			componentSizes = append(componentSizes, uf.size[root])
			seen[root] = true
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(componentSizes)))

	result := 1
	for i := 0; i < 3 && i < len(componentSizes); i++ {
		result *= componentSizes[i]
	}

	fmt.Printf("Part 1 : %d\n", result)
}
