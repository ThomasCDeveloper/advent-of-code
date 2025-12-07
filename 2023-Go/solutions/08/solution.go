package main

import (
	"ThomasCDeveloper/advent-of-code/2023/utils"
	"fmt"
	"os"
	"strings"
)

var directions = map[string][]string{}
var startingPoints = []string{}
var instructions = ""

func numberOfSteps(current string) int {
	i := 0
	for current[len(current)-1] != 'Z' {
		if instructions[i%len(instructions)] == 'R' {
			current = directions[current][1]
		} else {
			current = directions[current][0]
		}
		i += 1
	}
	return i
}

func part1(data []string) int {
	instructions = data[0]
	directionsStrings := data[2:]

	for _, d := range directionsStrings {
		in := strings.Split(d, " = ")[0]
		out := strings.Split(strings.Split(strings.Split(d, "(")[1], ")")[0], ", ")

		if in[len(in)-1] == 'A' {
			startingPoints = append(startingPoints, in)
		}

		directions[in] = out
	}

	current := "AAA"

	return numberOfSteps(current)
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	l := int64(1)
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		a := int64(v)
		g := gcd(l, a)
		l = (l / g) * a
		if l < 0 {
			l = -l
		}
	}

	return int(l)
}

func part2(data []string) int {
	steps := []int{}
	for _, startingPoint := range startingPoints {
		steps = append(steps, numberOfSteps(startingPoint))
	}

	return lcm(steps)
}

func main() {
	path := utils.GetPath(os.Args)
	data := utils.ParseInputLines(path)

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
