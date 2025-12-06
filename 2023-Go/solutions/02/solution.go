package main

import (
	"ThomasCDeveloper/advent-of-code/2023/utils"
	"fmt"
	"os"
	"strings"
)

func part1(data []string) int {
	total := 0

	for _, game := range data {
		possible := true
		id := utils.Atoi(strings.Split(strings.Split(game, ": ")[0], " ")[1])

		sets := strings.Split(strings.Split(game, ": ")[1], "; ")
		for _, set := range sets {
			for _, color := range strings.Split(set, ", ") {
				num := strings.Split(color, " ")[0]
				word := strings.Split(color, " ")[1]

				if word == "red" && utils.Atoi(num) > 12 {
					possible = false
				}
				if word == "green" && utils.Atoi(num) > 13 {
					possible = false
				}
				if word == "blue" && utils.Atoi(num) > 14 {
					possible = false
				}
			}
		}
		if possible {
			total += id
		}

	}

	return total
}

func part2(data []string) int {
	total := 0
	for _, game := range data {
		minimums := map[string]int{}

		sets := strings.Split(strings.Split(game, ": ")[1], "; ")
		for _, set := range sets {
			for _, color := range strings.Split(set, ", ") {
				num := utils.Atoi(strings.Split(color, " ")[0])
				word := strings.Split(color, " ")[1]

				if val, ok := minimums[word]; ok {
					minimums[word] = max(val, num)
				} else {
					minimums[word] = num
				}
			}
		}
		total += minimums["green"] * minimums["blue"] * minimums["red"]
	}
	return total
}

func main() {
	path := utils.GetPath(os.Args)
	data := utils.ParseInputLines(path)

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
