package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Check for command line argument to specify the input file
	file := "input.txt"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	// Read input file
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Initialize variables
	D := string(data)
	part2 := 0
	enabled := true

	// Iterate through the string
	for i := 0; i < len(D); i++ {
		// Check for "do()"
		if i+4 <= len(D) && D[i:i+4] == "do()" {
			enabled = true
		}
		// Check for "don't()"
		if i+7 <= len(D) && D[i:i+7] == "don't()" {
			enabled = false
		}

		// Look for occurrences of "mul("
		if i+4 <= len(D) && D[i:i+4] == "mul(" {
			// Find the closing parenthesis
			j := i + 4
			for j < len(D) && D[j] != ')' {
				j++
			}
			// Extract numbers inside "mul(x, y)"
			if j < len(D) && D[j] == ')' {
				// Use regex to find integers in the range
				re := regexp.MustCompile(`\d+`)
				matches := re.FindAllString(D[i:j+1], -1)

				if len(matches) == 2 {
					x, err1 := strconv.Atoi(matches[0])
					y, err2 := strconv.Atoi(matches[1])

					if err1 == nil && err2 == nil {
						// Make sure the closing parenthesis isn't followed by a number
						if j+1 >= len(D) || !strings.ContainsAny(string(D[j+1]), "0123456789") {
							// Perform the multiplication if enabled
							if enabled {
								part2 += x * y
							}
							fmt.Println(D[i : j+1]) // Print the found "mul(x, y)"
						}
					}
				}
			}
		}
	}

	// Output the final result
	fmt.Println(part2)
}
