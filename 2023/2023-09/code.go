package main

// CMD: go run *.go

import (
	"fmt"
	"strconv"
	"strings"
)

func GetDiffLevel(values []int, _start []int) (int, []int) {
	_start = append(_start, values[0])
	nb0 := 0
	for _, val := range values {
		if val == values[0] {
			nb0++
		}
	}
	if nb0 == len(values) {
		return len(values), _start
	}

	newValues := []int{}
	for i := 0; i < len(values)-1; i++ {
		newValues = append(newValues, values[i+1]-values[i])
	}

	return GetDiffLevel(newValues, _start)
}

func SolvePart1(data []string) int {
	sum := 0
	for _, line := range data {
		values := []int{}
		for _, strVal := range strings.Split(line, " ") {
			intVal, _ := strconv.Atoi(strVal)
			values = append(values, intVal)
		}

		gotDiff, starts := GetDiffLevel(values, []int{})

		lines := make([][]int, len(values)-gotDiff+1)
		for i := range lines {
			lines[i] = make([]int, len(values)+1)
			lines[i][0] = starts[i]
			if i == len(lines)-1 {
				for j := range lines[i] {
					lines[i][j] = starts[i]
				}
			}
		}
		for i := range lines[1:] {
			i = len(lines) - i - 1
			for j := 1; j < len(lines[0]); j++ {
				lines[i-1][j] = lines[i-1][j-1] + lines[i][j-1]
			}
		}
		sum += lines[0][len(lines[0])-1]
	}
	return sum
}

func SolvePart2(data []string) int {
	sum := 0
	for _, line := range data {
		values := []int{}
		for _, strVal := range strings.Split(line, " ") {
			intVal, _ := strconv.Atoi(strVal)
			values = append(values, intVal)
		}

		gotDiff, starts := GetDiffLevel(values, []int{})

		lines := make([][]int, len(values)-gotDiff+1)
		for i := range lines {
			lines[i] = make([]int, 2)
			lines[i][1] = starts[i]
			if i == len(lines)-1 {
				lines[i][0] = starts[i]
			}
		}
		for i := range lines[1:] {
			i = len(lines) - i - 1
			lines[i-1][0] = lines[i-1][1] - lines[i][0]

		}
		sum += lines[0][0]
	}
	return sum
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
