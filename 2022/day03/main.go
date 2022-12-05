package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func priority(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	}
	if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 27
	}
	return 0
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	sc := bufio.NewScanner(input)

	p1, p2 := 0, 0

	rsacks := make([]string, 0)

	for sc.Scan() {
		rsack := sc.Text()

		first, second := rsack[:len(rsack)/2], rsack[len(rsack)/2:]

		for _, g := range first {
			if strings.ContainsRune(second, g) {
				p1 += priority(g)
				break
			}
		}
		rsacks = append(rsacks, rsack)

		if len(rsacks) == 3 {
			for _, r := range rsacks[0] {
				if strings.ContainsRune(rsacks[1], r) && strings.ContainsRune(rsacks[2], r) {
					p2 += priority(r)
					break
				}
			}
			rsacks = make([]string, 0)
		}
	}
	fmt.Printf("Part 1: %[1]d\nPart 2: %[2]d\n", p1, p2)
}
