package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countPathsBetween(u, target string, graph map[string][]string, memo map[string]map[string]int) int {
	// Initialize memo map for node if needed
	if _, ok := memo[u]; !ok {
		memo[u] = make(map[string]int)
	}

	// Check memo
	if val, ok := memo[u][target]; ok {
		return val
	}

	// Base case
	if u == target {
		memo[u][target] = 1
		return 1
	}

	// Dead end
	if _, ok := graph[u]; !ok {
		memo[u][target] = 0
		return 0
	}

	total := 0
	for _, v := range graph[u] {
		total += countPathsBetween(v, target, graph, memo)
	}

	memo[u][target] = total
	return total
}

func solvePartTwo(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	graph := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		source := strings.TrimSpace(parts[0])
		dests := strings.Fields(strings.TrimSpace(parts[1]))
		graph[source] = dests
	}

	memo := make(map[string]map[string]int)

	// Scenario A: svr -> dac -> fft -> out
	pathsSvrDac := countPathsBetween("svr", "dac", graph, memo)
	pathsDacFft := countPathsBetween("dac", "fft", graph, memo)
	pathsFftOut := countPathsBetween("fft", "out", graph, memo)
	totalScenarioA := pathsSvrDac * pathsDacFft * pathsFftOut

	// Scenario B: svr -> fft -> dac -> out
	pathsSvrFft := countPathsBetween("svr", "fft", graph, memo)
	pathsFftDac := countPathsBetween("fft", "dac", graph, memo)
	pathsDacOut := countPathsBetween("dac", "out", graph, memo)
	totalScenarioB := pathsSvrFft * pathsFftDac * pathsDacOut

	totalValidPaths := totalScenarioA + totalScenarioB

	fmt.Printf("Scenario 1 (dac->fft): %d * %d * %d = %d\n", pathsSvrDac, pathsDacFft, pathsFftOut, totalScenarioA)
	fmt.Printf("Scenario 2 (fft->dac): %d * %d * %d = %d\n", pathsSvrFft, pathsFftDac, pathsDacOut, totalScenarioB)
	fmt.Printf("Total paths visiting both: %d\n", totalValidPaths)
}

func main() {
	solvePartTwo("input.txt")
}
