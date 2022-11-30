package main

import (
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func generatePermutations(n int, stars []string, perms *[][]string) {
	if n == 1 {
		strsCopy := make([]string, len(stars))
		copy(strsCopy, stars)

		*perms = append(*perms, strsCopy)
	} else {
		for i := 0; i < n-1; i++ {
			generatePermutations(n-1, stars, perms)
			if n%2 == 0 {
				swap(stars, i, n-1)
			} else {
				swap(stars, 0, n-1)
			}
		}
		generatePermutations(n-1, stars, perms)
	}
}

func swap(stars []string, i, j int) {
	stars[i], stars[j] = stars[j], stars[i]
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	rules := strings.Split(string(input), "\n")
	happinessRules := make(map[string]map[string]int)

	regex := regexp.MustCompile("\\A(\\w+) .* (gain|lose) (\\d+) .* to (\\w+)")

	for _, rule := range rules {
		matches := regex.FindStringSubmatch(rule)[1:]

		person := matches[0]
		negative := strings.Contains(matches[1], "lose")
		happiness, _ := strconv.Atoi(matches[2])
		nextTo := matches[3]

		if negative {
			happiness *= -1
		}

		if _, present := happinessRules[person]; !present {
			happinessRules[person] = make(map[string]int)
		}
		happinessRules[person][nextTo] = happiness
	}

	var guests []string

	for guest := range happinessRules {
		guests = append(guests, guest)
	}

	maxHappiness := math.MinInt32
	var perms [][]string
	generatePermutations(len(guests), guests, &perms)

	for _, seating := range perms {
		seating = append(seating, seating[0])

		sum := 0
		for i := 0; i < len(seating)-1; i++ {
			person := seating[i]
			nextTo := seating[i+1]
			sum += happinessRules[person][nextTo]
			sum += happinessRules[nextTo][person]
		}

		if sum > maxHappiness {
			maxHappiness = sum
		}
	}

	println(maxHappiness)
}
