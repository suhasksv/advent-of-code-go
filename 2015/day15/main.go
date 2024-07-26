package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parseInput(input string) []Ingredient {
	var ingredients []Ingredient
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		name := parts[0]
		capacity, _ := strconv.Atoi(strings.TrimRight(parts[2], ","))
		durability, _ := strconv.Atoi(strings.TrimRight(parts[4], ","))
		flavor, _ := strconv.Atoi(strings.TrimRight(parts[6], ","))
		texture, _ := strconv.Atoi(strings.TrimRight(parts[8], ","))
		calories, _ := strconv.Atoi(parts[10])
		ingredients = append(ingredients, Ingredient{name, capacity, durability, flavor, texture, calories})
	}
	return ingredients
}

func scoreRecipe(recipe []int, ingredients []Ingredient) int {
	capacity := 0
	durability := 0
	flavor := 0
	texture := 0
	calories := 0
	for i, amount := range recipe {
		capacity += amount * ingredients[i].capacity
		durability += amount * ingredients[i].durability
		flavor += amount * ingredients[i].flavor
		texture += amount * ingredients[i].texture
		calories += amount * ingredients[i].calories
	}
	if capacity < 0 || durability < 0 || flavor < 0 || texture < 0 || calories != 500 {
		return 0
	}
	return capacity * durability * flavor * texture
}

func findBestRecipe(ingredients []Ingredient, remaining int, recipe []int, level int) int {
	if level == len(ingredients)-1 {
		recipe[level] = remaining
		score := scoreRecipe(recipe, ingredients)
		recipe[level] = 0
		return score
	}
	maxScore := 0
	for i := 0; i <= remaining; i++ {
		recipe[level] = i
		score := findBestRecipe(ingredients, remaining-i, recipe, level+1)
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	input := ""
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}
	ingredients := parseInput(input)
	recipe := make([]int, len(ingredients))
	maxScore := findBestRecipe(ingredients, 100, recipe, 0)
	fmt.Println("The best score is", maxScore)
}
