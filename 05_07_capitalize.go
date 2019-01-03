package piscine

func Capitalize(s string) string {
	b := []byte(s)
	word := false
	for i, c := range b {
		if IsAlpha(c) {
			if !word {
				word = true
				b[i] = toUpper(c)
			} else {
				b[i] = toLower(c)
			}
		} else {
			word = false
		}
	}
	return string(b)
}
