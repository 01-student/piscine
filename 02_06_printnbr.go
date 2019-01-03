package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -1<<63 {
		z01.PrintRune('-')
		printDigit(9)
		n = 223372036854775808
	}
	if n < 0 {
		n = n * -1
		z01.PrintRune('-')
	}
	if n >= 10 {
		PrintNbr(n / 10)
	}
	printDigit(int(n % 10))
}
