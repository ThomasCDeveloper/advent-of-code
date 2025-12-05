package main

import (
	"ThomasCDeveloper/advent-of-code/2023/utils"
	"fmt"
	"os"
	"slices"
	"strings"
)

func ValueOfRune(r rune) int {
	for i, v := range []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'} {
		if v == r {
			return i
		}
	}
	return -1
}

func part1(data []string) int {
	total := 0

	for _, line := range data {
		subtotal := 0
		for _, r := range line {
			v := ValueOfRune(r)
			if v != -1 {
				subtotal = v
				break
			}
		}
		reversed := strings.Split(line, "")
		slices.Reverse(reversed)
		for _, r := range strings.Join(reversed, "") {
			v := ValueOfRune(r)
			if v != -1 {
				subtotal = subtotal*10 + v
				break
			}
		}

		total += subtotal
	}
	return total
}

var replacements = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3ree",
	"four":  "f4ur",
	"five":  "f5ve",
	"six":   "s6x",
	"seven": "s7ven",
	"eight": "e8ght",
	"nine":  "n9ne",
}

func part2(data []string) int {
	for i := range data {
		line := data[i]
		for key, val := range replacements {
			line = strings.ReplaceAll(line, key, val)
		}
		data[i] = line
	}
	return part1(data)
}

func main() {
	path := utils.GetPath(os.Args)
	data := utils.ParseInputLines(path)

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
