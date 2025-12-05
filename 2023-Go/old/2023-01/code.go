package main

// CMD: go run *.go

import (
	"fmt"
	"strconv"
	"strings"
)

func getFirstDigit(word string) string {
	digits := "0123456789"

	runes := []rune(word)

	for i := 0; i < len(word); i++ {
		char := string(runes[i])
		if strings.Index(digits, char) != -1 {
			return char
		}
	}

	return "0"
}

func getLastDigit(word string) string {
	digits := "0123456789"

	runes := []rune(word)

	for i := 0; i < len(word); i++ {
		char := string(runes[len(word)-i-1])
		if strings.Index(digits, char) != -1 {
			return char
		}
	}

	return "0"
}

func getFirstDigit2(word string) string {
	numbers := map[string]int{
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	numberPos := [10]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	for key, val := range numbers {
		position := strings.Index(word, key)
		if position != -1 {
			if position < numberPos[val] || numberPos[val] == -1 {
				numberPos[val] = position
			}
		}
	}

	min, val := len(word)+10, 0
	for i := 0; i < 10; i++ {
		if numberPos[i] != -1 {
			if numberPos[i] < min {
				min = numberPos[i]
				val = i
			}
		}
	}

	return strconv.Itoa(val)
}

func getLastDigit2(word string) string {
	numbers := map[string]int{
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	numberPos := [10]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	for key, val := range numbers {
		position := strings.LastIndex(word, key)
		if position != -1 {
			if position > numberPos[val] || numberPos[val] == -1 {
				numberPos[val] = position
			}
		}
	}

	max, val := -1, 0
	for i := 0; i < 10; i++ {
		if numberPos[i] != -1 {
			if numberPos[i] > max {
				max = numberPos[i]
				val = i
			}
		}
	}

	return strconv.Itoa(val)
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:")

	compt := 0
	for _, word := range data {
		value, _ := strconv.Atoi(getFirstDigit(word) + getLastDigit(word))
		compt += value
	}
	fmt.Println(compt)

	// PART 2
	fmt.Println("Part 2:")

	compt = 0
	for _, word := range data {
		value, _ := strconv.Atoi(getFirstDigit2(word) + getLastDigit2(word))
		compt += value
	}
	fmt.Println(compt)
}
