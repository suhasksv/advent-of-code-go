package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type linuxTer struct {
	name      string
	size      int
	isFile    bool
	childDir  map[string]*linuxTer
	parentDir *linuxTer
}

func main() {
	input, err := os.Open("./input.txt")
	defer input.Close()

	if err != nil {
		log.Fatalln(err)
	}

	sc := bufio.NewScanner(input)
	var currDir *linuxTer

	var dirs []*linuxTer

	for sc.Scan() {
		lines := strings.Fields(sc.Text())

		if len(lines) > 2 {
			if lines[2] == ".." {
				currDir = currDir.parentDir
			} else if lines[2] == "/" {
				currDir = &linuxTer{"/", 0, false, make(map[string]*linuxTer), nil}
			} else {
				currDir = currDir.childDir[lines[2]]
			}
		} else if lines[0] == "dir" {
			currDir.childDir[lines[1]] = &linuxTer{lines[1], 0, false, make(map[string]*linuxTer), currDir}
			dirs = append(dirs, currDir.childDir[lines[1]])
		} else if lines[0] != "$" {
			size, _ := strconv.Atoi(lines[0])
			currDir.childDir[lines[1]] = &linuxTer{lines[1], size, true, nil, currDir}
		}
	}

	var totalSize int

	for _, dir := range dirs {
		size := calcSize(*dir)
		if size <= 100000 {
			totalSize += size
		}
	}

	fmt.Printf("Part 1: %d\n", totalSize)
}

func calcSize(root linuxTer) (size int) {
	if root.isFile {
		return root.size
	}
	for _, g := range root.childDir {
		size += calcSize(*g)
	}
	return
}
