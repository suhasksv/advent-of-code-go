package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coordinate struct {
	x, y int
}

func main() {
	input, err := os.Open("./input.txt")
	defer input.Close()

	if err != nil {
		log.Fatalln(err)
	}

	sc := bufio.NewScanner(input)

	var forest [][]rune
	for sc.Scan() {
		var row []rune
		for _, tree := range sc.Text() {
			row = append(row, tree)
		}
		forest = append(forest, row)
	}

	isVisible := make(map[coordinate]bool)

	// Horizontal View

	leftMax := make([]rune, len(forest))
	rightMax := make([]rune, len(forest))

	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			if j == 0 {
				leftMax[i] = forest[i][j]
				rightMax[i] = forest[i][len(forest[0])-1]
				isVisible[coordinate{i, j}] = true
				isVisible[coordinate{i, len(forest[0]) - 1}] = true
				continue
			}
			if forest[i][j] > leftMax[i] {
				isVisible[coordinate{i, j}] = true
				leftMax[i] = forest[i][j]
			}
			if forest[i][len(forest[0])-1-j] > rightMax[i] {
				isVisible[coordinate{i, len(forest[0]) - 1 - j}] = true
				rightMax[i] = forest[i][len(forest[0])-1-j]
			}
		}
	}

	// Vertical View
	topMax := make([]rune, len(forest))
	downMax := make([]rune, len(forest))

	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			if i == 0 {
				topMax[j] = forest[i][j]
				downMax[j] = forest[len(forest)-1][j]
				isVisible[coordinate{i, j}] = true
				isVisible[coordinate{len(forest) - 1, j}] = true
				continue
			}
			if forest[i][j] > topMax[j] {
				isVisible[coordinate{i, j}] = true
				topMax[j] = forest[i][j]
			}
			if forest[len(forest)-1-i][j] > downMax[j] {
				isVisible[coordinate{len(forest) - 1 - i, j}] = true
				downMax[j] = forest[len(forest)-1-i][j]
			}
		}
	}
	fmt.Println("Part 1:", len(isVisible))
}
