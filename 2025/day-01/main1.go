package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solveSafePuzzle(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: Could not find file '%s'\n", filename)
		return
	}
	defer file.Close()

	// The dial starts at 50
	currentPos := 50

	// Dial positions: 0–99
	const MODULO = 100

	// Counter for landing on 0
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
			currentPos = (currentPos + amount) % MODULO
		} else if direction == 'L' {
			// Go modulo can be negative, so normalize
			currentPos = (currentPos - amount) % MODULO
			if currentPos < 0 {
				currentPos += MODULO
			}
		}

		if currentPos == 0 {
			zeroHits++
		}
	}

	fmt.Printf("Final Password: %d\n", zeroHits)
}

func main() {
	solveSafePuzzle("input.txt")
}
