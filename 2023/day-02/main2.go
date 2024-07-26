package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeCountPart2 struct {
	Red   int
	Green int
	Blue  int
}

func ParseGameDataPart2(filePath string) (map[int][]CubeCountPart2, error) {
	games := make(map[int][]CubeCountPart2)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, ": ")
		gameIDStr := strings.Split(parts[0], " ")[1]
		gameID, err := strconv.Atoi(gameIDStr)
		if err != nil {
			return nil, err
		}

		subsetsStr := strings.Split(parts[1], "; ")
		var gameData []CubeCountPart2

		for _, subsetStr := range subsetsStr {
			cubesStr := strings.Split(subsetStr, ", ")
			cubeCount := CubeCountPart2{}
			for _, cubeStr := range cubesStr {
				countColor := strings.Split(cubeStr, " ")
				count, err := strconv.Atoi(countColor[0])
				if err != nil {
					return nil, err
				}
				color := countColor[1]

				switch color {
				case "red":
					cubeCount.Red += count
				case "green":
					cubeCount.Green += count
				case "blue":
					cubeCount.Blue += count
				}
			}
			gameData = append(gameData, cubeCount)
		}

		games[gameID] = gameData
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func CalculateMinimumCubesPart2(gameData []CubeCountPart2) (int, int, int) {
	minRed, minGreen, minBlue := 0, 0, 0

	for _, subset := range gameData {
		if subset.Red > minRed {
			minRed = subset.Red
		}
		if subset.Green > minGreen {
			minGreen = subset.Green
		}
		if subset.Blue > minBlue {
			minBlue = subset.Blue
		}
	}

	return minRed, minGreen, minBlue
}

func CalculatePowerPart2(red, green, blue int) int {
	return red * green * blue
}

func FindTotalPowerPart2(filePath string) (int, error) {
	games, err := ParseGameDataPart2(filePath)
	if err != nil {
		return 0, err
	}

	totalPower := 0

	for _, gameData := range games {
		minRed, minGreen, minBlue := CalculateMinimumCubesPart2(gameData)
		power := CalculatePowerPart2(minRed, minGreen, minBlue)
		totalPower += power
	}

	return totalPower, nil
}

func main() {
	filePath := "input.txt"
	totalPower, err := FindTotalPowerPart2(filePath)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Total power: %d\n", totalPower)
}
