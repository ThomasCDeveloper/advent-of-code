package main

// CMD: go run *.go

import (
	"fmt"
	"math"
	"strings"
)

type galaxy struct {
	x, y int
}

func (g galaxy) getDist(_g galaxy, hor []int, ver []int, expansion int) int {
	baseValue := int(math.Abs(float64(g.x-_g.x)) + math.Abs(float64(g.y-_g.y)))
	nL := 0
	for _, val := range ver {
		if math.Min(float64(g.x), float64(_g.x)) <= float64(val) && math.Max(float64(g.x), float64(_g.x)) >= float64(val) {
			nL++
		}
	}
	for _, val := range hor {
		if math.Min(float64(g.y), float64(_g.y)) <= float64(val) && math.Max(float64(g.y), float64(_g.y)) >= float64(val) {
			nL++
		}
	}

	return baseValue + nL*(expansion-1)
}

func getEmptyLines(data []string) ([]int, []int) {
	ver := []int{}
	hor := []int{}

	for y, line := range data {
		if !strings.Contains(line, "#") {
			hor = append(hor, y)
		}
	}
	for x := range data[0] {
		ve := true
		for _, line := range data {
			ve = ve && !(string(line[x]) == "#")
		}
		if ve {
			ver = append(ver, x)
		}
	}

	return hor, ver
}

func SolvePart1(data []string) int {
	hor, ver := getEmptyLines(data)
	//data = addLines(data, hor, ver)

	galaxies := []galaxy{}
	for y, line := range data {
		for x := range line {
			if string(line[x]) == "#" {
				galaxies = append(galaxies, galaxy{x, y})
			}
		}
	}
	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += galaxies[i].getDist(galaxies[j], hor, ver, 1)
		}
	}

	return sum
}

func SolvePart2(data []string) int {
	hor, ver := getEmptyLines(data)

	//data = addLines(data, hor, ver)

	galaxies := []galaxy{}
	for y, line := range data {
		for x := range line {
			if string(line[x]) == "#" {
				galaxies = append(galaxies, galaxy{x, y})
			}
		}
	}
	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += galaxies[i].getDist(galaxies[j], hor, ver, 999999)
		}
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
