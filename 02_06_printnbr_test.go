package piscine_test

import (
	"fmt"
	"strconv"
	"testing"

	piscine "."
	"github.com/01-edu/z01/test"
)

func printNbr(n int) {
	fmt.Print(strconv.Itoa(n))
}

func TestPrintNbr(t *testing.T) {
	table := test.RandomInts(30)
	table = append(table,
		test.IntMin,
		test.IntMax,
		0,
	)
	for _, arg := range table {
		test.Challenge(t, piscine.PrintNbr, printNbr, arg)
	}
}
