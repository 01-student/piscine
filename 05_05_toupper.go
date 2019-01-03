package piscine

func toUpper(c byte) byte {
	if isLower(c) {
		return c - 'a' + 'A'
	}
	return c
}

func ToUpper(s string) string {
	b := []byte(s)
	for i, c := range b {
		b[i] = toUpper(c)
	}
	return string(b)
}
