package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Coord struct {
	x, y int
}

func assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func assertf(condition bool, message string, args ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(message, args...))
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type CacheEntry struct {
	code  string
	robot int
}

func handleLine(line string, robots int, cache map[CacheEntry]int) int {
	numericalValue, err := strconv.Atoi(line[:len(line)-1])
	check(err)
	// fmt.Printf("Numerical value: %d\n", numericalValue)

	code := expandCodeNumericalKeypad(line)

	length := calcLengthAtLevel(code, cache, 0, robots)

	fmt.Printf("%s: %d\n", line, length)

	return length * numericalValue
}

func calcLengthAtLevel(code string, cache map[CacheEntry]int, robotNumBefore, maxRobots int) int {
	codeExtended := expandCodeDirectionKeypad(code)

	if robotNumBefore == maxRobots {
		return len(code)
	}

	length := 0
	for _, curr := range splitOnA(codeExtended) {
		cacheEntry := CacheEntry{curr, robotNumBefore + 1}
		if _, ok := cache[cacheEntry]; ok {
			length += cache[cacheEntry]
			continue
		}
		count := calcLengthAtLevel(curr, cache, robotNumBefore+1, maxRobots)
		length += count
	}

	cache[CacheEntry{code, robotNumBefore}] = length
	return length
}

func splitOnA(code string) []string {
	splits := make([]string, 0)
	start := 0
	for i, c := range code {
		if c == 'A' {
			splits = append(splits, code[start:i+1])
			start = i + 1
		}
	}
	return splits
}

func expandCodeNumericalKeypad(code string) string {
	// Keypad matrix:
	// 7 8 9
	// 4 5 6
	// 1 2 3
	// x 0 A
	// starting at A

	expanded := ""
	curr := 'A'
	for i := 0; i < len(code); i++ {
		from := curr
		to := rune(code[i])
		expanded += numericalFromTo(from, to)
		expanded += "A"
		curr = to
	}

	return expanded
}

func numericalFromTo(from, to rune) string {
	coords := map[rune]Coord{
		'7': {0, 0},
		'8': {1, 0},
		'9': {2, 0},
		'4': {0, 1},
		'5': {1, 1},
		'6': {2, 1},
		'1': {0, 2},
		'2': {1, 2},
		'3': {2, 2},
		'0': {1, 3},
		'A': {2, 3},
	}

	fromCoord := coords[from]
	toCoord := coords[to]

	xDiff := toCoord.x - fromCoord.x
	yDiff := toCoord.y - fromCoord.y

	vertical := ""
	for yDiff < 0 {
		vertical += "^"
		yDiff++
	}
	for yDiff > 0 {
		vertical += "v"
		yDiff--
	}

	horizontal := ""
	for xDiff < 0 {
		horizontal += "<"
		xDiff++
	}
	for xDiff > 0 {
		horizontal += ">"
		xDiff--
	}

	xDiff = toCoord.x - fromCoord.x

	// Priority: < over ^ over v over >
	if fromCoord.y == 3 && toCoord.x == 0 {
		return vertical + horizontal
	} else if fromCoord.x == 0 && toCoord.y == 3 {
		return horizontal + vertical
	} else if xDiff < 0 {
		return horizontal + vertical
	} else {
		return vertical + horizontal
	}
}

func expandCodeDirectionKeypad(code string) string {
	// Keypad matrix:
	// x ^ A
	// < v >

	expanded := ""
	curr := 'A'
	for i := 0; i < len(code); i++ {
		from := curr
		to := rune(code[i])
		expanded += directionFromTo(from, to)
		expanded += "A"
		curr = to
	}

	return expanded
}

func directionFromTo(from, to rune) string {
	coords := map[rune]Coord{
		'^': {1, 0},
		'A': {2, 0},
		'<': {0, 1},
		'v': {1, 1},
		'>': {2, 1},
	}

	fromCoord := coords[from]
	toCoord := coords[to]

	xDiff := toCoord.x - fromCoord.x
	yDiff := toCoord.y - fromCoord.y

	vertical := ""
	for yDiff < 0 {
		vertical += "^"
		yDiff++
	}
	for yDiff > 0 {
		vertical += "v"
		yDiff--
	}

	horizontal := ""
	for xDiff < 0 {
		horizontal += "<"
		xDiff++
	}
	for xDiff > 0 {
		horizontal += ">"
		xDiff--
	}

	xDiff = toCoord.x - fromCoord.x

	// Priority: < over ^ over v over >
	if fromCoord.x == 0 && toCoord.y == 0 {
		return horizontal + vertical
	} else if fromCoord.y == 0 && toCoord.x == 0 {
		return vertical + horizontal
	} else if xDiff < 0 {
		return horizontal + vertical
	} else {
		return vertical + horizontal
	}
}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	const robots = 25
	cache := make(map[CacheEntry]int)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		sum += handleLine(line, robots, cache)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
	assert(sum != 211498201171116, "Answer is not correct")
	assert(sum != 1449630236257308, "Answer is not correct")
	assert(sum != 553709479156924, "Answer is not correct")
	// assert(sum == 213536 || sum == 258369757013802, "Answer is not correct")
}