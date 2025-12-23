package main

import (
	"bufio"
	"fmt"
	"os"
)

func solveLobbyBatteriesPart2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: Could not find file '%s'\n", filename)
		return
	}
	defer file.Close()

	const REQUIRED_LENGTH = 12
	totalJoltage := 0

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("Processing %d battery banks. Target Length: %d\n",
		len(lines), REQUIRED_LENGTH)

	for _, line := range lines {
		bank := line
		if bank == "" {
			continue
		}

		if len(bank) < REQUIRED_LENGTH {
			fmt.Printf("Skipping bank %s (too short)\n", bank)
			continue
		}

		// Monotonic stack
		stack := make([]byte, 0, len(bank))
		dropBudget := len(bank) - REQUIRED_LENGTH

		for i := 0; i < len(bank); i++ {
			digit := bank[i]

			for dropBudget > 0 && len(stack) > 0 && stack[len(stack)-1] < digit {
				stack = stack[:len(stack)-1]
				dropBudget--
			}

			stack = append(stack, digit)
		}

		// Truncate to required length
		stack = stack[:REQUIRED_LENGTH]

		// Convert digits to integer
		bestBankScore := 0
		for _, d := range stack {
			bestBankScore = bestBankScore*10 + int(d-'0')
		}

		totalJoltage += bestBankScore
	}

	fmt.Println("------------------------------")
	fmt.Printf("Total Output Joltage: %d\n", totalJoltage)
}

func main() {
	solveLobbyBatteriesPart2("input.txt")
}
