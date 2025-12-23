package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solveTrashCompactorPart1(filename string) {
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
	separatorIndices := []int{-1} // virtual separator at -1

	for col := 0; col < maxWidth; col++ {
		isEmptyColumn := true
		for row := 0; row < rows; row++ {
			if grid[row][col] != ' ' {
				isEmptyColumn = false
				break
			}
		}
		if isEmptyColumn {
			separatorIndices = append(separatorIndices, col)
		}
	}
	separatorIndices = append(separatorIndices, maxWidth) // virtual separator at end

	grandTotal := 0

	// Process vertical slices
	for i := 0; i < len(separatorIndices)-1; i++ {
		startCol := separatorIndices[i] + 1
		endCol := separatorIndices[i+1]

		if startCol >= endCol {
			continue
		}

		var numbers []int
		operator := ""
		hasContent := false

		for r := 0; r < rows; r++ {
			segment := strings.TrimSpace(grid[r][startCol:endCol])
			if segment == "" {
				continue
			}

			hasContent = true

			if segment == "+" {
				operator = "+"
			} else if segment == "*" {
				operator = "*"
			} else {
				num, err := strconv.Atoi(segment)
				if err == nil {
					numbers = append(numbers, num)
				}
			}
		}

		if !hasContent {
			continue
		}

		// Calculate result
		result := 0
		if operator == "+" {
			for _, n := range numbers {
				result += n
			}
		} else if operator == "*" {
			result = 1
			for _, n := range numbers {
				result *= n
			}
		}

		grandTotal += result
	}

	fmt.Println("------------------------------")
	fmt.Printf("Grand Total: %d\n", grandTotal)
}

func main() {
	solveTrashCompactorPart1("input.txt")
}
