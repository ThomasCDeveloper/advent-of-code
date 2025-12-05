package main

// CMD: go run *.go

import (
	"fmt"
	"math"

	"slices"
)

type terrain struct {
	ver []int
	hor []int
}

func getTerrain(lines []string) terrain {
	hors := []int{}
	for _, line := range lines {
		hor := 0
		for _, char := range line {
			hor *= 2
			if string(char) == "#" {
				hor += 1
			}
		}
		hors = append(hors, hor)
	}

	vers := make([]int, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		for _, line := range lines {
			vers[i] *= 2
			char := string(line[i])
			if char == "#" {
				vers[i] += 1
			}
		}
	}

	return terrain{vers, hors}
}

func getMirrors(ter terrain) int {
	verMirror := 0
	for i := 1; i < len(ter.ver); i++ {
		isMirror := true
		for j := 0; j < int(math.Min(float64(i), float64(len(ter.ver)-i))); j++ {
			isMirror = isMirror && (ter.ver[i+j] == ter.ver[i-j-1])
			if !isMirror {
				break
			}
		}
		if isMirror {
			verMirror += i
		}
	}

	horMirror := 0
	for i := 1; i < len(ter.hor); i++ {
		isMirror := true
		for j := 0; j < int(math.Min(float64(i), float64(len(ter.hor)-i))); j++ {
			isMirror = isMirror && (ter.hor[i+j] == ter.hor[i-j-1])
			if !isMirror {
				break
			}
		}
		if isMirror {
			horMirror += i
		}
	}

	return horMirror*100 + verMirror
}

func getMirrors2(ter terrain) int {
	powerOf2 := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768}
	verMirror := 0
	for i := 1; i < len(ter.ver); i++ {
		isMirror := true
		diffs := 0
		for j := 0; j < int(math.Min(float64(i), float64(len(ter.ver)-i))); j++ {
			is1diff := slices.Contains(powerOf2, int(math.Abs(float64(ter.ver[i+j])-float64(ter.ver[i-j-1]))))
			is0diff := (ter.ver[i+j] == ter.ver[i-j-1])
			isMirror = isMirror && (is1diff || is0diff)

			if is1diff {
				diffs += 1
			}

			if !isMirror || diffs > 1 {
				break
			}
		}
		if isMirror && diffs == 1 {
			verMirror += i
		}
	}

	horMirror := 0
	for i := 1; i < len(ter.hor); i++ {
		isMirror := true
		diffs := 0
		for j := 0; j < int(math.Min(float64(i), float64(len(ter.hor)-i))); j++ {
			is1diff := slices.Contains(powerOf2, int(math.Abs(float64(ter.hor[i+j])-float64(ter.hor[i-j-1]))))
			is0diff := (ter.hor[i+j] == ter.hor[i-j-1])
			isMirror = isMirror && (is1diff || is0diff)

			if is1diff {
				diffs += 1
			}

			if !isMirror || diffs > 1 {
				break
			}
		}
		if isMirror && diffs == 1 {
			horMirror += i
		}
	}

	return horMirror*100 + verMirror
}

func SolvePart1(data []string) int {
	terrains := []terrain{}
	newTerrainLines := []string{}

	for _, line := range data {
		if len(line) != 0 {
			newTerrainLines = append(newTerrainLines, line)
		} else {
			terrains = append(terrains, getTerrain(newTerrainLines))
			newTerrainLines = []string{}
		}
	}
	terrains = append(terrains, getTerrain(newTerrainLines))

	sum := 0
	for _, ter := range terrains {
		sum += getMirrors(ter)
	}

	return sum
}

func SolvePart2(data []string) int {
	terrains := []terrain{}
	newTerrainLines := []string{}

	for _, line := range data {
		if len(line) != 0 {
			newTerrainLines = append(newTerrainLines, line)
		} else {
			terrains = append(terrains, getTerrain(newTerrainLines))
			newTerrainLines = []string{}
		}
	}
	terrains = append(terrains, getTerrain(newTerrainLines))

	sum := 0
	for _, ter := range terrains {
		sum += getMirrors2(ter)
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
