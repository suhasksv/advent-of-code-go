package main

import (
	"encoding/json"
	"io/ioutil"
)

func findNumber(input interface{}) []int {
	var numbers []int

	switch input := input.(type) {
	case []interface{}:
		for _, val := range input {
			numbers = append(numbers, findNumber(val)...)
		}
	case map[string]interface{}:
		noRed := true

		for _, val := range input {
			if str, ok := val.(string); ok && str == "red" {
				noRed = false
				break
			}
		}

		if noRed {
			for _, val := range input {
				numbers = append(numbers, findNumber(val)...)
			}
		}

	case float64:
		numbers = append(numbers, int(input))
	}

	return numbers
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{}, 0)
	err1 := json.Unmarshal(input, &data)
	if err1 != nil {
		panic(err1)
	}

	sum := 0

	for _, num := range findNumber(data) {
		sum += num
	}

	println(sum)
}
