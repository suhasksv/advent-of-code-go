package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseLine(line string) (int, []int, error) {
	// Extract the target pattern inside brackets
	targetRe := regexp.MustCompile(`\[([.#]+)\]`)
	targetMatch := targetRe.FindStringSubmatch(line)
	if targetMatch == nil {
		return 0, nil, fmt.Errorf("no target pattern found")
	}

	targetStr := targetMatch[1]
	targetMask := 0
	for i, ch := range targetStr {
		if ch == '#' {
			targetMask |= (1 << i)
		}
	}

	// Extract buttons inside parentheses
	buttonRe := regexp.MustCompile(`\(([\d,]+)\)`)
	buttonMatches := buttonRe.FindAllStringSubmatch(line, -1)
	buttonMasks := []int{}

	for _, bm := range buttonMatches {
		indices := strings.Split(bm[1], ",")
		mask := 0
		for _, idxStr := range indices {
			idx, err := strconv.Atoi(idxStr)
			if err != nil {
				return 0, nil, fmt.Errorf("invalid button index")
			}
			mask |= (1 << idx)
		}
		buttonMasks = append(buttonMasks, mask)
	}

	return targetMask, buttonMasks, nil
}

func bfs(target int, buttons []int) int {
	type State struct {
		mask    int
		presses int
	}

	queue := []State{{0, 0}}
	visited := map[int]bool{0: true}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.mask == target {
			return current.presses
		}

		for _, b := range buttons {
			newMask := current.mask ^ b
			if !visited[newMask] {
				visited[newMask] = true
				queue = append(queue, State{newMask, current.presses + 1})
			}
		}
	}

	return -1 // impossible
}

func solve(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	totalPresses := 0
	machinesProcessed := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		targetMask, buttonMasks, err := parseLine(line)
		if err != nil {
			fmt.Println("Warning:", err, "Line:", line)
			continue
		}

		minPresses := bfs(targetMask, buttonMasks)
		if minPresses != -1 {
			totalPresses += minPresses
			machinesProcessed++
		} else {
			fmt.Println("Warning: Could not solve machine:", line)
		}
	}

	fmt.Printf("Processed %d machines.\n", machinesProcessed)
	fmt.Printf("Total fewest presses required: %d\n", totalPresses)
}

func main() {
	solve("input.txt")
}
