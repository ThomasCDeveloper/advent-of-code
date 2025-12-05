package main

// CMD: go run *.go

import (
	"fmt"
	"slices"
)

type co struct {
	r, c int
}

var tiles = map[co]bool{}
var distances = map[co]int{}

func printSituation(list []co) {
	fmt.Println()
	for r := range data {
		line := ""
		for c := range data[0] {
			if slices.Contains(list, co{r, c}) {
				line += "O"
			} else {
				if tiles[co{r, c}] {
					line += " "
				} else {
					line += "#"
				}
			}
		}
		fmt.Println(line)
	}
}

func deepCopy(list []co) []co {
	out := []co{}
	out = append(out, list...)
	return out
}

func getTile(rc co, loop bool) bool {
	if loop {
		for rc.c < 0 {
			rc.c += len(data)
		}
		for rc.r < 0 {
			rc.r += len(data[0])
		}
		return tiles[co{rc.r % len(data), rc.c % len(data[0])}]
	}
	if rc.c < 0 || rc.r < 0 || rc.r >= len(data) || rc.c >= len(data[0]) {
		return false
	}
	return tiles[rc]
}

func SolvePart1(data []string) int {
	startCo := co{0, 0}
	for r, line := range data {
		for c, char := range line {
			rcco := co{r, c}
			if char == 'S' {
				startCo = co{r, c}
			}
			if char == '#' {
				tiles[rcco] = false
			} else {
				tiles[rcco] = true
			}
		}
	}

	visitedTiles := []co{startCo}
	distances[startCo] = 0

	for i := 0; i < 131; i++ {
		newVisitedTiles := []co{}
		for _, visitedTile := range visitedTiles {
			for _, dir := range []co{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
				rc := co{visitedTile.r + dir.r, visitedTile.c + dir.c}
				if getTile(rc, false) {
					if !slices.Contains(newVisitedTiles, rc) {
						newVisitedTiles = append(newVisitedTiles, rc)
					}
					if val, check := distances[rc]; val == 0 || !check {
						distances[rc] = i + 1
					}
				}
			}
		}
		visitedTiles = newVisitedTiles
	}

	sum := 0

	for key, val := range distances {
		if (key.c+key.r)%2 == 0 && val <= 65 {
			sum += 1
		}
	}

	return sum
}

func SolvePart2(data []string) int {
	evenCorners, oddCorners, evenFull, oddFull := 0, 0, 0, 0

	for key, val := range distances {
		if val > 65 {
			if (key.c+key.r)%2 == 0 {
				evenCorners++
			} else {
				oddCorners++
			}
		}
		if (key.c+key.r)%2 == 0 {
			evenFull++
		} else {
			oddFull++
		}
	}

	n := 202300

	return ((n+1)*(n+1)*oddFull + n*n*evenFull - (n+1)*oddCorners + n*evenCorners)
}

var data = []string{}

func main() {
	data = GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
