package main

// CMD: go run *.go

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getNumbersFromList(nums string) []int {
	output := []int{}

	i := 0
	for i < len(nums) {
		if nums[i] == ' ' {
			num, _ := strconv.Atoi(string(nums[i+1]))
			output = append(output, num)
		} else {
			num, _ := strconv.Atoi(string(nums[i : i+2]))
			output = append(output, num)
		}
		i = i + 3
	}

	return output
}

func SolvePart1(data []string) int {
	result := 0
	for _, val := range data {
		sum := 0
		cardNumbers := getNumbersFromList(strings.Split(strings.Split(val, ": ")[1], "| ")[0])
		pickedNumbers := getNumbersFromList(strings.Split(strings.Split(val, ": ")[1], "| ")[1])

		for _, pickedNumber := range pickedNumbers {
			for _, cardNumber := range cardNumbers {
				if cardNumber == pickedNumber {
					if sum == 0 {
						sum += 1
					} else {
						sum *= 2
					}
				}
			}
		}

		result += sum
	}

	return result
}

func SolvePart2(data []string) int {
	result := 0

	numberOfCards := make([]int, len(data))
	for i := 0; i < len(data); i++ {
		numberOfCards[i] = 1
	}

	for i, val := range data {
		cardNumbers := getNumbersFromList(strings.Split(strings.Split(val, ": ")[1], "| ")[0])
		pickedNumbers := getNumbersFromList(strings.Split(strings.Split(val, ": ")[1], "| ")[1])

		numberOfMatches := 0

		for _, pickedNumber := range pickedNumbers {
			for _, cardNumber := range cardNumbers {
				if cardNumber == pickedNumber {
					numberOfMatches += 1
				}
			}
		}

		for j := i + 1; j < int(math.Min(float64(len(data)), float64(i+numberOfMatches+1))); j++ {
			numberOfCards[j] += numberOfCards[i]
		}

	}

	for _, val := range numberOfCards {
		result += val
	}

	return result
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:")
	fmt.Println(SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:")
	fmt.Println(SolvePart2(data))
}
