package main

// CMD: go run *.go

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type sequence struct {
	springs string
	cmd     string
}

var cache = make(map[sequence]int)

func GetCombinaisons(cmd []uint8, springs string) int {
	if springs == "" {
		if len(cmd) == 0 { // fin du bloc
			return 1
		} else {
			return 0
		}
	}

	if len(cmd) == 0 {
		if strings.Contains(springs, "#") { // plus de commande mais au moins un spring restant
			return 0
		} else {
			return 1
		}
	}

	key := sequence{springs, string(cmd)}
	cachedValue, isCached := cache[key]
	if isCached {
		return cachedValue
	}

	sum := 0

	firstSpring := string(springs[0])

	if strings.Contains(".?", firstSpring) {
		// 1er spring traité comme .
		sum += GetCombinaisons(cmd, springs[1:])
	}

	if strings.Contains("#?", firstSpring) {
		// 1er spring traité comme #
		firstNum := int(cmd[0])
		if firstNum <= len(springs) && !strings.Contains(springs[:firstNum], ".") && (firstNum == len(springs) || string(springs[firstNum]) != "#") {
			min := int(math.Min(float64(firstNum+1), float64(len(springs))))
			sum += GetCombinaisons(cmd[1:], springs[min:])
		}
	}

	cache[key] = sum

	return sum
}

func SolvePart1(data []string) int {
	cmds := [][]uint8{}
	springss := []string{}
	for _, line := range data {
		re := regexp.MustCompile("[0-9]+")

		cmd := []uint8{}
		for _, stringcmd := range re.FindAllString(line, -1) {
			intcmd, _ := strconv.Atoi(stringcmd)
			int8cmd := uint8(intcmd)
			cmd = append(cmd, int8cmd)
		}

		cmds = append(cmds, cmd)
		springss = append(springss, strings.Split(line, " ")[0])
	}

	sum := 0
	for i := range springss {
		sum += GetCombinaisons(cmds[i], springss[i])
	}

	return sum
}

func SolvePart2(data []string) int {
	cmds := [][]uint8{}
	springss := []string{}
	for _, line := range data {
		re := regexp.MustCompile("[0-9]+")

		cmd := []uint8{}
		for _, stringcmd := range re.FindAllString(line, -1) {
			intcmd, _ := strconv.Atoi(stringcmd)
			int8cmd := uint8(intcmd)
			cmd = append(cmd, int8cmd)
		}

		cmds = append(cmds, cmd)
		springss = append(springss, strings.Split(line, " ")[0])
	}

	sum := 0
	for i := range springss {
		cmd := []uint8{}
		springs := ""
		for j := 0; j < 5; j++ {
			cmd = append(cmd, cmds[i]...)
			springs += "?" + springss[i]
		}
		sum += GetCombinaisons(cmd, springs[1:])
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
