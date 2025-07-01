package piscine

func StrLen(s string) int {
	// Initialize a counter to count the runes
	count := 0

	// Loop through each rune in the string and increment the count
	for range s {
		count++
	}

	// Return the final count of runes
	return count
}
