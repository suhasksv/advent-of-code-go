package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func solveTrashCompactorPart2(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error: Could not find file '%s'\n", filename)
		return
	}

	lines := strings.Split(string(data), "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], "\n")
	}

	if len(lines) == 0 {
		return
	}

	// Find maximum width
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Pad lines
	grid := make([]string, len(lines))
	for i, line := range lines {
		grid[i] = line + strings.Repeat(" ", maxWidth-len(line))
	}

	rows := len(grid)

	// Identify vertical separators
	separatorIndices := []int{-1}
	for col := 0; col < maxWidth; col++ {
		isEmptyColumn := true
		for r := 0; r < rows; r++ {
			if grid[r][col] != ' ' {
				isEmptyColumn = false
				break
			}
		}
		if isEmptyColumn {
			separatorIndices = append(separatorIndices, col)
		}
	}
	separatorIndices = append(separatorIndices, maxWidth)

	grandTotal := 0
	fmt.Printf("Detected %d separate problems.\n", len(separatorIndices)-1)

	// Process each vertical slice
	for i := 0; i < len(separatorIndices)-1; i++ {
		startCol := separatorIndices[i] + 1
		endCol := separatorIndices[i+1]

		if startCol >= endCol {
			continue
		}

		// PART 2 LOGIC: Right-to-Left, Top-to-Bottom digits, operator at bottom
		numbers := []int{}

		operatorSegment := strings.TrimSpace(grid[rows-1][startCol:endCol])
		if operatorSegment == "" {
			continue
		}
		operator := rune(operatorSegment[0])

		// Columns: Right to Left
		for c := endCol - 1; c >= startCol; c-- {
			numStr := ""
			for r := 0; r < rows-1; r++ { // Top to row before operator
				char := rune(grid[r][c])
				if unicode.IsDigit(char) {
					numStr += string(char)
				}
			}
			if numStr != "" {
				num, _ := strconv.Atoi(numStr)
				numbers = append(numbers, num)
			}
		}

		if len(numbers) == 0 {
			continue
		}

		// Calculate result
		result := 0
		if operator == '+' {
			for _, n := range numbers {
				result += n
			}
		} else if operator == '*' {
			result = 1
			for _, n := range numbers {
				result *= n
			}
		}

		// fmt.Printf("Problem %d: %v %c = %d\n", i+1, numbers, operator, result)
		grandTotal += result
	}

	fmt.Println("------------------------------")
	fmt.Printf("Grand Total: %d\n", grandTotal)
}

func main() {
	solveTrashCompactorPart2("input.txt")
}
