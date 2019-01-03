package piscine

import "github.com/01-edu/z01"

func show2(a int, b int) {
	if a < b {
		printDigit(a / 10)
		printDigit(a % 10)
		z01.PrintRune(' ')
		printDigit(b / 10)
		printDigit(b % 10)
		if !(a == 98 && b == 99) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}

func PrintComb2() {
	a := 0
	b := 0
	for a <= 98 {
		b = a + 1
		for b <= 99 {
			show2(a, b)
			b++
		}
		a++
	}
	z01.PrintRune('\n')
}
