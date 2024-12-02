package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//func main() {
//	var total1, total2 int
//	var left, right []int
//
//	// Open the input file
//	file, err := os.Open("input.txt")
//	if err != nil {
//		fmt.Println("Error opening file:", err)
//		return
//	}
//	defer file.Close()
//
//	// Read input line by line
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		line := scanner.Text()
//		parts := strings.Fields(line)
//		if len(parts) != 2 {
//			fmt.Println("Invalid input format")
//			return
//		}
//		l, err1 := strconv.Atoi(parts[0])
//		r, err2 := strconv.Atoi(parts[1])
//		if err1 != nil || err2 != nil {
//			fmt.Println("Error parsing numbers:", err1, err2)
//			return
//		}
//		left = append(left, l)
//		right = append(right, r)
//	}
//
//	if err := scanner.Err(); err != nil {
//		fmt.Println("Error reading file:", err)
//		return
//	}
//
//	// Sort both slices
//	sort.Ints(left)
//	sort.Ints(right)
//
//	// Part 1: Calculate the total distance
//	for i := range left {
//		total1 += abs(left[i] - right[i])
//	}
//
//	// Part 2: Calculate the sum based on frequency matching
//	rightCount := make(map[int]int)
//	for _, num := range right {
//		rightCount[num]++
//	}
//
//	for _, l := range left {
//		if count, found := rightCount[l]; found {
//			total2 += l * count
//		}
//	}
//
//	// Print results for both parts
//	fmt.Println("Part 1:", total1)
//	fmt.Println("Part 2:", total2)
//}
//
//// Helper function for absolute value
//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
