package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// --- Shape Struct ---
type Shape struct {
	id         int
	coords     [][2]int
	variations [][][2]int
	area       int
}

// Normalize coords to top-left (0,0)
func normalize(coords [][2]int) [][2]int {
	minR, minC := 1<<30, 1<<30
	for _, rc := range coords {
		r, c := rc[0], rc[1]
		if r < minR {
			minR = r
		}
		if c < minC {
			minC = c
		}
	}
	norm := make([][2]int, len(coords))
	for i, rc := range coords {
		norm[i] = [2]int{rc[0] - minR, rc[1] - minC}
	}
	return norm
}

// Generate all 8 variations (rotations + flips)
func generateVariations(coords [][2]int) [][][2]int {
	var result [][][2]int
	current := coords
	for i := 0; i < 4; i++ {
		// Original
		result = append(result, normalize(current))
		// Horizontal flip
		flipped := make([][2]int, len(current))
		for j, rc := range current {
			flipped[j] = [2]int{rc[0], -rc[1]}
		}
		result = append(result, normalize(flipped))
		// Rotate 90: (r,c) -> (c,-r)
		newCoords := make([][2]int, len(current))
		for j, rc := range current {
			newCoords[j] = [2]int{rc[1], -rc[0]}
		}
		current = newCoords
	}
	return result
}

// --- Parse Input ---
func parseInput(filename string) ([]*Shape, []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	shapes := []*Shape{}
	queries := []string{}

	var currentID int
	var currentLines []string
	var readingShape bool = false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasSuffix(line, ":") && !strings.Contains(line, "x") {
			// Shape header
			if readingShape && len(currentLines) > 0 {
				coords := parseGrid(currentLines)
				shapes = append(shapes, &Shape{
					id:         currentID,
					coords:     coords,
					variations: generateVariations(coords),
					area:       len(coords),
				})
				currentLines = []string{}
			}
			idStr := strings.TrimSuffix(line, ":")
			currentID, _ = strconv.Atoi(idStr)
			readingShape = true
		} else if strings.Contains(line, "x") && strings.Contains(line, ":") {
			// Query line
			queries = append(queries, line)
		} else if readingShape {
			currentLines = append(currentLines, line)
		}
	}
	// Last shape
	if readingShape && len(currentLines) > 0 {
		coords := parseGrid(currentLines)
		shapes = append(shapes, &Shape{
			id:         currentID,
			coords:     coords,
			variations: generateVariations(coords),
			area:       len(coords),
		})
	}
	return shapes, queries
}

func parseGrid(lines []string) [][2]int {
	var coords [][2]int
	for r, row := range lines {
		for c, ch := range row {
			if ch == '#' {
				coords = append(coords, [2]int{r, c})
			}
		}
	}
	return coords
}

// --- Solve Query ---
func solveQuery(query string, shapes []*Shape) bool {
	parts := strings.Split(query, ":")
	size := strings.TrimSpace(parts[0])
	countsStr := strings.Fields(strings.TrimSpace(parts[1]))

	dims := strings.Split(size, "x")
	W, _ := strconv.Atoi(dims[0])
	H, _ := strconv.Atoi(dims[1])

	counts := make([]int, len(countsStr))
	for i, s := range countsStr {
		counts[i], _ = strconv.Atoi(s)
	}

	// Build list of pieces to place
	pieces := []*Shape{}
	totalArea := 0
	for i, cnt := range counts {
		for j := 0; j < cnt; j++ {
			pieces = append(pieces, shapes[i])
			totalArea += shapes[i].area
		}
	}
	if totalArea > W*H {
		return false
	}

	// Bitmask backtracking
	N := len(pieces)
	gridSize := W * H
	gridMask := make([]bool, gridSize) // simple bool slice instead of big.Int

	var backtrack func(k int) bool
	backtrack = func(k int) bool {
		if k == N {
			return true
		}
		p := pieces[k]
		for _, varCoords := range p.variations {
			// Try all positions
			maxR, maxC := 0, 0
			for _, rc := range varCoords {
				if rc[0] > maxR {
					maxR = rc[0]
				}
				if rc[1] > maxC {
					maxC = rc[1]
				}
			}
			for dr := 0; dr <= H-1-maxR; dr++ {
				for dc := 0; dc <= W-1-maxC; dc++ {
					collision := false
					pos := make([]int, len(varCoords))
					for i, rc := range varCoords {
						idx := (rc[0]+dr)*W + (rc[1] + dc)
						pos[i] = idx
						if gridMask[idx] {
							collision = true
							break
						}
					}
					if collision {
						continue
					}
					// Place piece
					for _, idx := range pos {
						gridMask[idx] = true
					}
					if backtrack(k + 1) {
						return true
					}
					// Remove piece
					for _, idx := range pos {
						gridMask[idx] = false
					}
				}
			}
		}
		return false
	}

	return backtrack(0)
}

// --- Main ---
func main() {
	inputFile := "input.txt"
	shapes, queries := parseInput(inputFile)
	successCount := 0
	fmt.Printf("Loaded %d shapes and %d regions.\n", len(shapes), len(queries))
	for _, q := range queries {
		if solveQuery(q, shapes) {
			successCount++
		}
	}
	fmt.Printf("Total regions that can fit all presents: %d\n", successCount)
}
