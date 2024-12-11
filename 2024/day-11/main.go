// ----- Theoretically should work but taking a lot of time for Part 2 -----

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func splitEvenDigits(num int) (int, int) {
	numStr := strconv.Itoa(num)
	mid := len(numStr) / 2
	left, _ := strconv.Atoi(numStr[:mid])
	right, _ := strconv.Atoi(numStr[mid:])
	return left, right
}

func simulateBlinks(stones []int, blinks int) []int {
	for i := 0; i < blinks; i++ {
		newStones := []int{}
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if len(strconv.Itoa(stone))%2 == 0 {
				left, right := splitEvenDigits(stone)
				newStones = append(newStones, left, right)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}
	return stones
}

func main() {
	// Read input from input.txt
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse initial stones
	line := strings.TrimSpace(string(data))
	elements := strings.Split(line, " ")
	initialStones := []int{}
	for _, elem := range elements {
		stone, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}
		initialStones = append(initialStones, stone)
	}

	// Number of blinks to simulate
	numBlinks1 := 25
	numBlinks2 := 75

	// Simulate and count the number of stones after 75 blinks
	part1 := simulateBlinks(initialStones, numBlinks1)
	fmt.Printf("Number of stones after %d blinks: %d\n", numBlinks1, len(part1))

	part2 := simulateBlinks(initialStones, numBlinks2)
	fmt.Printf("Number of stones after %d blinks: %d\n", numBlinks2, len(part2))
}
