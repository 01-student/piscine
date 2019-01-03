package main

import "github.com/01-edu/z01"

func main() {
	r := '0'
	for r <= '9' {
		z01.PrintRune(r)
		r++
	}
	z01.PrintRune('\n')
}
