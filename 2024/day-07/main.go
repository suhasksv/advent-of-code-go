package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// ParseInput parses the input text into equations.
func ParseInput(inputText string) []Equation {
	var equations []Equation
	lines := strings.Split(strings.TrimSpace(inputText), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		testValue, _ := strconv.Atoi(parts[0])
		numbers := parseNumbers(parts[1])
		equations = append(equations, Equation{TestValue: testValue, Numbers: numbers})
	}
	return equations
}

// parseNumbers converts a space-separated string of numbers into a slice of integers.
func parseNumbers(numbersText string) []int {
	numbers := strings.Fields(numbersText)
	var result []int
	for _, num := range numbers {
		parsedNum, _ := strconv.Atoi(num)
		result = append(result, parsedNum)
	}
	return result
}

// EvaluateLeftToRight evaluates the numbers and operators left-to-right.
func EvaluateLeftToRight(numbers []int, operators []string) int {
	result := numbers[0]
	for i, op := range operators {
		switch op {
		case "+":
			result += numbers[i+1]
		case "*":
			result *= numbers[i+1]
		case "||":
			result, _ = strconv.Atoi(fmt.Sprintf("%d%d", result, numbers[i+1]))
		}
	}
	return result
}

// IsValidEquation checks if a test value can be achieved with the given numbers and any combination of operators.
func IsValidEquation(testValue int, numbers []int, includeConcat bool) bool {
	numOperators := len(numbers) - 1
	operators := []string{"+", "*"}
	if includeConcat {
		operators = append(operators, "||")
	}
	combinations := generateOperatorCombinations(operators, numOperators)

	for _, combination := range combinations {
		if EvaluateLeftToRight(numbers, combination) == testValue {
			return true
		}
	}
	return false
}

// generateOperatorCombinations generates all combinations of operators of a given length.
func generateOperatorCombinations(operators []string, length int) [][]string {
	if length == 0 {
		return [][]string{{}}
	}

	smallerCombinations := generateOperatorCombinations(operators, length-1)
	var combinations [][]string
	for _, comb := range smallerCombinations {
		for _, op := range operators {
			newComb := append([]string{}, comb...)
			newComb = append(newComb, op)
			combinations = append(combinations, newComb)
		}
	}
	return combinations
}

// CalculateTotalCalibration calculates the total calibration result for valid equations.
func CalculateTotalCalibration(inputText string, includeConcat bool) int {
	equations := ParseInput(inputText)
	totalCalibration := 0

	for _, eq := range equations {
		if IsValidEquation(eq.TestValue, eq.Numbers, includeConcat) {
			totalCalibration += eq.TestValue
		}
	}

	return totalCalibration
}

// Equation represents a test value and a list of numbers.
type Equation struct {
	TestValue int
	Numbers   []int
}

func main() {
	// Read input from a file named "input.txt"
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	inputText := string(inputBytes)

	// Part 1: Without concatenation operator
	part1Result := CalculateTotalCalibration(inputText, false)
	fmt.Println("Part 1 Total Calibration:", part1Result)

	// Part 2: With concatenation operator
	part2Result := CalculateTotalCalibration(inputText, true)
	fmt.Println("Part 2 Total Calibration:", part2Result)
}

