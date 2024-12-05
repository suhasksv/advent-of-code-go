package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input file
	infile := "input.txt"
	if len(os.Args) > 1 {
		infile = os.Args[1]
	}

	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the input into edges and queries
	parts := strings.Split(strings.Join(lines, "\n"), "\n\n")
	if len(parts) != 2 {
		fmt.Println("Invalid input format")
		return
	}

	edges := strings.Split(parts[0], "\n")
	queries := strings.Split(parts[1], "\n")

	// E[x] is the set of pages that must come before x
	// ER[x] is the set of pages that must come after x
	E := make(map[int]map[int]struct{})
	ER := make(map[int]map[int]struct{})

	for _, line := range edges {
		parts := strings.Split(line, "|")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		if _, exists := E[y]; !exists {
			E[y] = make(map[int]struct{})
		}
		E[y][x] = struct{}{}

		if _, exists := ER[x]; !exists {
			ER[x] = make(map[int]struct{})
		}
		ER[x][y] = struct{}{}
	}

	p1 := 0
	p2 := 0

	// Process queries
	for _, query := range queries {
		nums := strings.Split(query, ",")
		vs := make([]int, len(nums))
		for i, num := range nums {
			vs[i], _ = strconv.Atoi(num)
		}

		if len(vs)%2 != 1 {
			fmt.Println("Invalid query length")
			continue
		}

		ok := true
		for i, x := range vs {
			for j, y := range vs {
				if i < j {
					if _, exists := E[x][y]; exists {
						ok = false
					}
				}
			}
		}

		if ok {
			p1 += vs[len(vs)/2]
		} else {
			// Perform topological sorting
			good := []int{}
			D := make(map[int]int)
			queue := []int{}

			for _, v := range vs {
				D[v] = 0
				for pre := range E[v] {
					if contains(vs, pre) {
						D[v]++
					}
				}
				if D[v] == 0 {
					queue = append(queue, v)
				}
			}

			for len(queue) > 0 {
				x := queue[0]
				queue = queue[1:]
				good = append(good, x)

				for y := range ER[x] {
					if _, exists := D[y]; exists {
						D[y]--
						if D[y] == 0 {
							queue = append(queue, y)
						}
					}
				}
			}
			p2 += good[len(good)/2]
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

// Helper function to check if a slice contains a value
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
