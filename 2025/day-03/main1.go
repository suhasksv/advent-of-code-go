package main

import (
	"bufio"
	"fmt"
	"os"
)

func solveLobbyBatteriesPart1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: Could not find file '%s'\n", filename)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalJoltage := 0
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("Processing %d battery banks...\n", len(lines))

	for _, line := range lines {
		bank := line
		if bank == "" {
			continue
		}

		bestBankScore := 0
		found := false

		// Try tens digit from 9 down to 0
		for tensDigit := 9; tensDigit >= 0; tensDigit-- {
			target := byte('0' + tensDigit)
			firstIdx := -1

			// Find first occurrence of tens digit
			for i := 0; i < len(bank); i++ {
				if bank[i] == target {
					firstIdx = i
					break
				}
			}

			if firstIdx != -1 && firstIdx+1 < len(bank) {
				// Find max digit in remaining suffix
				maxUnits := byte('0')
				for j := firstIdx + 1; j < len(bank); j++ {
					if bank[j] > maxUnits {
						maxUnits = bank[j]
					}
				}

				unitsDigit := int(maxUnits - '0')
				bestBankScore = tensDigit*10 + unitsDigit
				found = true
				break
			}
		}

		if found {
			totalJoltage += bestBankScore
		}
	}

	fmt.Println("------------------------------")
	fmt.Printf("Total Output Joltage: %d\n", totalJoltage)
}

func main() {
	solveLobbyBatteriesPart1("input.txt")
}
