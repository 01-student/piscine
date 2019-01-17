package piscine_test

import (
	"fmt"
	"testing"

	piscine "."
	"github.com/01-edu/z01/test"
)

func isNegative(n int) {
	if n < 0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}

func TestIsNegative(t *testing.T) {
	table := test.RandomInts(30)
	table = append(table,
		test.IntMin,
		test.IntMax,
		0,
	)
	for _, arg := range table {
		test.Challenge(t, piscine.IsNegative, isNegative, arg)
	}
}
