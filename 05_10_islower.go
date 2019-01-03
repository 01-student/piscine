package piscine

func isLower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func IsLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if !isLower(s[i]) {
			return false
		}
	}
	return true
}
