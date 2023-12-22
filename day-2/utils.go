package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func findPossibleGamesSum(games []string) int {
	sum := 0
	for i, game := range games {
		if checkSetOfGames(game, i) {
			sum += i + 1
		}
	}

	return sum
}

var cubeColorLimits = map[string]int{
	"green": 13,
	"blue":  14,
	"red":   12,
}

func checkSetOfGames(gameLine string, idx int) bool {
	individualGames := strings.Split(string(gameLine), ";")
	pattern := regexp.MustCompile(`(\d+)\s*(red|green|blue)`)

	for _, game := range individualGames {
		matches := pattern.FindAllStringSubmatch(game, -1)
		for _, match := range matches {
			amount, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println("Error parsing number:", err)
				return false
			}

			color := match[2]
			if cubeColorLimits[color] < amount {
				return false
			}
		}
	}
	return true
}

func readInput() []string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return strings.Split(string(content), "\n")
}
