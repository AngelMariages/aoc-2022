package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	elfsCalories := getElfsCalories(input)

	println(getHighestCalories(elfsCalories))

	println(getSumOfHighest3Calories(elfsCalories))
}

func getElfsCalories(input string) map[int]int {
	elfsCalories := make(map[int]int)
	currentElf := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			currentElf++
			continue
		}
		elfsCalories[currentElf] += parseLine(line)
	}

	return elfsCalories
}

func getHighestCalories(elfsCalories map[int]int) int {
	highestCalories := 0

	for _, calories := range elfsCalories {
		if calories > highestCalories {
			highestCalories = calories
		}
	}

	return highestCalories
}

func getSumOfHighest3Calories(elfsCalories map[int]int) int {
	//sort elfsCalories by value
	var keys []int
	for k := range elfsCalories {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return elfsCalories[keys[i]] > elfsCalories[keys[j]]
	})

	sum := 0
	for i := 0; i < 3; i++ {
		sum += elfsCalories[keys[i]]
	}

	return sum
}

func parseLine(line string) int {
	int, _ := strconv.Atoi(line)

	return int
}
