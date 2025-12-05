package aoc02

import (
	"ThomasCDeveloper/advent-of-code/2023/utils"
	"fmt"
	"strconv"
	"strings"
)

var path = "inputs/02.input"

func Run() {
	data := utils.ParseInputLines(path)

	maxs := []map[string]int{}
	for _, d := range data {
		reveals := strings.Split(strings.ReplaceAll(strings.Split(d, ": ")[1], ",", ";"), "; ")
		max := map[string]int{}
		for _, r := range reveals {
			c := strings.Split(r, " ")[1]
			v, _ := strconv.Atoi(strings.Split(r, " ")[0])

			if _, ok := max[c]; !ok {
				max[c] = v
			} else {
				if v > max[c] {
					max[c] = v
				}
			}
		}
		maxs = append(maxs, max)
	}

	SolveA(maxs)
	SolveB(maxs)
}

func SolveA(maxs []map[string]int) {
	sum := 0
	for i, m := range maxs {
		if !(m["red"] > 12 || m["green"] > 13 || m["blue"] > 14) {
			sum += i + 1
		}
	}

	fmt.Printf("Solution to part 1: %d\n", sum)
}

func SolveB(maxs []map[string]int) {
	sum := 0
	for _, max := range maxs {
		sum += max["red"] * max["green"] * max["blue"]
	}

	fmt.Printf("Solution to part 2: %d\n", sum)
}
