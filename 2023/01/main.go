package aoc01

import (
	"ThomasCDeveloper/advent-of-code/2023/utils"
	"fmt"
	"slices"
	"strings"
)

var path = "inputs/01.input"

func Run() {
	SolveA()
	SolveB()
}

func GetFirstDigit(s string) int {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return int(c - '0')
		}
	}
	return -1
}

func GetLastDigit(s string) int {
	split := strings.Split(s, "")
	slices.Reverse(split)
	return GetFirstDigit(strings.Join(split, ""))
}

func SolveA() {
	lines := utils.ParseInputLines(path)

	sum := 0
	for _, line := range lines {
		sum += 10*GetFirstDigit(line) + GetLastDigit(line)
	}

	fmt.Printf("Solution to part 1: %d\n", sum)
}

var replacements = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3ee",
	"four":  "f4ur",
	"five":  "f5ve",
	"six":   "s6x",
	"seven": "s7ven",
	"eight": "e8ght",
	"nine":  "n9ne",
	"zero":  "z0ro",
}

func SolveB() {
	lines := utils.ParseInputLines(path)

	sum := 0
	for _, line := range lines {
		for k, v := range replacements {
			line = strings.ReplaceAll(line, k, v)
		}
		sum += 10*GetFirstDigit(line) + GetLastDigit(line)
	}

	fmt.Printf("Solution to part 2: %d\n", sum)
}
