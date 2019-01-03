package piscine

func Index(s string, tofind string) int {
	if len(tofind) == 0 {
		return 0
	}
	a := []rune(s)
	b := []rune(tofind)
	if len(a) < len(b) {
		return -1
	}
	start := 0
	j := 0
	for i := range s {
		if a[i] == b[j] {
			j++
		} else {
			j = 0
			start = i + 1
		}
		if j == len(b) {
			return start
		}
	}
	return -1
}
