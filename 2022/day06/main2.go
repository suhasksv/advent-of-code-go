package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	sc := bufio.NewScanner(input)
	sc.Scan()

	const reqDiff = 14

	for i, _ := range sc.Text() {
		charset := make(map[byte]bool)
		for j := 0; j < reqDiff; j++ {
			charset[sc.Text()[i+j]] = true
		}
		if len(charset) == reqDiff {
			fmt.Println(i + reqDiff)
			break
		}
	}
}
