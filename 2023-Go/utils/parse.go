package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func GetPath(args []string) string {
	if len(os.Args) < 3 {
		panic("missing day argument")
	}
	day := args[2]
	filename := "input.txt"
	if args[1] == "test" {
		filename = "test.txt"
	}
	path := filepath.Join("solutions", day, filename)
	return path
}

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
