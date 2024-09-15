package main

// CMD: go run *.go

import (
	"fmt"
	"slices"
	"strings"
)

type module struct {
	name       string
	isFlipFlop bool

	outputs []string

	isHigh   bool
	memories map[string]bool
}

type pulse struct {
	sender   string
	receiver string
	high     bool
}

var pulsesQueue = []pulse{}
var mods = map[string]module{}

func addPulseToQueue(p pulse) {
	pulsesQueue = append(pulsesQueue, p)
	/*fmt.Printf(p.sender)
	if p.high {
		fmt.Printf(" -high-> ")
	} else {
		fmt.Printf(" -low-> ")
	}
	fmt.Printf(p.receiver + "\n")*/
}

func processPulse(p pulse) {
	m := mods[p.receiver]
	sig := p.high
	sen := p.sender
	if m.isFlipFlop {
		if !sig {
			m.isHigh = !m.isHigh
			for _, output := range m.outputs {
				addPulseToQueue(pulse{m.name, output, m.isHigh})
			}
		}
	} else {
		m.memories[sen] = sig
		allHigh := true
		for _, mem := range m.memories {
			allHigh = allHigh && mem
		}
		for _, output := range m.outputs {
			addPulseToQueue(pulse{m.name, output, !allHigh})
		}
	}
	mods[p.receiver] = m

	if p.high {
		nbHigh++
	} else {
		nbLow++
	}
}

func pushButton() {
	nbLow++
	for _, receiver := range starters {
		addPulseToQueue(pulse{"broadcaster", receiver, false})
	}
}

func getMods(data []string) {
	mods = map[string]module{}
	starters = []string{}
	for _, line := range data {
		nameOutput := strings.Split(line, " -> ")
		modName := nameOutput[0]

		if modName == "broadcaster" {
			starters = append(starters, strings.Split(nameOutput[1], ", ")...)
		} else {
			isFlipFlop := false
			if modName[0] == '%' {
				isFlipFlop = true
			}
			if isFlipFlop {
				mods[modName[1:]] = module{modName[1:], true, strings.Split(nameOutput[1], ", "), false, map[string]bool{}}
			} else {
				mods[modName[1:]] = module{modName[1:], false, strings.Split(nameOutput[1], ", "), false, map[string]bool{}}
			}
		}

	}

	for _, mod := range mods {
		for _, receiver := range mod.outputs {
			if _, ok := mods[receiver]; !ok {
				mods[receiver] = module{receiver, false, []string{}, false, map[string]bool{}}
			}
			mods[receiver].memories[mod.name] = false
		}

		if slices.Contains(mod.outputs, "nc") { // last before "rx"
			loopsSize[mod.name] = 0
		}
	}
}

var (
	nbHigh   = 0
	nbLow    = 0
	starters = []string{}
)

func SolvePart1(data []string) int {
	getMods(data)

	for i := 0; i < 1000; i++ {
		pushButton()

		for len(pulsesQueue) > 0 {
			processPulse(pulsesQueue[0])
			pulsesQueue = pulsesQueue[1:]
		}
	}

	return nbHigh * nbLow
}

func product(l map[string]int) int {
	v := 1
	for _, val := range l {
		v *= val
	}
	return v
}

var loopsSize = map[string]int{}

func SolvePart2(data []string) int {
	getMods(data)

	i := 0
	for product(loopsSize) == 0 {
		pushButton()
		i++

		for len(pulsesQueue) > 0 {
			if slices.Contains([]string{"hh", "fh", "fn", "lk"}, pulsesQueue[0].receiver) && !pulsesQueue[0].high {
				loopsSize[pulsesQueue[0].receiver] = i
			}
			processPulse(pulsesQueue[0])
			pulsesQueue = pulsesQueue[1:]
		}
	}
	return product(loopsSize)
}

func main() {
	data := GetInput("input.txt")

	// PART 1
	fmt.Println("Part 1:", SolvePart1(data))

	// PART 2
	fmt.Println("Part 2:", SolvePart2(data))
}
