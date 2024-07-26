package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeCountPart1 struct {
	Red   int
	Green int
	Blue  int
}

func ParseGameDataPart1(filePath string) (map[int][]CubeCountPart1, error) {
	games := make(map[int][]CubeCountPart1)

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
		var gameData []CubeCountPart1

		for _, subsetStr := range subsetsStr {
			cubesStr := strings.Split(subsetStr, ", ")
			cubeCount := CubeCountPart1{}
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

func IsGamePossiblePart1(gameData []CubeCountPart1, redLimit, greenLimit, blueLimit int) bool {
	for _, subset := range gameData {
		if subset.Red > redLimit || subset.Green > greenLimit || subset.Blue > blueLimit {
			return false
		}
	}
	return true
}

func FindPossibleGames(filePath string, redLimit, greenLimit, blueLimit int) (int, error) {
	games, err := ParseGameDataPart1(filePath)
	if err != nil {
		return 0, err
	}

	totalIDSum := 0

	for gameID, gameData := range games {
		if IsGamePossiblePart1(gameData, redLimit, greenLimit, blueLimit) {
			totalIDSum += gameID
		}
	}

	return totalIDSum, nil
}

func main() {
	filePath := "input.txt"
	redLimit, greenLimit, blueLimit := 12, 13, 14

	totalIDSum, err := FindPossibleGames(filePath, redLimit, greenLimit, blueLimit)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Sum of IDs of possible games: %d\n", totalIDSum)
}
