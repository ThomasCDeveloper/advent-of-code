package main

import (
	"ThomasCDeveloper/advent-of-code/2023/utils"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

type Node struct {
	neighbors []Pos
}

func getPipeLength(startPos Pos, nodeMap map[Pos]Node) int {
	steps := 1
	lastPos := startPos
	nextNode := nodeMap[startPos].neighbors[0]

	for nextNode != startPos {
		if nodeMap[nextNode].neighbors[0] != lastPos {
			lastPos = nextNode
			nextNode = nodeMap[nextNode].neighbors[0]
		} else {
			lastPos = nextNode
			nextNode = nodeMap[nextNode].neighbors[1]
		}
		steps++
	}

	return steps
}

var dirs = map[string]map[[2]int]string{
	"S": {
		{0, -1}: "7F|",
		{0, 1}:  "LJ|",
		{-1, 0}: "-FL",
		{1, 0}:  "-J7",
	},
	"|": {
		{0, -1}: "7SF|",
		{0, 1}:  "LSJ|",
	},
	"-": {
		{-1, 0}: "-FSL",
		{1, 0}:  "-SJ7",
	},
	"L": {
		{0, -1}: "7FS|",
		{1, 0}:  "-SJ7",
	},
	"J": {
		{-1, 0}: "-FSL",
		{0, -1}: "S7F|",
	},
	"7": {
		{-1, 0}: "-FSL",
		{0, 1}:  "LJS|",
	},
	"F": {
		{1, 0}: "-JS7",
		{0, 1}: "LJ|S",
	},
}

func createNodes(data []string) map[Pos]Node {
	nodes := map[Pos]Node{}
	for _y, line := range data {
		for _x, char := range line {
			neighbors := []Pos{}

			rules, ok := dirs[string(char)]
			if ok {
				for d, allowed := range rules {
					nx := _x + d[0]
					ny := _y + d[1]

					if nx < 0 || ny < 0 || ny >= len(data) || nx >= len(line) {
						continue
					}

					if strings.Contains(allowed, string(data[ny][nx])) {
						neighbors = append(neighbors, Pos{nx, ny})
					}
				}
			}

			nodes[Pos{_x, _y}] = Node{neighbors}
		}
	}
	return nodes
}

func part1(data []string) int {
	x, y := 0, 0
	for _y, line := range data {
		_x := strings.Index(line, "S")
		if _x != -1 {
			x, y = _x, _y
			break
		}
	}
	nodes := createNodes(data)

	return getPipeLength(Pos{x, y}, nodes) / 2
}

func getMaze(startPos Pos, nodeMap map[Pos]Node) []Pos {
	nodes := []Pos{startPos}

	lastPos := startPos
	nextNode := nodeMap[startPos].neighbors[0]

	for nextNode != startPos {
		nodes = append(nodes, nextNode)
		if nodeMap[nextNode].neighbors[0] != lastPos {
			lastPos = nextNode
			nextNode = nodeMap[nextNode].neighbors[0]
		} else {
			lastPos = nextNode
			nextNode = nodeMap[nextNode].neighbors[1]
		}
	}
	return nodes
}

func part2(data []string) int {
	total := 0
	x, y := 0, 0
	for _y, line := range data {
		_x := strings.Index(line, "S")
		if _x != -1 {
			x, y = _x, _y
			break
		}
	}

	nodes := getMaze(Pos{x, y}, createNodes(data))
	inCycle := make(map[Pos]bool)
	for _, p := range nodes {
		inCycle[p] = true
	}

	for y := 1; y < len(data)-1; y++ {
		walls := 0
		for x := 1; x < len(data[0])-1; x++ {
			if inCycle[Pos{x, y}] {
				switch data[y][x] {
				case 'J', 'L', '|', 'S':
					walls++
				}
			} else {
				if walls%2 == 1 {
					total++
				}
			}
		}
	}

	return total
}

func main() {
	path := utils.GetPath(os.Args)
	data := utils.ParseInputLines(path)

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
