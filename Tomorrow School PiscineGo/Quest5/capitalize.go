package piscine

func Capitalize(s string) string {
	result := ""
	capitalize := true

	for _, char := range s {
		// Check if the character is alphanumeric (letter or number)
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if capitalize {
				// Capitalize the first letter of each word
				if char >= 'a' && char <= 'z' {
					result += string(char - 'a' + 'A') // Convert lowercase to uppercase
				} else {
					result += string(char) // Keep uppercase as is
				}
				capitalize = false
			} else {
				// Lowercase the rest of the characters
				if char >= 'A' && char <= 'Z' {
					result += string(char + 'a' - 'A') // Convert uppercase to lowercase
				} else {
					result += string(char) // Keep digits as is
				}
			}
		} else {
			// For non-alphanumeric characters, just add them without change
			result += string(char)
			capitalize = true // Next alphanumeric character should be capitalized
		}
	}

	return result
}
