package main

// CMD: go run *.go

import (
	"fmt"
	"slices"
	"strings"
)

type Pos struct {
	x, y int
}

type Node struct {
	neighbors []Pos
}

func SolveMaze(startPos Pos, nodeMap map[Pos]Node) int {
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

	return steps / 2
}

func CreateNodes(data []string) map[Pos]Node {
	nodes := map[Pos]Node{}
	for _y, line := range data {
		for _x, char := range line {
			neighbors := []Pos{}
			switch string(char) {
			case "S":
				if _x > 0 {
					if strings.Contains("-FL", string(data[_y][_x-1])) {
						neighbors = append(neighbors, Pos{_x - 1, _y})
					}
				}
				if _x < len(line)-1 {
					if strings.Contains("-J7", string(data[_y][_x+1])) {
						neighbors = append(neighbors, Pos{_x + 1, _y})
					}
				}
				if _y > 0 {
					if strings.Contains("7F|", string(data[_y-1][_x])) {
						neighbors = append(neighbors, Pos{_x, _y - 1})
					}
				}
				if _y < len(data)-1 {
					if strings.Contains("LJ|", string(data[_y+1][_x])) {
						neighbors = append(neighbors, Pos{_x, _y + 1})
					}
				}
			case "|":
				if _y > 0 {
					if strings.Contains("7SF|", string(data[_y-1][_x])) {
						neighbors = append(neighbors, Pos{_x, _y - 1})
					}
				}
				if _y < len(data)-1 {
					if strings.Contains("LSJ|", string(data[_y+1][_x])) {
						neighbors = append(neighbors, Pos{_x, _y + 1})
					}
				}
			case "-":
				if _x > 0 {
					if strings.Contains("-FSL", string(data[_y][_x-1])) {
						neighbors = append(neighbors, Pos{_x - 1, _y})
					}
				}
				if _x < len(line)-1 {
					if strings.Contains("-SJ7", string(data[_y][_x+1])) {
						neighbors = append(neighbors, Pos{_x + 1, _y})
					}
				}
			case "L":
				if _y > 0 {
					if strings.Contains("7FS|", string(data[_y-1][_x])) {
						neighbors = append(neighbors, Pos{_x, _y - 1})
					}
				}
				if _x < len(line)-1 {
					if strings.Contains("-SJ7", string(data[_y][_x+1])) {
						neighbors = append(neighbors, Pos{_x + 1, _y})
					}
				}
			case "J":
				if _x > 0 {
					if strings.Contains("-FSL", string(data[_y][_x-1])) {
						neighbors = append(neighbors, Pos{_x - 1, _y})
					}
				}
				if _y > 0 {
					if strings.Contains("S7F|", string(data[_y-1][_x])) {
						neighbors = append(neighbors, Pos{_x, _y - 1})
					}
				}
			case "7":
				if _x > 0 {
					if strings.Contains("-FSL", string(data[_y][_x-1])) {
						neighbors = append(neighbors, Pos{_x - 1, _y})
					}
				}
				if _y < len(data)-1 {
					if strings.Contains("LJS|", string(data[_y+1][_x])) {
						neighbors = append(neighbors, Pos{_x, _y + 1})
					}
				}
			case "F":
				if _x < len(line)-1 {
					if strings.Contains("-JS7", string(data[_y][_x+1])) {
						neighbors = append(neighbors, Pos{_x + 1, _y})
					}
				}
				if _y < len(data)-1 {
					if strings.Contains("LJ|S", string(data[_y+1][_x])) {
						neighbors = append(neighbors, Pos{_x, _y + 1})
					}
				}
			default:
				continue
			}

			nodes[Pos{_x, _y}] = Node{neighbors}
		}
	}
	return nodes
}

func SolvePart1(data []string) int {
	x, y := 0, 0
	for _y, line := range data {
		_x := strings.Index(line, "S")
		if _x != -1 {
			x, y = _x, _y
			break
		}
	}

	nodes := CreateNodes(data)

	return SolveMaze(Pos{x, y}, nodes)
}

func GetMaze(startPos Pos, nodeMap map[Pos]Node) []Pos {
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

func SolvePart2(data []string) int {
	x, y := 0, 0
	for _y, line := range data {
		_x := strings.Index(line, "S")
		if _x != -1 {
			x, y = _x, _y
			break
		}
	}

	nodes := GetMaze(Pos{x, y}, CreateNodes(data))

	insideNodes := []Pos{}
	for y := 1; y < len(data)-1; y++ {
		for x := 1; x < len(data[0])-1; x++ {
			if !slices.Contains(nodes, Pos{x, y}) {
				walls := 0
				for dx := x + 1; dx < len(data[0]); dx++ {
					if slices.Contains(nodes, Pos{dx, y}) {
						if strings.Contains("JL|S", string(data[y][dx])) {
							walls++
						}
					}
				}
				if walls%2 == 1 {
					insideNodes = append(insideNodes, Pos{x, y})
				}
			}
		}
	}

	return len(insideNodes)
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
