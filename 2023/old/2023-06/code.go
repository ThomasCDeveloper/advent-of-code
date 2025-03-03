package main

// CMD: go run *.go

import (
	"fmt"
	"math"
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

func SolvePart1(data []string) int {
	durationsS := strings.Split(data[0], " ")
	recordsS := strings.Split(data[1], " ")

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

func SolvePart2(data []string) int {
	duration, _ := strconv.Atoi(strings.Join(strings.Split(data[0], " "), ""))
	record, _ := strconv.Atoi(strings.Join(strings.Split(data[1], " "), ""))

	return solveFor(duration, record)
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:")
	fmt.Println(SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:")
	fmt.Println(SolvePart2(data))
}
