package main

// CMD: go run *.go

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type part struct {
	x, m, a, s int // change to array ?
}

type partrange struct {
	minmax [][]int
}

func (p partrange) count() int {
	return (p.minmax[0][1] - p.minmax[0][0] + 1) * (p.minmax[1][1] - p.minmax[1][0] + 1) * (p.minmax[2][1] - p.minmax[2][0] + 1) * (p.minmax[3][1] - p.minmax[3][0] + 1)
}

type cmd struct {
	v     int // 0: x, 1:m etc
	sign  int // 0: inf 1: sup -1: none
	value int
	issue string
}

type rule struct {
	cmds []cmd
}

var workflow = map[string]rule{}

func evaluate(p part, r rule) bool {
	for _, c := range r.cmds {
		if c.sign == -1 {
			if c.issue == "A" {
				return true
			}
			if c.issue == "R" {
				return false
			}
			return evaluate(p, workflow[c.issue])
		}

		valueToCheck := 0
		switch c.v {
		case 1:
			valueToCheck = p.m
		case 2:
			valueToCheck = p.a
		case 3:
			valueToCheck = p.s
		default:
			valueToCheck = p.x
		}

		if c.sign == 0 {
			if valueToCheck < c.value {
				if c.issue == "A" {
					return true
				}
				if c.issue == "R" {
					return false
				}
				return evaluate(p, workflow[c.issue])
			}
		}
		if c.sign == 1 {
			if valueToCheck > c.value {
				if c.issue == "A" {
					return true
				}
				if c.issue == "R" {
					return false
				}
				return evaluate(p, workflow[c.issue])
			}
		}
	}

	return false
}

func copyRange(pRange partrange) partrange {
	newRange := partrange{}

	for _, v := range pRange.minmax {
		newRange.minmax = append(newRange.minmax, []int{v[0], v[1]})
	}

	return newRange
}

func evaluateRange(pRange partrange, r rule) int {
	for _, c := range r.cmds {

		if c.sign == -1 { // issue directly
			if c.issue == "R" {
				return 0
			}
			if c.issue == "A" {
				return pRange.count()
			}
			return evaluateRange(pRange, workflow[c.issue])
		}

		sign := c.sign
		value := c.value

		newRangeL := copyRange(pRange)
		newRangeR := copyRange(pRange)

		minToCheck := pRange.minmax[c.v][0]
		maxToCheck := pRange.minmax[c.v][1]

		if sign == 0 { // <
			if maxToCheck < value {
				if c.issue == "A" {
					return pRange.count()
				}
				return evaluateRange(pRange, workflow[c.issue])
			} else if minToCheck < value {
				newRangeL.minmax[c.v] = []int{minToCheck, value - 1}
				newRangeR.minmax[c.v] = []int{value, maxToCheck}

				returnVal := evaluateRange(newRangeR, r)
				if c.issue == "A" {
					return newRangeL.count() + returnVal
				}
				return evaluateRange(newRangeL, workflow[c.issue]) + returnVal
			}
		}
		if sign == 1 { // >
			if minToCheck > value {
				if c.issue == "A" {
					return pRange.count()
				}
				return evaluateRange(pRange, workflow[c.issue])
			} else if maxToCheck > value {
				newRangeL.minmax[c.v] = []int{minToCheck, value}
				newRangeR.minmax[c.v] = []int{value + 1, maxToCheck}

				returnVal := evaluateRange(newRangeL, r)
				if c.issue == "A" {
					return newRangeR.count() + returnVal
				}

				return evaluateRange(newRangeL, r) + evaluateRange(newRangeR, workflow[c.issue])
			}
		}
	}

	return 0
}

func SolvePart1(data []string) int {
	re := regexp.MustCompile("[0-9]+")
	parts := []part{}

	inputIsPart := false
	for _, line := range data {
		if len(line) == 0 {
			inputIsPart = true
			continue
		}

		if inputIsPart {
			values := re.FindAllString(line, -1)
			x, _ := strconv.Atoi(values[0])
			m, _ := strconv.Atoi(values[1])
			a, _ := strconv.Atoi(values[2])
			s, _ := strconv.Atoi(values[3])

			parts = append(parts, part{x, m, a, s})
		} else {
			nrule := rule{}

			ruleName := strings.Split(line, "{")[0]

			strcmds := strings.Split(strings.Split(line, "{")[1], "}")[0]

			for _, strcmd := range strings.Split(strcmds, ",") {
				ncmd := cmd{}

				commasplit := strings.Split(strcmd, ":")
				if len(commasplit) == 1 {
					ncmd.sign = -1
					ncmd.issue = strcmd
				} else {
					ncmd.v = strings.Index("xmas", string(commasplit[0][0]))
					ncmd.issue = commasplit[1]
					nv, _ := strconv.Atoi(re.FindAllString(commasplit[0], -1)[0])

					ncmd.value = nv

					if strings.Contains(commasplit[0], "<") {
						ncmd.sign = 0
					} else {
						ncmd.sign = 1
					}
				}

				nrule.cmds = append(nrule.cmds, ncmd)
			}

			workflow[ruleName] = nrule
		}
	}

	sum := 0
	for _, p := range parts {
		if evaluate(p, workflow["in"]) {
			sum += p.x + p.m + p.a + p.s
		}
	}

	return sum
}

func SolvePart2(data []string) int {
	pRange0 := partrange{
		[][]int{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}},
	}
	return evaluateRange(pRange0, workflow["in"])
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
