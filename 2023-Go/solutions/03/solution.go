package main

import (
	"ThomasCDeveloper/advent-of-code/2023/utils"
	"fmt"
	"os"
	"strconv"
)

type window struct {
	minX, maxX, y int
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSymbol(x, y int, data []string) bool {
	if y < 0 || y >= len(data) || x < 0 || x >= len(data[y]) {
		return false
	}
	c := data[y][x]
	return !isDigit(c) && c != '.'
}

func extractNumbers(data []string) map[window]int {
	values := map[window]int{}

	for y, line := range data {
		for x := 0; x < len(line); x++ {
			if !isDigit(line[x]) {
				continue
			}

			start := x
			for x+1 < len(line) && isDigit(line[x+1]) {
				x++
			}

			numStr := line[start : x+1]
			v, _ := strconv.Atoi(numStr)

			values[window{start, x, y}] = v
		}
	}

	return values
}

func part1(data []string) int {
	valueWindow := extractNumbers(data)

	sum := 0

	for w, val := range valueWindow {
		found := false
		for y := w.y - 1; y <= w.y+1 && !found; y++ {
			for x := w.minX - 1; x <= w.maxX+1; x++ {
				if isSymbol(x, y, data) {
					found = true
					break
				}
			}
		}
		if found {
			sum += val
		}
	}

	return sum
}

func part2(data []string) int {
	valueWindow := extractNumbers(data)

	rowIndex := map[int][]window{}
	for w := range valueWindow {
		rowIndex[w.y] = append(rowIndex[w.y], w)
	}

	sum := 0

	for y, line := range data {
		for x := 0; x < len(line); x++ {
			if line[x] != '*' {
				continue
			}

			neighbors := []int{}

			for dy := -1; dy <= 1; dy++ {
				row := y + dy
				if row < 0 || row >= len(data) {
					continue
				}

				for _, w := range rowIndex[row] {
					if x >= w.minX-1 && x <= w.maxX+1 {
						neighbors = append(neighbors, valueWindow[w])
					}
				}
			}

			if len(neighbors) >= 2 {
				product := 1
				for _, v := range neighbors {
					product *= v
				}
				sum += product
			}
		}
	}

	return sum
}

func main() {
	path := utils.GetPath(os.Args)
	data := utils.ParseInputLines(path)

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
