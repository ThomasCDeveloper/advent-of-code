package main

// CMD: go run *.go

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type window struct {
	minX, maxX, y int
}

func isSymbol(x int, y int, data []string) bool {
	if x < 0 || y < 0 || y >= len(data) || x >= len(data[0]) {
		return false
	}

	if strings.Index("0123456789.", string(data[y][x])) == -1 {
		return true
	}

	return false
}

func SolvePart1(data []string) int {
	valueWindow := map[window]int{}
	for y := 0; y < len(data); y++ {
		line := data[y]

		for x := 0; x < len(line); x++ {
			dx := 0
			if strings.Index("0123456789", string(line[x])) != -1 {
				for strings.Index("0123456789", string(line[x+dx])) != -1 && x+dx+1 < len(line) {
					dx++
				}
				if x+dx == len(line)-1 {
					if line[len(line)-1] == '.' {
						valueWindow[window{x, x + dx - 1, y}], _ = strconv.Atoi(line[x:x+dx] + "")
					} else {
						valueWindow[window{x, x + dx, y}], _ = strconv.Atoi(line[x:x+dx+1] + "")
					}
				} else {
					valueWindow[window{x, x + dx - 1, y}], _ = strconv.Atoi(line[x:x+dx] + "")
				}
				x = x + dx
			}
		}
	}

	sum := 0
	for key, val := range valueWindow {
		compt := 0
		for y := key.y - 1; y < key.y+2; y++ {
			for x := key.minX - 1; x < key.maxX+2; x++ {
				if isSymbol(x, y, data) {
					compt += 1
				}
			}
		}
		if compt != 0 {
			sum += val
		}
	}
	return sum
}

func SolvePart2(data []string) int {
	valueWindow := map[window]int{}
	for y := 0; y < len(data); y++ {
		line := data[y]

		curatedLine := line
		for _, charToDelete := range "#+$.%/-@=&" {
			curatedLine = strings.ReplaceAll(curatedLine, string(charToDelete), " ")
		}
		for x := 0; x < len(line); x++ {
			dx := 0
			if strings.Index("0123456789", string(line[x])) != -1 {
				for strings.Index("0123456789", string(line[x+dx])) != -1 && x+dx+1 < len(line) {
					dx++
				}
				if x+dx == len(line)-1 {
					if line[len(line)-1] == '.' {
						valueWindow[window{x, x + dx - 1, y}], _ = strconv.Atoi(line[x:x+dx] + "")
					} else {
						valueWindow[window{x, x + dx, y}], _ = strconv.Atoi(line[x:x+dx+1] + "")
					}
				} else {
					valueWindow[window{x, x + dx - 1, y}], _ = strconv.Atoi(line[x:x+dx] + "")
				}
				x = x + dx
			}
		}
	}

	sum := 0

	for y := 0; y < len(data); y++ {
		line := data[y]

		for x := 0; x < len(line); x++ {
			if line[x] == '*' {
				neighbors := []int{}
				for key, val := range valueWindow {
					if math.Abs(float64(key.y-y)) <= 1 {
						if x >= key.minX-1 && x <= key.maxX+1 {
							neighbors = append(neighbors, val)
						}
					}
				}

				if len(neighbors) >= 2 {
					product := 1
					for _, val := range neighbors {
						product *= val
					}
					sum += product
				}
			}
		}
	}

	return sum
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
