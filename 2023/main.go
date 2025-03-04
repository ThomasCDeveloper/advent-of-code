package main

import (
	aoc01 "ThomasCDeveloper/advent-of-code/2023/01"
	aoc02 "ThomasCDeveloper/advent-of-code/2023/02"
	aoc03 "ThomasCDeveloper/advent-of-code/2023/03"
	"fmt"
)

var days = map[string]func(){
	"01": aoc01.Run,
	"02": aoc02.Run,
	"03": aoc03.Run,
}

func main() {
	Benchmark(RunAll, 50)
}

func RunAll() {
	for k, v := range days {
		fmt.Printf("Day %s:\n", k)
		v()
	}
}
