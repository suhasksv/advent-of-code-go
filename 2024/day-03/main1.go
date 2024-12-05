package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Determine the input file
	File := "input.txt"
	if len(os.Args) > 1 {
		File = os.Args[1]
	}

	// Open the file
	file, err := os.Open(File)
	if err != nil {
		fmt.Println("Error opening the file", err)
		return
	}
	defer file.Close()

	// Read the entire file into a single string
	scanner := bufio.NewScanner(file)
	var data string
	for scanner.Scan() {
		data += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	part1 := 0
	mulRegex := regexp.MustCompile(`mul\(\d+,\d+\)`)

	// Find all "mul(x, y)" patterns
	matches := mulRegex.FindAllString(data, -1)
	for _, match := range matches {
		// Extract the numbers using a regex
		nums := regexp.MustCompile(`\d+`).FindAllString(match, -1)
		if len(nums) == 2 {
			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])
			part1 += x * y
			fmt.Println(match)
		}
	}

	// Output the result
	fmt.Println(part1)
}
