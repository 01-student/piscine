package piscine

func BasicAtoi(s string) int {
	result := 0
	i := 0
	for i < len(s) {
		result = result * 10
		result = result + int(s[i]) - '0'
		i++
	}
	return result
}
