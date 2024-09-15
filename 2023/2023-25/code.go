package main

// CMD: go run *.go

import (
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"strings"
)

type link struct {
	a, b string
}

func GetPath(source string, sink string, currentPath []string, steps int) ([]string, int) {
	cache[source] = steps

	if source == sink {
		return currentPath, steps
	}

	path := []string{}
	min := 99999999999

	for _, n := range nodes[source].neighbors {
		if val, ok := cache[n]; val > steps+1 || !ok {
			cPath, v := GetPath(n, sink, append(currentPath, n), steps+1)

			if v < min {
				min = v
				path = cPath
			}
		}
	}

	return path, min
}

type Node struct {
	neighbors []string
}

var nodes = map[string]Node{}
var cache = map[string]int{}

type passages struct {
	l link
	n int
}

var loopCache = []string{}

func FillLoopCache(n0 string) {
	if !slices.Contains(loopCache, n0) {
		loopCache = append(loopCache, n0)

		for _, n1 := range nodes[n0].neighbors {
			FillLoopCache(n1)
		}
	}
}

func SolvePart1(data []string) int {
	links := []link{}
	nodeNames := []string{}

	for _, line := range data {
		a := strings.Split(line, ": ")[0]
		if !slices.Contains(nodeNames, a) {
			nodeNames = append(nodeNames, a)
		}
		for _, b := range strings.Split(strings.Split(line, ": ")[1], " ") {
			links = append(links, link{a, b})
			if !slices.Contains(nodeNames, b) {
				nodeNames = append(nodeNames, b)
			}
		}
	}

	for _, link := range links {
		if _, ok := nodes[link.a]; !ok {
			nodes[link.a] = Node{[]string{}}
		}
		if _, ok := nodes[link.b]; !ok {
			nodes[link.b] = Node{[]string{}}
		}

		if !slices.Contains(nodes[link.a].neighbors, link.b) {
			nodes[link.a] = Node{append(nodes[link.a].neighbors, link.b)}
		}
		if !slices.Contains(nodes[link.b].neighbors, link.a) {
			nodes[link.b] = Node{append(nodes[link.b].neighbors, link.a)}
		}
	}

	counts := map[link]int{}

	iter := 100
	for i := 0; i < iter; i++ {

		cache = map[string]int{}

		source := nodeNames[rand.Intn(len(nodeNames))]
		sink := source
		for sink == source {
			sink = nodeNames[rand.Intn(len(nodeNames))]
		}

		p, _ := GetPath(source, sink, []string{source}, 0)

		for j := range p[1:] {
			counts[link{p[j], p[j+1]}] += 1
		}
	}

	pass := []passages{}

	for key, val := range counts {
		pass = append(pass, passages{key, val})
	}

	sort.Slice(pass, func(a, b int) bool {
		return pass[a].n > pass[b].n
	})

	fmt.Println(pass[:10])

	cuts := []link{pass[0].l, pass[1].l}

	fmt.Println(pass[:10])

	nodes = map[string]Node{}
	links = []link{}

	for _, line := range data {
		a := strings.Split(line, ": ")[0]
		for _, b := range strings.Split(strings.Split(line, ": ")[1], " ") {
			if !(slices.Contains(cuts, link{a, b}) || slices.Contains(cuts, link{b, a})) {
				links = append(links, link{a, b})
			}
		}
	}

	for _, link := range links {
		if _, ok := nodes[link.a]; !ok {
			nodes[link.a] = Node{[]string{}}
		}
		if _, ok := nodes[link.b]; !ok {
			nodes[link.b] = Node{[]string{}}
		}

		if !slices.Contains(nodes[link.a].neighbors, link.b) {
			nodes[link.a] = Node{append(nodes[link.a].neighbors, link.b)}
		}
		if !slices.Contains(nodes[link.b].neighbors, link.a) {
			nodes[link.b] = Node{append(nodes[link.b].neighbors, link.a)}
		}
	}

	counts = map[link]int{}
	for i := 0; i < iter; i++ {

		fmt.Println("%", float64(i+1+iter)*100/float64(2*iter))

		cache = map[string]int{}

		source := nodeNames[rand.Intn(len(nodeNames))]
		sink := source
		for sink == source {
			sink = nodeNames[rand.Intn(len(nodeNames))]
		}

		p, _ := GetPath(source, sink, []string{source}, 0)

		for j := range p[1:] {
			counts[link{p[j], p[j+1]}] += 1
		}
	}

	pass = []passages{}

	for key, val := range counts {
		pass = append(pass, passages{key, val})
	}

	sort.Slice(pass, func(a, b int) bool {
		return pass[a].n > pass[b].n
	})

	fmt.Println(pass[:10])
	cuts = append(cuts, pass[0].l)

	nodes = map[string]Node{}
	links = []link{}

	for _, line := range data {
		a := strings.Split(line, ": ")[0]
		for _, b := range strings.Split(strings.Split(line, ": ")[1], " ") {
			if !(slices.Contains(cuts, link{a, b}) || slices.Contains(cuts, link{b, a})) {
				links = append(links, link{a, b})
			}
		}
	}

	for _, link := range links {
		if _, ok := nodes[link.a]; !ok {
			nodes[link.a] = Node{[]string{}}
		}
		if _, ok := nodes[link.b]; !ok {
			nodes[link.b] = Node{[]string{}}
		}

		if !slices.Contains(nodes[link.a].neighbors, link.b) {
			nodes[link.a] = Node{append(nodes[link.a].neighbors, link.b)}
		}
		if !slices.Contains(nodes[link.b].neighbors, link.a) {
			nodes[link.b] = Node{append(nodes[link.b].neighbors, link.a)}
		}
	}

	FillLoopCache(nodeNames[0])

	return len(loopCache) * (len(nodeNames) - len(loopCache))
}

func main() {
	data := GetInput("input.txt")

	// PART F
	fmt.Println("Part F:", SolvePart1(data))
}
