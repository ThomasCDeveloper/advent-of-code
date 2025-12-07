package main

import (
	"ThomasCDeveloper/advent-of-code/2023/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	from int
	to   int
	n    int
}

type ConvertMap struct {
	ranges []Range
}

func processSeed(seed int, ma ConvertMap) int {
	for _, ran := range ma.ranges {
		if seed >= ran.from && seed <= ran.from+ran.n {
			return ran.to + seed - ran.from
		}
	}

	return seed
}

func part1(data []string) int {
	seeds := []int{}
	for _, val := range strings.Split(data[0], " ")[1:len(strings.Split(data[0], " "))] {
		convert, _ := strconv.Atoi(val)
		seeds = append(seeds, convert)
	}

	maps := []ConvertMap{}
	newMap := ConvertMap{}
	for _, line := range data[2:] {
		if line == "" {
			continue
		}

		if strings.Contains("abcdefghijklmnopqrstuvwxyz", string(line[0])) {
			if len(newMap.ranges) != 0 {
				maps = append(maps, newMap)
				newMap = ConvertMap{}
			}
			continue
		}

		to, _ := strconv.Atoi(strings.Split(line, " ")[0])
		from, _ := strconv.Atoi(strings.Split(line, " ")[1])
		n, _ := strconv.Atoi(strings.Split(line, " ")[2])

		newMap.ranges = append(newMap.ranges, Range{from, to, n})
	}

	maps = append(maps, newMap)

	results := []int{}

	for _, seed := range seeds {
		for _, ma := range maps {
			seed = processSeed(seed, ma)
		}
		results = append(results, seed)
	}

	min := results[0]
	for _, res := range results {
		if res < min {
			min = res
		}
	}

	return min
}

func applyMapToRanges(ranges [][2]int, m ConvertMap) [][2]int {
	out := make([][2]int, 0, len(ranges))

	for _, r := range ranges {
		start, end := r[0], r[1]

		remaining := [][2]int{{start, end}}

		for _, tr := range m.ranges {
			newRemaining := make([][2]int, 0, len(remaining))
			mapStart := tr.from
			mapEnd := tr.from + tr.n

			for _, seg := range remaining {
				s, e := seg[0], seg[1]

				if e < mapStart || s > mapEnd {
					newRemaining = append(newRemaining, seg)
					continue
				}

				if s < mapStart {
					newRemaining = append(newRemaining, [2]int{s, mapStart - 1})
					s = mapStart
				}

				if e > mapEnd {
					newRemaining = append(newRemaining, [2]int{mapEnd + 1, e})
					e = mapEnd
				}

				offset := tr.to - tr.from
				out = append(out, [2]int{s + offset, e + offset})
			}

			remaining = newRemaining
		}

		out = append(out, remaining...)
	}

	return out
}

func part2(data []string) int {
	maps := []ConvertMap{}
	newMap := ConvertMap{}

	for _, line := range data[2:] {
		if line == "" {
			continue
		}
		if line[0] >= 'a' && line[0] <= 'z' {
			if len(newMap.ranges) > 0 {
				maps = append(maps, newMap)
				newMap = ConvertMap{}
			}
			continue
		}

		f := strings.Fields(line)
		to, _ := strconv.Atoi(f[0])
		from, _ := strconv.Atoi(f[1])
		n, _ := strconv.Atoi(f[2])
		newMap.ranges = append(newMap.ranges, Range{from, to, n})
	}
	maps = append(maps, newMap)

	parts := strings.Fields(data[0])[1:]
	seedRanges := make([][2]int, 0, len(parts)/2)

	for i := 0; i < len(parts); i += 2 {
		start, _ := strconv.Atoi(parts[i])
		n, _ := strconv.Atoi(parts[i+1])
		seedRanges = append(seedRanges, [2]int{start, start + n})
	}

	for _, m := range maps {
		seedRanges = applyMapToRanges(seedRanges, m)
	}

	min := seedRanges[0][0]
	for _, r := range seedRanges {
		if r[0] < min {
			min = r[0]
		}
	}
	return min
}

func main() {
	path := utils.GetPath(os.Args)
	data := utils.ParseInputLines(path)

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
