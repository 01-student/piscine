package piscine_test

import (
	"testing"

	piscine "."
)

func TestPointOne(t *testing.T) {
	n := 0
	piscine.PointOne(&n)
	if n != 1 {
		t.Errorf("PointOne(&n), n == %d instead of 1", n)
	}
}
