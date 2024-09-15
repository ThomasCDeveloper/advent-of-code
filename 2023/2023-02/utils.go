package main

import (
	"bufio"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetInput(fileName string) []string {
	output := []string{}

	f, err := os.Open(fileName)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}
