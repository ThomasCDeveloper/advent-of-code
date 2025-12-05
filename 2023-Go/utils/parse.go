package utils

import (
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseRawInput(path string) string {
	dat, err := os.ReadFile(path)
	Check(err)
	return string(dat)
}

func ParseInputLines(path string) []string {
	raw := ParseRawInput(path)
	return strings.Split(raw, "\n")
}
