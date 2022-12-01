package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("./input.txt")
	defer input.Close()
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(input)

	//lines := strings.Split(string(input), "\n\n\n")

	var currcal = 0
	var maxcal = 0

	for sc.Scan() {
		food, err := strconv.Atoi(sc.Text())
		currcal += food

		if err != nil {
			if currcal > maxcal {
				maxcal = currcal
			}
			currcal = 0
		}
	}
	fmt.Println(maxcal)
}
