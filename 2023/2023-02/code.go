package main

// CMD: go run *.go

import (
	"fmt"
	"strconv"
	"strings"
)

func ProcessGame(line string) int {
	picks := strings.Split(line, ": ")[1]
	gameNumber, _ := strconv.Atoi(strings.Split(strings.Split(line, ": ")[0], "e ")[1])

	for _, pick := range strings.Split(picks, "; ") {
		colors := []int{0, 0, 0}

		for _, pic := range strings.Split(pick, ", ") {
			number, _ := strconv.Atoi(strings.Split(pic, " ")[0])
			color := strings.Split(pic, " ")[1]

			switch color {
			case "red":
				colors[0] += number
				if colors[0] > 12 {
					return 0
				}
			case "green":
				colors[1] += number
				if colors[1] > 13 {
					return 0
				}
			case "blue":
				colors[2] += number
				if colors[2] > 14 {
					return 0
				}
			}
		}
	}
	return gameNumber
}

func ProcessGame2(line string) int {
	picks := strings.Split(line, ": ")[1]

	max := []int{0, 0, 0}
	for _, pick := range strings.Split(picks, "; ") {
		colors := []int{0, 0, 0}

		for _, pic := range strings.Split(pick, ", ") {
			number, _ := strconv.Atoi(strings.Split(pic, " ")[0])
			color := strings.Split(pic, " ")[1]

			switch color {
			case "red":
				colors[0] += number
				if colors[0] > max[0] {
					max[0] = colors[0]
				}
			case "green":
				colors[1] += number
				if colors[1] > max[1] {
					max[1] = colors[1]
				}
			case "blue":
				colors[2] += number
				if colors[2] > max[2] {
					max[2] = colors[2]
				}
			}
		}
	}
	return max[0] * max[1] * max[2]
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:")

	compt := 0
	for _, line := range data {
		compt += ProcessGame(line)
	}
	fmt.Println(compt)

	// PART 2
	fmt.Println("Part 2:")

	compt = 0
	for _, line := range data {
		compt += ProcessGame2(line)
	}
	fmt.Println(compt)
}
