package piscine

func IsPrintable(s string) bool {
	// Iterate over each character in the string
	for _, char := range s {
		// Check if the character is not a printable character (ASCII 32 to 126)
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
