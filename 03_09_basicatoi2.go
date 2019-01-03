package piscine

func BasicAtoi2(s string) int {
	for _, b := range []byte(s) {
		if b < '0' || b > '9' {
			return 0
		}
	}
	return BasicAtoi(s)
}
