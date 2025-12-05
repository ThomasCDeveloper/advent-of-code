package main

// CMD: go run *.go

import (
	"fmt"
	"strings"
)

type Node struct {
	childs []string
}

func PGCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func PPCM(a, b int) int {
	result := a * b / PGCD(a, b)

	return result
}

func GetStepOf(node string, nodes map[string]Node, cmd string) int {
	currentNode := node
	steps := 0

	for {
		for _, char := range cmd {
			if string(char) == "L" {
				currentNode = nodes[currentNode].childs[0]
			} else {
				currentNode = nodes[currentNode].childs[1]
			}
			steps++
			if string(currentNode[2]) == "Z" {
				return steps
			}
		}
	}
}

func SolvePart1(data []string) int {
	cmd := data[0]

	nodes := map[string]Node{}

	nodesDescription := data[2:]
	for i := range nodesDescription {
		nodeName := strings.Split(nodesDescription[i], " ")[0]
		nodeChild1 := strings.Split(strings.Split(nodesDescription[i], " (")[1], ",")[0]
		nodeChild2 := strings.Split(strings.Split(nodesDescription[i], ", ")[1], ")")[0]

		nodes[nodeName] = Node{[]string{nodeChild1, nodeChild2}}
	}

	return GetStepOf("AAA", nodes, cmd)
}

func SolvePart2(data []string) int {
	cmd := data[0]

	nodes := map[string]Node{}

	nodesDescription := data[2:]
	for i := range nodesDescription {
		nodeName := strings.Split(nodesDescription[i], " ")[0]
		nodeChild1 := strings.Split(strings.Split(nodesDescription[i], " (")[1], ",")[0]
		nodeChild2 := strings.Split(strings.Split(nodesDescription[i], ", ")[1], ")")[0]

		nodes[nodeName] = Node{[]string{nodeChild1, nodeChild2}}
	}

	currentNodes := []string{}

	for key := range nodes {
		if string(key[2]) == "A" {
			currentNodes = append(currentNodes, key)
		}
	}

	steps := []int{}

	for i := 0; i < len(currentNodes); i++ {
		steps = append(steps, GetStepOf(currentNodes[i], nodes, cmd))
	}

	ppcm := steps[0]
	for i := 1; i < len(steps); i++ {
		ppcm = PPCM(ppcm, steps[i])
	}

	return ppcm

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
