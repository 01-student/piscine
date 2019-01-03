package piscine

func Atoi(s string) int {
	sign := 1
	i := 0
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}
	return BasicAtoi2(s[i:]) * sign
}
