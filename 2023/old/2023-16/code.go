package main

// CMD: go run *.go

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Beam struct {
	x, y, dir int // dir 0 right 1 down 2 left 3 up
}

var fullBeam = []Beam{}

func Walk(beam Beam, data []string) {
	if beam.x < 0 || beam.y < 0 || beam.x >= len(data[0]) || beam.y >= len(data) {
		return
	}
	if slices.Contains(fullBeam, beam) {
		return
	}

	fullBeam = append(fullBeam, beam)

	currentTile := data[beam.y][beam.x]

	if currentTile == '.' || (currentTile == '|' && beam.dir%2 == 1) || (currentTile == '-' && beam.dir%2 == 0) {
		switch beam.dir {
		case 0:
			Walk(Beam{beam.x + 1, beam.y, beam.dir}, data)
		case 1:
			Walk(Beam{beam.x, beam.y + 1, beam.dir}, data)
		case 2:
			Walk(Beam{beam.x - 1, beam.y, beam.dir}, data)
		default:
			Walk(Beam{beam.x, beam.y - 1, beam.dir}, data)
		}
		return
	}
	if currentTile == '\\' {
		switch beam.dir {
		case 0:
			Walk(Beam{beam.x, beam.y + 1, 1}, data)
		case 1:
			Walk(Beam{beam.x + 1, beam.y, 0}, data)
		case 2:
			Walk(Beam{beam.x, beam.y - 1, 3}, data)
		default:
			Walk(Beam{beam.x - 1, beam.y, 2}, data)
		}
	}
	if currentTile == '/' {
		switch beam.dir {
		case 0:
			Walk(Beam{beam.x, beam.y - 1, 3}, data)
		case 1:
			Walk(Beam{beam.x - 1, beam.y, 2}, data)
		case 2:
			Walk(Beam{beam.x, beam.y + 1, 1}, data)
		default:
			Walk(Beam{beam.x + 1, beam.y, 0}, data)
		}
	}
	if currentTile == '|' {
		Walk(Beam{beam.x, beam.y + 1, 1}, data)
		Walk(Beam{beam.x, beam.y - 1, 3}, data)
		return
	}
	if currentTile == '-' {
		Walk(Beam{beam.x + 1, beam.y, 0}, data)
		Walk(Beam{beam.x - 1, beam.y, 2}, data)
		return
	}
}

func SolvePart1(data []string, x int, y int, dir int) int {
	fullBeam = []Beam{}
	Walk(Beam{x, y, dir}, data)
	for i, beam := range fullBeam {
		beam.dir = 0
		fullBeam[i] = beam
	}

	walkedTiles := []Beam{}
	for _, beam := range fullBeam {
		if !slices.Contains(walkedTiles, beam) {
			walkedTiles = append(walkedTiles, beam)
		}
	}

	return len(walkedTiles)
}

func SolvePart2(data []string) int {
	fullBeam = []Beam{}
	max := 0

	for y := range data {
		val := SolvePart1(data, 0, y, 0)
		if val > max {
			max = val
		}
		val = SolvePart1(data, len(data[0])-1, y, 2)
		if val > max {
			max = val
		}
	}
	for x := range data[0] {
		val := SolvePart1(data, x, 0, 1)
		if val > max {
			max = val
		}
		val = SolvePart1(data, x, len(data)-1, 3)
		if val > max {
			max = val
		}
	}

	return max
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data, 0, 0, 0))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
