package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// DFS with memoization
func countPaths(node string, graph map[string][]string, memo map[string]int) int {
	if node == "out" {
		return 1
	}

	if val, ok := memo[node]; ok {
		return val
	}

	total := 0
	for _, neighbor := range graph[node] {
		total += countPaths(neighbor, graph, memo)
	}

	memo[node] = total
	return total
}

func solveReactor(filename string) {
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
		destStr := strings.TrimSpace(parts[1])
		if destStr == "" {
			graph[source] = []string{}
		} else {
			destinations := strings.Fields(destStr)
			graph[source] = destinations
		}
	}

	if _, ok := graph["you"]; !ok {
		fmt.Println("Error: Could not find starting node 'you' in input.")
		return
	}

	memo := make(map[string]int)
	result := countPaths("you", graph, memo)

	fmt.Printf("Graph loaded with %d devices.\n", len(graph))
	fmt.Printf("Total distinct paths from 'you' to 'out': %d\n", result)
}

func main() {
	solveReactor("input.txt")
}
