package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type linuxTerm struct {
	name      string
	size      int
	isFile    bool
	childDir  map[string]*linuxTerm
	parentDir *linuxTerm
}

func main() {
	//Read input file
	input, err := os.Open("./input.txt")
	defer input.Close()

	if err != nil {
		log.Fatalln(err)
	}

	sc := bufio.NewScanner(input)
	var currDir *linuxTerm

	var dirs []*linuxTerm

	for sc.Scan() {
		line := strings.Fields(sc.Text())
		if len(line) > 2 {
			if line[2] == ".." {
				currDir = currDir.parentDir
			} else if line[2] == "/" {
				currDir = &linuxTerm{"/", 0, false, make(map[string]*linuxTerm), nil}
				dirs = append(dirs, currDir)
			} else {
				currDir = currDir.childDir[line[2]]
			}
		} else if line[0] == "dir" {
			currDir.childDir[line[1]] = &linuxTerm{line[1], 0, false, make(map[string]*linuxTerm), currDir}
			dirs = append(dirs, currDir.childDir[line[1]])
		} else if line[0] != "$" {
			size, _ := strconv.Atoi(line[0])
			currDir.childDir[line[1]] = &linuxTerm{line[1], size, true, nil, currDir}
		}
	}

	toFree := 30000000 - (70000000 - calSize(*dirs[0]))
	var smallestEnoughSize = calSize(*dirs[0])

	for _, dir := range dirs {
		size := calSize(*dir)
		if size > toFree && size-toFree < smallestEnoughSize-toFree {
			smallestEnoughSize = size
		}
	}

	fmt.Println("Part 2: ", smallestEnoughSize)
}

func calSize(root linuxTerm) (size int) {
	if root.isFile {
		return root.size
	}
	for _, d := range root.childDir {
		size += calSize(*d)
	}
	return
}
