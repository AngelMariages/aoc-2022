package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

var (
	rock     = "AX"
	paper    = "BY"
	scissors = "CZ"
)

var scores = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

func main() {
	sumOfScores := getSumOfScores(input, 1)

	println(sumOfScores)

	sumOfScores = getSumOfScores(input, 2)

	println(sumOfScores)
}

func getSumOfScores(input string, typ int) int {
	rounds := strings.Split(input, "\n")

	sumOfScores := 0

	for _, round := range rounds {
		if round == "" {
			continue
		}

		var roundScore int

		if typ == 1 {
			roundScore = parseRound_part1(round)
		} else {
			roundScore = parseRound_part2(round)
		}

		sumOfScores += roundScore
	}

	return sumOfScores
}

func parseRound_part2(round string) int {
	chooses := strings.Split(round, " ")

	otherChoose := parseChoose(chooses[0])
	parseRoundResult := parseRoundResult(chooses[1])

	if parseRoundResult == 2 {
		return scores[getLoosingChoose(otherChoose)] + 0
	} else if parseRoundResult == 1 {
		return scores[getWinningChoose(otherChoose)] + 6
	}

	return scores[otherChoose] + 3
}

func getLoosingChoose(otherChoose string) string {
	if otherChoose == "rock" {
		return "scissors"
	} else if otherChoose == "paper" {
		return "rock"
	}

	return "paper"
}

func getWinningChoose(otherChoose string) string {
	if otherChoose == "rock" {
		return "paper"
	} else if otherChoose == "paper" {
		return "scissors"
	}

	return "rock"
}

func parseRoundResult(choose string) int {
	if choose == "Z" {
		return 1
	} else if choose == "X" {
		return 2
	}

	return 0
}

func parseRound_part1(round string) int {
	chooses := strings.Split(round, " ")

	otherChoose := parseChoose(chooses[0])
	myChoose := parseChoose(chooses[1])

	winner := getWinner(myChoose, otherChoose)

	scoreFromMatch := 3

	if winner == 1 {
		scoreFromMatch = 6
	} else if winner == 2 {
		scoreFromMatch = 0
	}

	return scores[myChoose] + scoreFromMatch
}

func getWinner(elf1, elf2 string) int {
	if elf1 == elf2 {
		return 0
	}

	if elf1 == "rock" && elf2 == "scissors" {
		return 1
	} else if elf1 == "scissors" && elf2 == "paper" {
		return 1
	} else if elf1 == "paper" && elf2 == "rock" {
		return 1
	}

	return 2
}

func parseChoose(choose string) string {
	if strings.Contains(rock, choose) {
		return "rock"
	} else if strings.Contains(paper, choose) {
		return "paper"
	} else if strings.Contains(scissors, choose) {
		return "scissors"
	}

	return ""
}
