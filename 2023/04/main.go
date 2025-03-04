package aoc04

import (
	"ThomasCDeveloper/advent-of-code/2023/utils"
	"fmt"
	"math"
	"slices"
	"strings"
)

var path = "inputs/04.input"

func Run() {
	SolveA()
	SolveB()
}

func SolveA() {
	data := utils.ParseInputLines(path)
	sum := 0
	for _, line := range data {
		correct := 0
		curated := strings.Split(strings.Split(strings.ReplaceAll(line, "  ", " "), ": ")[1], " | ")
		cardNumbers := strings.Split(curated[0], " ")
		pickedNumbers := strings.Split(curated[1], " ")

		for _, n := range pickedNumbers {
			if slices.Index(cardNumbers, n) != -1 {
				correct += 1
			}
		}
		if correct != 0 {
			sum += int(math.Pow(2, float64(correct-1)))
		}
	}

	fmt.Printf("Solution to part 1: %d\n", sum)
}

func SolveB() {
	data := utils.ParseInputLines(path)
	var nbScratch = make([]int, len(data))
	for n := range nbScratch {
		nbScratch[n] = 1
	}
	for n, line := range data {
		correct := 0
		curated := strings.Split(strings.Split(strings.ReplaceAll(line, "  ", " "), ": ")[1], " | ")
		cardNumbers := strings.Split(curated[0], " ")
		pickedNumbers := strings.Split(curated[1], " ")

		for _, n := range pickedNumbers {
			if slices.Index(cardNumbers, n) != -1 {
				correct += 1
			}
		}
		if correct != 0 {
			for nexts := range correct {
				nbScratch[n+nexts+1] += nbScratch[n]
			}
		}
	}

	sum := 0
	for _, n := range nbScratch {
		sum += n
	}

	fmt.Printf("Solution to part 2: %d\n", sum)
}
