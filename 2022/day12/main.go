package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type point struct {
	x, y int
}

func main() {
	input, err := os.Open("./input.txt")

	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)

	heightmap := make([][]rune, 0)
	var start, end point

	for sc.Scan() {
		var line []rune
		for i, elevation := range sc.Text() {
			if elevation == 'S' {
				start = point{i, len(heightmap)}
				elevation = 'a'
			}
			if elevation == 'E' {
				end = point{i, len(heightmap)}
				elevation = 'z'
			}
			line = append(line, elevation)
		}
		heightmap = append(heightmap, line)
	}

	visited := make(map[point]bool)
	toVisit := []point{start}
	distFromStart := map[point]int{start: 0}

	for {
		currPoint := toVisit[0]
		visited[currPoint] = true

		if currPoint == end {
			fmt.Println(distFromStart[end])
			break
		}

		for _, near := range [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} {
			j, i := near[1], near[0]
			nextPoint := point{currPoint.x + j, currPoint.y + i}

			if !visited[nextPoint] && nextPoint.x >= 0 && nextPoint.y >= 0 && nextPoint.x < len(heightmap[0]) && nextPoint.y < len(heightmap) && (heightmap[nextPoint.y][nextPoint.x]-heightmap[currPoint.y][currPoint.x] <= 1) {
				if distFromStart[nextPoint] == 0 {
					toVisit = append(toVisit, nextPoint)
					distFromStart[nextPoint] = distFromStart[currPoint] + 1
				}
				if distFromStart[nextPoint] >= distFromStart[currPoint]+1 {
					distFromStart[nextPoint] = distFromStart[currPoint] + 1
				}
			}
		}
		sort.Slice(toVisit, func(i, j int) bool {
			return distFromStart[toVisit[i]] < distFromStart[toVisit[j]]
		})
	}
}
