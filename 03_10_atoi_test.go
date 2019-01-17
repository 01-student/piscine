package piscine_test

import (
	"strconv"
	"testing"

	piscine "."
	"github.com/01-edu/z01/test"
)

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func TestAtoi(t *testing.T) {
	table := make([]string, 30)
	for i := range table {
		table[i] = strconv.Itoa(test.RandomInt())
	}
	table = append(table,
		strconv.Itoa(test.IntMin),
		strconv.Itoa(test.IntMax),
		"",
		"-",
		"+",
		"0",
		"+0",
		"-Invalid123",
		"--123",
		"-+123",
		"++123",
		"123-",
		"123+",
		"123.",
		"123.0",
	)
	for _, arg := range table {
		test.Challenge(t, piscine.Atoi, atoi, arg)
	}
}
