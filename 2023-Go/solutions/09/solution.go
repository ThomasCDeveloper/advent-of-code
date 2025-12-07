package main

import (
	"ThomasCDeveloper/advent-of-code/2023/utils"
	"fmt"
	"os"
	"strings"
)

func getReduction(nums []int) []int {
	out := []int{}
	for i := range nums[1:] {
		out = append(out, nums[i+1]-nums[i])
	}
	return out
}

func checkIfAllZeros(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

func part1(data []string) int {
	total := 0
	for _, l := range data {
		nums := []int{}
		for _, s := range strings.Split(l, " ") {
			nums = append(nums, utils.Atoi(s))
		}

		lastValues := []int{}
		for !checkIfAllZeros(nums) {
			lastValues = append(lastValues, nums[len(nums)-1])
			nums = getReduction(nums)
		}

		for _, n := range lastValues {
			total += n
		}
	}
	return total
}

func part2(data []string) int {
	total := 0
	for _, l := range data {
		nums := []int{}
		for _, s := range strings.Split(l, " ") {
			nums = append(nums, utils.Atoi(s))
		}

		firstValues := []int{}
		for !checkIfAllZeros(nums) {
			firstValues = append(firstValues, nums[0])
			nums = getReduction(nums)
		}

		for i, n := range firstValues {
			if i%2 == 1 {
				total -= n
			} else {
				total += n
			}
		}
	}
	return total
}

func main() {
	path := utils.GetPath(os.Args)
	data := utils.ParseInputLines(path)

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
