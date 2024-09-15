package main

// CMD: go run *.go

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Bet struct {
	cards    []int
	bet      int
	wins     int
	handType int
	raw      string
}

func Compare(bet1, bet2 Bet) bool {
	if bet1.handType > bet2.handType {
		return true
	}
	if bet2.handType > bet1.handType {
		return false
	}

	for i := range bet1.cards {
		if bet1.cards[i] > bet2.cards[i] {
			return true
		}
		if bet2.cards[i] > bet1.cards[i] {
			return false
		}
	}

	return true
}

func GetBet(line string, part2 bool) Bet {
	bet := Bet{}

	cards := make([]int, 5)
	values := "23456789TJQKA"
	if part2 {
		values = "J23456789TQKA"
	}
	rawDeck := strings.Split(line, " ")[0]
	for i := 0; i < 5; i++ {
		cards[i] = strings.Index(values, string(rawDeck[i]))
	}

	repartition := GetCardRepartition(cards)

	Js := 0
	if part2 {
		Js = repartition[0]
		repartition = repartition[1:]
	}

	sort.Slice(repartition, func(i, j int) bool {
		return repartition[i] > repartition[j]
	})
	if part2 {
		repartition[0] += Js
	}

	handType := 0
	if repartition[0] == 5 {
		handType = 6
	} else if repartition[0] == 4 {
		handType = 5
	} else if repartition[0] == 3 && repartition[1] == 2 {
		handType = 4
	} else if repartition[0] == 3 {
		handType = 3
	} else if repartition[0] == 2 && repartition[1] == 2 {
		handType = 2
	} else if repartition[0] == 2 {
		handType = 1
	}

	bet.cards = cards
	bet.bet, _ = strconv.Atoi(strings.Split(line, " ")[1])
	bet.handType = handType
	bet.wins = 0
	bet.raw = rawDeck

	return bet
}

func GetCardRepartition(cards []int) []int {
	output := make([]int, 13)
	for i := 0; i < 5; i++ {
		output[cards[i]]++
	}
	return output
}

func SolvePart1(data []string) int {
	tournament := []Bet{}
	for _, line := range data {
		tournament = append(tournament, GetBet(line, false))
	}

	for i := 0; i < len(tournament)-1; i++ {
		for j := i + 1; j < len(tournament); j++ {
			if Compare(tournament[i], tournament[j]) {
				tournament[i].wins++
			} else {
				tournament[j].wins++
			}
		}
	}

	sort.Slice(tournament, func(i, j int) bool {
		return tournament[i].wins < tournament[j].wins
	})

	sum := 0

	for i := range tournament {
		sum += tournament[i].bet * (i + 1)
	}

	return sum
}

func SolvePart2(data []string) int {
	tournament := []Bet{}
	for _, line := range data {
		tournament = append(tournament, GetBet(line, true))
	}

	for i := 0; i < len(tournament)-1; i++ {
		for j := i + 1; j < len(tournament); j++ {
			if Compare(tournament[i], tournament[j]) {
				tournament[i].wins++
			} else {
				tournament[j].wins++
			}
		}
	}

	sort.Slice(tournament, func(i, j int) bool {
		return tournament[i].wins < tournament[j].wins
	})

	sum := 0

	for i := range tournament {
		sum += tournament[i].bet * (i + 1)
	}

	return sum
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
