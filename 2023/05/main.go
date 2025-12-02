package aoc05

import (
	"ThomasCDeveloper/advent-of-code/2023/utils"
	"fmt"
	"strings"
)

var path = "inputs/05.input.test"

func Run() {
	SolveA()
	SolveB()
}

type Map struct {
	ranges [][3]int
}

func (m Map) ProcessNumber(n int) int {
	for _, subMap := range m.ranges {
		if n >= subMap[0] && n < subMap[0]+subMap[2] {
			return n - subMap[0] + subMap[1]
		}
	}
	return n
}

func (m Map) ProcessRanges(r [][2]int) [][2]int {
	intervals := r
	for _, subMap := range m.ranges {
		var newIntervals [][2]int
		for _, iv := range intervals {
			pieces := applyMappingRule(iv, subMap)
			newIntervals = append(newIntervals, pieces...)
		}
		intervals = newIntervals
	}
	return intervals
}

func applyMappingRule(iv [2]int, subMap [3]int) [][2]int {
	a, b := iv[0], iv[1]
	srcStart := subMap[0]
	srcEnd := subMap[0] + subMap[2] - 1
	var res [][2]int

	// If no overlap, return the interval unchanged.
	if b < srcStart || a > srcEnd {
		res = append(res, iv)
		return res
	}

	// Left part: values before the mapping rule.
	if a < srcStart {
		res = append(res, [2]int{a, srcStart - 1})
	}

	// Overlapping part: the intersection.
	overlapStart := max(a, srcStart)
	overlapEnd := min(b, srcEnd)
	newStart := subMap[1] + (overlapStart - srcStart)
	newEnd := subMap[1] + (overlapEnd - srcStart)
	res = append(res, [2]int{newStart, newEnd})

	// Right part: values after the overlapping portion.
	if b > overlapEnd {
		res = append(res, [2]int{overlapEnd + 1, b})
	}
	return res
}

func MapOfString(s string) Map {
	m := Map{}
	for _, e := range strings.Split(s, "\n")[1:] {
		m.ranges = append(m.ranges, [3]int{
			utils.Atoi(strings.Split(e, " ")[1]),
			utils.Atoi(strings.Split(e, " ")[0]),
			utils.Atoi(strings.Split(e, " ")[2]),
		})
	}
	return m
}

func SolveA() {
	data := strings.Split(utils.ParseRawInput(path), "\n\n")
	seedsString := strings.Split(strings.Split(data[0], ": ")[1], " ")

	maps := data[1:]
	min := 999999999999
	for _, sString := range seedsString {
		seed := utils.Atoi(sString)
		for _, mS := range maps {
			m := MapOfString(mS)
			seed = m.ProcessNumber(seed)
		}
		if seed < min {
			min = seed
		}
	}

	fmt.Printf("Solution to part 1: %d\n", min)
}

func SolveB() {
	data := strings.Split(utils.ParseRawInput(path), "\n\n")

	seedsParts := strings.Split(strings.Split(data[0], ": ")[1], " ")
	var seedIntervals [][2]int
	for i := 0; i < len(seedsParts); i += 2 {
		start := utils.Atoi(seedsParts[i])
		length := utils.Atoi(seedsParts[i+1])
		seedIntervals = append(seedIntervals, [2]int{start, start + length - 1})
	}

	mapsData := data[1:]
	intervals := seedIntervals
	fmt.Println(intervals)
	for _, mS := range mapsData {
		m := MapOfString(mS)
		fmt.Println(m)
		intervals = m.ProcessRanges(intervals)
		fmt.Println(intervals)
	}

	min := 999999999999
	for _, iv := range intervals {
		if iv[0] < min {
			min = iv[0]
		}
	}

	fmt.Printf("Solution to part 2: %d\n", min)
}
