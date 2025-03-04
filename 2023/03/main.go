package aoc03

import (
	"ThomasCDeveloper/advent-of-code/2023/utils"
	"fmt"
	"strconv"
	"strings"
)

var path = "inputs/03.input"

var data = []string{}
var windows = map[Window]int{}

func Run() {
	data = utils.ParseInputLines(path)
	for y, line := range data {
		currentNum := ""
		for x := 0; x < len(line); x++ {
			if strings.Index("0123456789", string(line[x])) != -1 {
				currentNum += string(line[x])
			} else {
				if currentNum != "" {
					value, _ := strconv.Atoi(currentNum)
					windows[Window{x - len(currentNum), y, len(currentNum)}] = value
				}
				currentNum = ""
			}
		}
		if currentNum != "" {
			value, _ := strconv.Atoi(currentNum)
			windows[Window{len(line) - len(currentNum), y, len(currentNum)}] = value
		}
	}

	SolveA()
	SolveB()
}

func GetRuneAtXY(data []string, x int, y int) rune {
	if x < 0 || y < 0 || x >= len(data[0]) || y >= len(data) {
		return '.'
	}
	return rune(data[y][x%len(data[y])])
}

type Window struct {
	x, y, length int
}

func SolveA() {
	sum := 0
	for window, val := range windows {
		partsConnected := 0
		for x := window.x - 1; x < window.x+window.length+1; x++ {
			for y := window.y - 1; y < window.y+2; y++ {
				r := GetRuneAtXY(data, x, y)
				if strings.Index("0123456789.", string(r)) == -1 {
					partsConnected++
				}
			}
		}
		if partsConnected > 0 {
			sum += val
		}
	}

	fmt.Printf("Solution to part 1: %d\n", sum)
}

func SolveB() {
	sum := 0
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if GetRuneAtXY(data, x, y) == '*' {
				currentNeighbors := []int{}
				for window, val := range windows {
					for x2 := window.x - 1; x2 < window.x+window.length+1; x2++ {
						for y2 := window.y - 1; y2 < window.y+2; y2++ {
							if x == x2 && y == y2 {
								currentNeighbors = append(currentNeighbors, val)
							}
						}
					}
				}
				if len(currentNeighbors) == 2 {
					sum += currentNeighbors[0] * currentNeighbors[1]
				}
			}
		}
	}

	fmt.Printf("Solution to part 2: %d\n", sum)
}
