package main

import (
	"ThomasCDeveloper/advent-of-code/2023/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solveFor(T, r int) int {
	a := float64(1)
	b := float64(-T)
	c := float64(r + 1)

	D := b*b - 4*a*c
	x1 := math.Floor((-b + math.Sqrt(D)) / 2)
	x2 := math.Ceil((-b - math.Sqrt(D)) / 2)

	return int(x1 - x2 + 1)
}

func part1(data []string) int {
	for strings.Contains(data[0], "  ") {
		data0 := strings.ReplaceAll(data[0], "  ", " ")
		data[0] = data0
	}
	for strings.Contains(data[1], "  ") {
		data1 := strings.ReplaceAll(data[1], "  ", " ")
		data[1] = data1
	}
	durationsS := strings.Split(strings.Split(data[0], ": ")[1], " ")
	recordsS := strings.Split(strings.Split(data[1], ": ")[1], " ")

	fmt.Println(durationsS)

	durations := []int{}
	records := []int{}

	for i := range durationsS {
		val, _ := strconv.Atoi(durationsS[i])
		durations = append(durations, val)
		val, _ = strconv.Atoi(recordsS[i])
		records = append(records, val)
	}

	sum := 1
	for i := range durationsS {
		T := durations[i]
		r := records[i]

		sum *= solveFor(T, r)
	}

	return sum
}

func part2(data []string) int {
	for strings.Contains(data[0], "  ") {
		data0 := strings.ReplaceAll(data[0], "  ", " ")
		data[0] = data0
	}
	for strings.Contains(data[1], "  ") {
		data1 := strings.ReplaceAll(data[1], "  ", " ")
		data[1] = data1
	}
	duration := utils.Atoi(strings.Join(strings.Split(strings.Split(data[0], ": ")[1], " "), ""))
	record := utils.Atoi(strings.Join(strings.Split(strings.Split(data[1], ": ")[1], " "), ""))

	return solveFor(duration, record)
}

func main() {
	path := utils.GetPath(os.Args)
	data := utils.ParseInputLines(path)

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
