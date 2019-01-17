package piscine_test

import (
	"testing"
	"unicode/utf8"

	piscine "."
	"github.com/01-edu/z01/test"
)

func strLen(s string) int {
	return utf8.RuneCountInString(s)
}

func TestStrLen(t *testing.T) {
	table := test.RandomASCIIStrings(30)
	table = append(table, "")
	for _, arg := range table {
		test.Challenge(t, piscine.StrLen, strLen, arg)
	}
}
