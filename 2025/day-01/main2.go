package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func floorDiv(a, b int) int {
	if a >= 0 {
		return a / b
	}
	return -(((-a) + b - 1) / b)
}

func solveSafePuzzle(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: Could not find file '%s'\n", filename)
		return
	}
	defer file.Close()

	currentPos := 50
	const MODULO = 100
	zeroHits := 0

	fmt.Printf("Starting Position: %d\n", currentPos)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := strings.TrimSpace(scanner.Text())
		if command == "" {
			continue
		}

		direction := command[0]
		amount, _ := strconv.Atoi(command[1:])

		if direction == 'R' {
			zeroHits += floorDiv(currentPos+amount, MODULO) -
				floorDiv(currentPos, MODULO)

			currentPos = (currentPos + amount) % MODULO

		} else if direction == 'L' {
			zeroHits += floorDiv(currentPos-1, MODULO) -
				floorDiv(currentPos-amount-1, MODULO)

			currentPos = (currentPos - amount) % MODULO
			if currentPos < 0 {
				currentPos += MODULO
			}
		}
	}

	fmt.Printf("Final Password (total clicks on 0): %d\n", zeroHits)
}

func main() {
	solveSafePuzzle("input.txt")
}
