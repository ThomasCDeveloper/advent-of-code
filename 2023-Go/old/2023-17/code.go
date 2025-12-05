package main

// CMD: go run *.go

import (
	"container/heap"
	"fmt"
	"strconv"
)

type xydir struct {
	x     int
	y     int
	dx    int
	dy    int
	compt int
}

type Item struct {
	priority int
	state    xydir
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	y := len(*pq)
	item := x.(*Item)
	item.index = y
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	m := len(old)
	item := old[m-1]
	old[m-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : m-1]
	return item
}

/*func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}*/

func SolvePart1(data []string) int {
	tiles := [][]int{}
	for y, line := range data {
		intLine := []int{}

		for x := range line {
			val, _ := strconv.Atoi(string(data[y][x]))
			intLine = append(intLine, val)
		}
		tiles = append(tiles, intLine)
	}

	pq := make(PriorityQueue, 2)
	minScores := map[xydir]int{}

	minScores[xydir{
		x:     0,
		y:     0,
		dx:    1,
		dy:    0,
		compt: 0,
	}] = 0
	minScores[xydir{
		x:     0,
		y:     0,
		dx:    0,
		dy:    1,
		compt: 0,
	}] = 0

	pq[0] = &Item{
		state: xydir{
			x:     0,
			y:     0,
			dx:    1,
			dy:    0,
			compt: 0,
		},

		priority: 0,
		index:    0,
	}
	pq[1] = &Item{
		state: xydir{
			x:     0,
			y:     0,
			dx:    0,
			dy:    1,
			compt: 0,
		},

		priority: 0,
		index:    1,
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if minScores[item.state] < item.priority {
			continue
		}

		if item.state.y == len(tiles)-1 && item.state.x == len(tiles[0])-1 {
			return item.priority
		}

		for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			if ((item.state.compt == 3) && dir[0] == item.state.dx && dir[1] == item.state.dy) || (dir[0] == -item.state.dx && dir[1] == -item.state.dy) {
				continue
			}

			ni, nj := item.state.y+dir[1], item.state.x+dir[0]
			nextMoves := item.state.compt

			if !(dir[0] == item.state.dx && dir[1] == item.state.dy) {
				nextMoves = 1
			} else {
				nextMoves = nextMoves + 1
			}

			if ni < 0 || ni >= len(tiles) || nj < 0 || nj >= len(tiles[0]) {
				continue
			}

			nextState := xydir{y: ni, x: nj, compt: nextMoves, dx: dir[0], dy: dir[1]}
			nextHeatLoss := tiles[ni][nj]
			if _, ok := minScores[nextState]; ok && minScores[nextState] <= item.priority+nextHeatLoss {
				continue
			}

			minScores[nextState] = item.priority + nextHeatLoss
			heap.Push(&pq, &Item{priority: item.priority + nextHeatLoss, state: nextState})
		}
	}

	return 0
}

func SolvePart2(data []string) int {
	tiles := [][]int{}
	for y, line := range data {
		intLine := []int{}

		for x := range line {
			val, _ := strconv.Atoi(string(data[y][x]))
			intLine = append(intLine, val)
		}
		tiles = append(tiles, intLine)
	}

	pq := make(PriorityQueue, 2)
	minScores := map[xydir]int{}

	minScores[xydir{
		x:     0,
		y:     0,
		dx:    1,
		dy:    0,
		compt: 0,
	}] = 0
	minScores[xydir{
		x:     0,
		y:     0,
		dx:    0,
		dy:    1,
		compt: 0,
	}] = 0

	pq[0] = &Item{
		state: xydir{
			x:     0,
			y:     0,
			dx:    1,
			dy:    0,
			compt: 0,
		},

		priority: 0,
		index:    0,
	}
	pq[1] = &Item{
		state: xydir{
			x:     0,
			y:     0,
			dx:    0,
			dy:    1,
			compt: 0,
		},

		priority: 0,
		index:    1,
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if minScores[item.state] < item.priority {
			continue
		}

		if item.state.y == len(tiles)-1 && item.state.x == len(tiles[0])-1 && item.state.compt >= 4 {
			return item.priority
		}

		for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			if ((item.state.compt == 10) && dir[0] == item.state.dx && dir[1] == item.state.dy) || (dir[0] == -item.state.dx && dir[1] == -item.state.dy) {
				continue
			}

			ni, nj := item.state.y+dir[1], item.state.x+dir[0]
			nextMoves := item.state.compt

			if nextMoves < 4 {
				if dir[0] != item.state.dx || dir[1] != item.state.dy {
					continue
				}
				nextMoves += 1
			} else {
				if dir[0] != item.state.dx || dir[1] != item.state.dy {
					nextMoves = 1
				} else {
					nextMoves = nextMoves%10 + 1
				}
			}

			if ni < 0 || ni >= len(tiles) || nj < 0 || nj >= len(tiles[0]) {
				continue
			}

			nextState := xydir{y: ni, x: nj, compt: nextMoves, dx: dir[0], dy: dir[1]}
			nextHeatLoss := tiles[ni][nj]
			if _, ok := minScores[nextState]; ok && minScores[nextState] <= item.priority+nextHeatLoss {
				continue
			}

			minScores[nextState] = item.priority + nextHeatLoss
			heap.Push(&pq, &Item{priority: item.priority + nextHeatLoss, state: nextState})
		}
	}

	return 0
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
