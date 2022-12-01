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
	var maxcal0, maxcal1, maxcal2 = 0, 0, 0

	for sc.Scan() {
		food, err := strconv.Atoi(sc.Text())
		currcal += food

		if err != nil {
			if currcal > maxcal2 {
				maxcal2 = currcal
			}
			if maxcal2 > maxcal1 {
				maxcal2, maxcal1 = maxcal1, maxcal2
			}
			if maxcal2 > maxcal0 {
				maxcal1, maxcal0 = maxcal0, maxcal1
			}
			currcal = 0
		}
	}
	fmt.Println(maxcal0 + maxcal1 + maxcal2)
}
