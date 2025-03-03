package main

// CMD: go run *.go

import (
	"fmt"
	"strconv"
	"strings"
)

func getHash(word string) int {
	val := 0

	for _, letter := range word {
		val += int(letter)
		val *= 17
		val %= 256
	}

	return val
}

func decomposeWord(word string) (int, string, string, int) {
	separator := " "
	if strings.Contains(word, "=") {
		separator = "="
	} else {
		separator = "-"
	}

	parts := strings.Split(word, separator)
	label := getHash(parts[0])
	lens := 0
	if len(parts[1]) != 0 {
		lens, _ = strconv.Atoi(parts[1])
	}

	return label, parts[0], separator, lens
}

func SolvePart1(data []string) int {
	sum := 0

	for _, word := range data {
		sum += getHash(word)
	}
	return sum
}

type Label struct {
	label string
	lens  int
}

func SolvePart2(data []string) int {
	boxes := make(map[int][]Label)
	for _, word := range data {
		boxNumber, boxlabel, separator, lens := decomposeWord(word)
		if _, test := boxes[boxNumber]; !test {
			boxes[boxNumber] = []Label{}
		}
		if separator == "=" {
			hasLabel := false
			for i, lab := range boxes[boxNumber] {
				if lab.label == boxlabel {
					lab.lens = lens
					boxes[boxNumber][i] = lab
					hasLabel = true
				}
			}
			if !hasLabel {
				boxes[boxNumber] = append(boxes[boxNumber], Label{boxlabel, lens})
			}
		} else {
			for i, lab := range boxes[boxNumber] {
				if lab.label == boxlabel {
					bef := boxes[boxNumber][:i]
					aft := boxes[boxNumber][i+1:]
					newLine := append([]Label{}, bef...)
					newLine = append(newLine, aft...)

					boxes[boxNumber] = newLine
				}
			}
		}
	}

	sum := 0
	for box := range boxes {
		for lab := range boxes[box] {
			sum += (box + 1) * (lab + 1) * boxes[box][lab].lens
		}
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
