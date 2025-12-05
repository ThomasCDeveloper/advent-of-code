package main

// CMD: go run *.go

import (
	"fmt"
	"slices"
	"strings"
)

type co struct {
	r, c int
}

/*
	type Node struct {
		n []co
	}

	func canWalk1(rc co, data []string) map[co]int {
		out := map[co]int{}

		char := rune(data[rc.r][rc.c])

		if char == '.' {
			if strings.Contains(".v>", string(data[rc.r+0][rc.c+1])) {
				out[co{rc.r + 0, rc.c + 1}] = 1
			}
			if strings.Contains(".v", string(data[rc.r+0][rc.c-1])) {
				out[co{rc.r + 0, rc.c - 1}] = 1
			}
			if strings.Contains(".v>", string(data[rc.r+1][rc.c+0])) {
				out[co{rc.r + 1, rc.c + 0}] = 1
			}
			if strings.Contains(".>", string(data[rc.r-1][rc.c+0])) {
				out[co{rc.r - 1, rc.c + 0}] = 1
			}
		}
		if char == '>' {
			if strings.Contains(".v>", string(data[rc.r+0][rc.c+1])) {
				out[co{rc.r + 0, rc.c + 1}] = 1
			}
		}
		if char == 'v' {
			if strings.Contains(".v>", string(data[rc.r+1][rc.c+0])) {
				out[co{rc.r + 1, rc.c + 0}] = 1
			}
		}

		return out
	}

	func Walk(rc co, nodes map[co]Node, max map[co]int, visited []co) {
		visited = append(visited, rc)

		for _, next := range nodes[rc].n {
			if !slices.Contains(visited, next) {
				val := max[next]

				if val < max[rc]+1 {
					max[next] = max[rc] + 1
					Walk(next, nodes, max, visited)
				}
			}
		}
	}

func SolvePart1(data []string) int {

		nodes := map[co]Node{}

		for r := 1; r < len(data)-1; r++ {
			for c := 1; c < len(data[0])-1; c++ {
				if data[r][c] != '#' {
					nodes[co{r, c}] = Node{canWalk1(co{r, c}, data)}
				}
			}
		}

		startCo := co{1, 1}
		max := map[co]int{startCo: 1}

		Walk(startCo, nodes, max, []co{})

		return max[co{len(data) - 1, len(data[0]) - 2}]
	}
*/
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

var seen = []co{}
var graph = map[co]map[co]int{}
var end = co{}

func dfs(rc co) int {
	if rc.r == end.r && rc.c == end.c {
		return 0
	}

	m := 0

	if !slices.Contains(seen, rc) {
		seen = append(seen, rc)
	}

	for nx := range graph[rc] {
		if !slices.Contains(seen, nx) {
			value := dfs(nx) + graph[rc][nx]
			m = max(m, value)
		}
	}

	seen = seen[:len(seen)-1]

	return m
}

type item struct {
	n, r, c int
}

func SolvePart2(data []string) int {
	start := co{0, strings.Index(data[0], ".")}
	end = co{len(data) - 1, strings.Index(data[len(data)-1], ".")}

	points := []co{start, end}

	for r, row := range data {
		for c, char := range row {
			if char == '#' {
				continue
			}

			neighbors := 0
			for _, dco := range []co{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nr, nc := r+dco.r, c+dco.c
				if nr >= 0 && nr < len(data) && nc >= 0 && nc < len(data[0]) {
					if data[nr][nc] != '#' {
						neighbors++
					}
				}
			}

			if neighbors >= 3 {
				points = append(points, co{r, c})
			}
		}
	}

	for _, pt := range points {
		graph[pt] = map[co]int{}
	}

	for _, pt := range points {
		sr := pt.r
		sc := pt.c

		stack := []item{{0, sr, sc}}
		seen = []co{{sr, sc}}

		for len(stack) > 0 {
			n, r, c := stack[0].n, stack[0].r, stack[0].c

			stack = stack[1:]

			if n != 0 && slices.Contains(points, co{r, c}) {
				graph[co{sr, sc}][co{r, c}] = n
				continue
			}

			for _, dco := range []co{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nr, nc := r+dco.r, c+dco.c
				if nr >= 0 && nr < len(data) && nc >= 0 && nc < len(data[0]) {
					if data[nr][nc] != '#' {
						if !slices.Contains(seen, co{nr, nc}) {
							stack = append(stack, item{n + 1, nr, nc})
							seen = append(seen, co{nr, nc})
						}
					}
				}
			}
		}
	}

	seen = []co{}

	return dfs(start)
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	//fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
