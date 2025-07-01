package piscine

func IsNumeric(s string) bool {
	// Iterate over each character in the string
	for _, char := range s {
		// Check if the character is not a digit
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
