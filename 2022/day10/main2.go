package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	input, err := os.Open("./input.txt")
	defer input.Close()

	if err != nil {
		log.Fatalln(err)
	}

	sc := bufio.NewScanner(input)

	for sc.Scan() {

	}
}
