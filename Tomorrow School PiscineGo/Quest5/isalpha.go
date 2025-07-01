package piscine

func IsAlpha(s string) bool {
	// Iterate through each character in the string
	for _, char := range s {
		// Check if the character is not an uppercase letter, lowercase letter, or digit
		if !(char >= 'A' && char <= 'Z') && !(char >= 'a' && char <= 'z') && !(char >= '0' && char <= '9') {
			return false
		}
	}
	return true
}
