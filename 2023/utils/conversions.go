package utils

import "strconv"

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}
