package main

import "github.com/01-edu/z01"

func main() {
	r := 'a'
	for r <= 'z' {
		z01.PrintRune(r)
		r++
	}
	z01.PrintRune('\n')
}
