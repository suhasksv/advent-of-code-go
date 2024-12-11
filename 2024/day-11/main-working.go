// ----- Practically this is working and is fast -----
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var dp = make(map[[2]int]int)

func solve(x, t int) int {

	key := [2]int{x, t}
	if val, exists := dp[key]; exists {
		return val
	}

	var ret int
	if t == 0 {
		ret = 1
	} else if x == 0 {
		ret = solve(1, t-1)
	} else if len(strconv.Itoa(x))%2 == 0 {
		dstr := strconv.Itoa(x)
		mid := len(dstr) / 2
		left, _ := strconv.Atoi(dstr[:mid])
		right, _ := strconv.Atoi(dstr[mid:])
		ret = solve(left, t-1) + solve(right, t-1)
	} else {
		ret = solve(x*2024, t-1)
	}

	dp[key] = ret
	return ret
}

func solveAll(D []int, t int) int {
	sum := 0
	for _, x := range D {
		sum += solve(x, t)
	}
	return sum
}

func main() {
	// Read input from input.txt
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse initial stones
	line := strings.TrimSpace(string(data))
	elements := strings.Split(line, " ")
	D := []int{}
	for _, elem := range elements {
		num, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}
		D = append(D, num)
	}

	// Calculate results for 25 and 75 blinks
	result25 := solveAll(D, 25)
	result75 := solveAll(D, 75)

	fmt.Printf("Result after 25 blinks: %d\n", result25)
	fmt.Printf("Result after 75 blinks: %d\n", result75)
}
