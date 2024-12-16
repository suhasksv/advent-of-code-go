package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	A     [2]int // Button A increments (X, Y)
	B     [2]int // Button B increments (X, Y)
	Prize [2]int // Prize position (X, Y)
}

func findMinCostToWinPrizes(clawMachines []Machine) (int, int) {
	const maxPresses = 100
	totalCost := 0
	prizesWon := 0

	for _, machine := range clawMachines {
		xA, yA := machine.A[0], machine.A[1]
		xB, yB := machine.B[0], machine.B[1]
		xPrize, yPrize := machine.Prize[0], machine.Prize[1]

		minCost := int(^uint(0) >> 1) // Set to maximum int value
		foundSolution := false

		for nA := 0; nA <= maxPresses; nA++ {
			for nB := 0; nB <= maxPresses; nB++ {
				if nA*xA+nB*xB == xPrize && nA*yA+nB*yB == yPrize {
					foundSolution = true
					cost := 3*nA + 1*nB
					if cost < minCost {
						minCost = cost
					}
				}
			}
		}

		if foundSolution {
			prizesWon++
			totalCost += minCost
		}
	}

	return prizesWon, totalCost
}

func readInputFile(filePath string) ([]Machine, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var clawMachines []Machine
	var machineData []string
	
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		machineData = append(machineData, line)

		if len(machineData) == 3 { // Expecting exactly 3 lines per machine
			machine, err := parseMachineData(machineData, lineNumber)
			if err != nil {
				fmt.Printf("Error parsing machine data at lines %d-%d: %v\n", lineNumber-2, lineNumber, err)
			} else {
				clawMachines = append(clawMachines, machine)
			}
			machineData = nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(machineData) > 0 {
		fmt.Printf("Incomplete machine data at the end of file: %v\n", machineData)
	}

	return clawMachines, nil
}

func parseMachineData(machineData []string, lineNumber int) (Machine, error) {
	var machine Machine
	var err error

	parseValues := func(input string) ([2]int, error) {
		parts := strings.Split(input, ", ")
		if len(parts) != 2 {
			return [2]int{}, fmt.Errorf("invalid format: %s", input)
		}
		
		x, err1 := strconv.Atoi(parts[0][2:]) // Skip "X+"
		y, err2 := strconv.Atoi(parts[1][2:]) // Skip "Y+"
		if err1 != nil || err2 != nil {
			return [2]int{}, fmt.Errorf("failed to parse values: %v, %v", err1, err2)
		}
		return [2]int{x, y}, nil
	}

	aValues, err := parseValues(strings.Split(machineData[0], ": ")[1])
	if err != nil {
		return machine, err
	}

	bValues, err := parseValues(strings.Split(machineData[1], ": ")[1])
	if err != nil {
		return machine, err
	}

	prizeValues, err := parseValues(strings.Split(machineData[2], ": ")[1])
	if err != nil {
		return machine, err
	}

	machine.A = aValues
	machine.B = bValues
	machine.Prize = prizeValues
	return machine, nil
}

func main() {
	filePath := "input.txt"
	clawMachines, err := readInputFile(filePath)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	prizesWon, totalCost := findMinCostToWinPrizes(clawMachines)
	fmt.Printf("Prizes won: %d\n", prizesWon)
	fmt.Printf("Total cost: %d\n", totalCost)
}

