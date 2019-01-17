package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	switch s[0] {
	case '-':
		return BasicAtoi2(s[1:]) * -1
	case '+':
		return BasicAtoi2(s[1:])
	default:
		return BasicAtoi2(s)
	}
}
