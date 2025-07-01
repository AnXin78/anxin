package piscine

func NRune(s string, n int) rune {
	if n < 1 || n > len(s) {
		return 0
	}
	for i, r := range s {
		if i+1 == n {
			return r
		}
	}
	return 0
}
