package piscine

func Join(a []string, sep string) (b string) {
	for i, s := range a {
		b += s
		if i < len(a)-1 {
			b += sep
		}
	}
	return b
}
