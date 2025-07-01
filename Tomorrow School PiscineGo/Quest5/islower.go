package piscine

func IsLower(s string) bool {
	// Iterate through the string and check each character
	for _, char := range s {
		if char < 'a' || char > 'z' {
			return false
		}
	}
	return true
}
