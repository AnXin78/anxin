package piscine

func AppendRange(min, max int) []int {
	// If min is greater than or equal to max, return nil
	if min >= max {
		return nil
	}

	// Initialize an empty slice to hold the result
	var result []int

	// Loop from min to max-1 and append each value to the slice
	for i := min; i < max; i++ {
		result = append(result, i)
	}

	// Return the resulting slice
	return result
}
