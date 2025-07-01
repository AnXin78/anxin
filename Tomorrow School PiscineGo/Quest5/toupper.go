package piscine

func ToUpper(s string) string {
	result := ""

	for _, char := range s {
		// Check if the character is a lowercase letter
		if char >= 'a' && char <= 'z' {
			// Convert to uppercase by subtracting the difference between 'a' and 'A'
			result += string(char - 'a' + 'A')
		} else {
			// If it's not a lowercase letter, keep it unchanged
			result += string(char)
		}
	}

	return result
}
