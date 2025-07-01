package piscine

func DivMod(a int, b int, div *int, mod *int) {
	// Calculate the quotient and store it in the variable pointed by div
	*div = a / b
	// Calculate the remainder and store it in the variable pointed by mod
	*mod = a % b
}
