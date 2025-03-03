package main

// CMD: go run *.go

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type co struct {
	r int
	c int
}

func printTiles(tiles [][]bool) {
	for r := 0; r < len(tiles); r++ {
		line := ""
		for c := 0; c < len(tiles[0]); c++ {
			if tiles[r][c] {
				line += "# "
			} else {
				line += "  "
			}
		}
		fmt.Println(line)
	}
}

func SolvePart1(data []string) int {
	listTiles := map[co]bool{}

	currentPos := co{0, 0}
	dirs := map[string]co{"R": {0, 1}, "L": {0, -1}, "U": {-1, 0}, "D": {1, 0}}

	for _, line := range data {
		dir := strings.Split(line, " ")[0]
		n, _ := strconv.Atoi(strings.Split(line, " ")[1])

		for i := 0; i < n; i++ {
			listTiles[currentPos] = true

			currentPos.c += dirs[dir].c
			currentPos.r += dirs[dir].r
		}
	}

	n := len(listTiles)

	minr, minc, maxr, maxc := currentPos.r, currentPos.c, currentPos.r, currentPos.c
	for key := range listTiles {
		if key.c > maxc {
			maxc = key.c
		}
		if key.c < minc {
			minc = key.c
		}
		if key.r > maxr {
			maxr = key.r
		}
		if key.r < minr {
			minr = key.r
		}
	}

	tiles := [][]bool{}
	for r := 0; r < maxr-minr+3; r++ {
		line := make([]bool, maxc-minc+3)
		tiles = append(tiles, line)
	}

	for key := range listTiles {
		tiles[key.r-minr+1][key.c-minc+1] = true
	}

	startFill := co{0, 0}
	neighbors := []co{startFill}
	for len(neighbors) > 0 {
		pos := neighbors[0]
		tiles[pos.r][pos.c] = true

		for _, dir := range []co{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nr, nc := pos.r+dir.r, pos.c+dir.c

			if nr >= 0 && nr < len(tiles) && nc >= 0 && nc < len(tiles[0]) {
				if !tiles[nr][nc] {
					if !slices.Contains(neighbors, co{nr, nc}) {
						neighbors = append(neighbors, co{nr, nc})
					}
				}
			}
		}

		neighbors = neighbors[1:]
	}

	sum := len(tiles)*len(tiles[0]) + n
	for r := 0; r < len(tiles); r++ {
		for c := 0; c < len(tiles[0]); c++ {
			if tiles[r][c] {
				sum -= 1
			}
		}
	}

	return sum
}

func SolvePart2(data []string) int {
	points := []co{}

	currentPos := co{0, 0}
	listDirs := []string{"R", "D", "L", "U"}
	dirs := map[string]co{"R": {0, 1}, "L": {0, -1}, "U": {-1, 0}, "D": {1, 0}}

	lenLine := 0
	for _, line := range data {
		points = append(points, currentPos)
		hexa := strings.Split(strings.Split(line, "#")[1], ")")[0]

		dirInt, _ := strconv.Atoi(string(hexa[len(hexa)-1]))
		dir := listDirs[dirInt]
		n, _ := strconv.ParseInt(hexa[:len(hexa)-1], 16, 64)

		lenLine += int(n)

		currentPos.c += dirs[dir].c * int(n)
		currentPos.r += dirs[dir].r * int(n)
	}

	sum := 0
	for i := range points[1:] {
		sum += points[i].c*points[i+1].r - points[i].r*points[i+1].c
	}

	return sum/2 + lenLine/2 + 1
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
