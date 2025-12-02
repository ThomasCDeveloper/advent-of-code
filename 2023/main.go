package main

import (
	aoc01 "ThomasCDeveloper/advent-of-code/2023/01"
	aoc02 "ThomasCDeveloper/advent-of-code/2023/02"
	aoc03 "ThomasCDeveloper/advent-of-code/2023/03"
	aoc04 "ThomasCDeveloper/advent-of-code/2023/04"
	aoc05 "ThomasCDeveloper/advent-of-code/2023/05"
	"fmt"
)

var days = map[string]func(){
	"01": aoc01.Run,
	"02": aoc02.Run,
	"03": aoc03.Run,
	"04": aoc04.Run,
	"05": aoc05.Run,
}

func main() {
	days["05"]()
	//Benchmark(RunAll, 50)
}

func RunAll() {
	for k, v := range days {
		fmt.Printf("Day %s:\n", k)
		v()
	}
}
