package piscine

func toLower(c byte) byte {
	if isUpper(c) {
		return c + 'a' - 'A'
	}
	return c
}

func ToLower(s string) string {
	b := []byte(s)
	for i, c := range b {
		b[i] = toLower(c)
	}
	return string(b)
}
