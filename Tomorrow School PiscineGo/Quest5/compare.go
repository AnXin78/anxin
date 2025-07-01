package piscine

func Compare(a, b string) int {
	// Compare the strings lexicographically
	if a == b {
		return 0
	}

	if a < b {
		return -1
	}

	return 1
}
