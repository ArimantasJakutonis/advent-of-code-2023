package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getAllSum(lines []string) int {
	result := 0
	for _, line := range lines {
		result += concatFirstAndLastDigitsInLine(line)
	}
	return result
}

func concatFirstAndLastDigitsInLine(line string) int {
	var first rune = -1
	var last rune
	for _, char := range line {
		if unicode.IsDigit(char) {
			if first == -1 {
				first = char
			}
			last = char
		}

	}

	result, err := strconv.Atoi(string(first) + string(last))

	if err != nil {
		fmt.Println("Error parsing string to int:", err)
	}
	return result
}

func readInput() []string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return strings.Split(string(content), "\n")
}
