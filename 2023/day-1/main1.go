package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func findCalibrationValue(line string) int {
	re := regexp.MustCompile(`\d`)
	matches := re.FindAllString(line, -1)

	firstDigit, _ := strconv.Atoi((matches[0]))
	lastDigit, _ := strconv.Atoi(matches[len(matches)-1])

	return firstDigit*10 + lastDigit
}

func calculateTotal(filename string) (int, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(data), "\n")
	total := 0

	for _, line := range lines {
		if line != "" {
			total += findCalibrationValue(line)
		}
	}

	return total, nil
}

func main() {
	filename := "input.txt"

	total, err := calculateTotal(filename)

	if err != nil {
		fmt.Println("Error: Failure", err)
		return
	}

	fmt.Println("Sum of all calibration values:", total)
}
