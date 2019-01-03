package piscine

func IsAlpha(c byte) bool {
	return IsNumeric(c) ||
		'a' <= c && c <= 'z' ||
		'A' <= c && c <= 'Z'
}
