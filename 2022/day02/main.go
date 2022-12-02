package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	defer input.Close()

	sc := bufio.NewScanner(input)

	points := map[string]struct{ p1, p2 int }{
		"A X": {1 + 3, 3 + 0}, "A Y": {2 + 6, 1 + 3}, "A Z": {3 + 0, 2 + 6},
		"B X": {1 + 0, 1 + 0}, "B Y": {2 + 3, 2 + 3}, "B Z": {3 + 6, 3 + 6},
		"C X": {1 + 6, 2 + 0}, "C Y": {2 + 0, 3 + 3}, "C Z": {3 + 3, 1 + 6},
	}
	sc1, sc2 := 0, 0
	for sc.Scan() {
		sc1 += points[sc.Text()].p1
		sc2 += points[sc.Text()].p2
	}
	fmt.Printf("Part 1: %d\nPart 2: %d\n", sc1, sc2)
}
