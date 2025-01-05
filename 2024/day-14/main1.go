package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
//	"strings"
)

type Robot struct {
	px, py, vx, vy int
}

var DIRS = [4][2]int{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func ints(s string) []int {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(s, -1)
	res := make([]int, len(matches))
	for i, match := range matches {
		res[i], _ = strconv.Atoi(match)
	}
	return res
}

func main() {
	infile := "input.txt"
	if len(os.Args) >= 2 {
		infile = os.Args[1]
	}

	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	robots := []Robot{}

	for scanner.Scan() {
		line := scanner.Text()
		values := ints(line)
		robots = append(robots, Robot{values[0], values[1], values[2], values[3]})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
		return
	}

	const X, Y = 101, 103
	// Uncomment for smaller grid
	// const X, Y = 11, 7

	q1, q2, q3, q4 := 0, 0, 0, 0

	for t := 1; t < int(math.Pow10(6)); t++ {
		G := make([][]rune, Y)
		for i := range G {
			G[i] = make([]rune, X)
			for j := range G[i] {
				G[i][j] = '.'
			}
		}

		if t == 100 {
			q1, q2, q3, q4 = 0, 0, 0, 0
			mx, my := X/2, Y/2
			for i := range robots {
				robots[i].px = (robots[i].px + robots[i].vx) % X
				robots[i].py = (robots[i].py + robots[i].vy) % Y
				if robots[i].px < 0 {
					robots[i].px += X
				}
				if robots[i].py < 0 {
					robots[i].py += Y
				}
				G[robots[i].py][robots[i].px] = '#'

				if robots[i].px < mx && robots[i].py < my {
					q1++
				} else if robots[i].px > mx && robots[i].py < my {
					q2++
				} else if robots[i].px < mx && robots[i].py > my {
					q3++
				} else if robots[i].px > mx && robots[i].py > my {
					q4++
				}
			}
			fmt.Println(q1 * q2 * q3 * q4)
		}

		components := 0
		seen := map[[2]int]bool{}

		for y := 0; y < Y; y++ {
			for x := 0; x < X; x++ {
				if G[y][x] == '#' && !seen[[2]int{x, y}] {
					components++
					queue := [][2]int{{x, y}}
					for len(queue) > 0 {
						sx, sy := queue[0][0], queue[0][1]
						queue = queue[1:]
						if seen[[2]int{sx, sy}] {
							continue
						}
						seen[[2]int{sx, sy}] = true
						for _, d := range DIRS {
							nx, ny := sx+d[0], sy+d[1]
							if nx >= 0 && nx < X && ny >= 0 && ny < Y && G[ny][nx] == '#' {
								queue = append(queue, [2]int{nx, ny})
							}
						}
					}
				}
			}
		}

		if components <= 200 {
			fmt.Println(t)
			for _, row := range G {
				fmt.Println(string(row))
			}
			break
		}
	}
}

