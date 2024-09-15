package main

// CMD: go run *.go

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

type brick struct {
	x, y, z    int
	dx, dy, dz int
	idx        int
}

type co struct {
	x, y, z int
}

func dropBrick(b brick, stack map[co]int, zsum int, supports map[int][]int) {
	minZ := 0
	for _x := b.x; _x <= b.dx; _x++ {
		for _y := b.y; _y <= b.dy; _y++ {
			for _z := zsum; _z > -5; _z-- {
				if _, check := stack[co{_x, _y, _z - 1}]; check || _z == 0 {
					minZ = max(minZ, _z)
				}
			}
		}
	}

	for _x := b.x; _x <= b.dx; _x++ {
		for _z := 0; _z <= b.dz-b.z; _z++ {
			for _y := b.y; _y <= b.dy; _y++ {
				stack[co{_x, _y, _z + minZ}] = b.idx

				if v, check := stack[co{_x, _y, _z + minZ - 1}]; check {
					if v != b.idx {
						if !slices.Contains(supports[b.idx], v) {
							supports[b.idx] = append(supports[b.idx], v)
						}
						if !slices.Contains(supports2[v], b.idx) {
							supports2[v] = append(supports2[v], b.idx)
						}
					}
				}
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var bricks = []brick{}
var stack = map[co]int{}
var supports = map[int][]int{}
var supports2 = map[int][]int{}
var canBeRemoved = map[int]bool{}

func SolvePart1(data []string) int {
	zsum := 0

	re := regexp.MustCompile("[0-9]+")
	for i, line := range data {
		coords := re.FindAllString(line, -1)

		c0, _ := strconv.Atoi(coords[0])
		c1, _ := strconv.Atoi(coords[1])
		c2, _ := strconv.Atoi(coords[2])
		c3, _ := strconv.Atoi(coords[3])
		c4, _ := strconv.Atoi(coords[4])
		c5, _ := strconv.Atoi(coords[5])

		x, dx := min(c0, c3), max(c0, c3)
		y, dy := min(c1, c4), max(c1, c4)
		z, dz := min(c2, c5), max(c2, c5)

		zsum += dz

		bricks = append(bricks, brick{x, y, z, dx, dy, dz, i})
		canBeRemoved[i] = true
	}

	slices.SortFunc(bricks, func(a, b brick) int {
		return a.z - b.z
	})

	for _, b := range bricks {
		dropBrick(b, stack, zsum, supports)
	}

	for _, sup := range supports {
		if len(sup) == 1 {
			canBeRemoved[sup[0]] = false
		}
	}

	sum := 0
	for _, v := range canBeRemoved {
		if v {
			sum += 1
		}
	}

	return sum
}

var broken = []int{}

func GetFalls(idx int) int {
	broken = append(broken, idx)

	sum := 0
	for _, sup := range supports2[idx] {
		nSup := supports[sup]
		falling := true
		for _, i := range nSup {
			if !slices.Contains(broken, i) {
				falling = false
			}
		}
		if falling {
			sum += GetFalls(sup) + 1
		}
	}
	return sum
}

func SolvePart2(data []string) int {
	sum := 0
	for i := range data {
		broken = []int{}
		sum += GetFalls(i)
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
