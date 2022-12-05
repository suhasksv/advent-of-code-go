package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.Open("./input.txt")
	defer input.Close()
	if err != nil {
		log.Fatalln(err)
	}

	sc := bufio.NewScanner(input)

	var subsets, overlap int

	for sc.Scan() {
		var startst, endst, startnd, endnd int
		fmt.Sscanf(sc.Text(), "%d-%d,%d-%d", &startst, &endst, &startnd, &endnd)
		if startnd >= startst && endnd <= endst || startst >= startnd && endst <= endnd {
			subsets++
		}
		if startnd <= endst && endnd >= startst || startst <= endnd && endst >= startnd {
			overlap++
		}
	}

	fmt.Printf("Part 1: %[1]d\nPart 2: %[2]d\n", subsets, overlap)
}
