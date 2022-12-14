package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("./input.txt")
	defer input.Close()

	if err != nil {
		log.Fatalln(err)
	}

	cycleNum, registerEntry := 0, 1

	var ansSum int
	sc := bufio.NewScanner(input)

	for sc.Scan() {
		operations := strings.Fields(sc.Text())
		incvalue(&cycleNum, &registerEntry, &cycleNum)
		if operations[0] == "addx" {
			val, _ := strconv.Atoi(operations[1])
			incvalue(&cycleNum, &registerEntry, &ansSum)
			registerEntry += val
		}
	}
	fmt.Println("Part 1:", ansSum)
}

func incvalue(cycleNum, registerEntry, ansSum *int) {
	*cycleNum++
	if (*cycleNum-20)%40 == 0 && *cycleNum <= 220 {
		*ansSum += *registerEntry * *cycleNum
	}
}
