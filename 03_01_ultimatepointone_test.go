package piscine_test

import (
	"testing"

	piscine "."
)

func TestUltimatePointOne(t *testing.T) {
	n := 0
	nn := &n
	nnn := &nn
	piscine.UltimatePointOne(&nnn)
	if n != 1 {
		t.Errorf("UltimatePointOne(&nnn), n == %d instead of 1", n)
	}
}
