package piscine

func isUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func IsUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if !isUpper(s[i]) {
			return false
		}
	}
	return true
}
