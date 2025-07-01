package piscine

func Swap(a *int, b *int) {
	// Use a temporary variable to swap the values
	temp := *a
	*a = *b
	*b = temp
}
