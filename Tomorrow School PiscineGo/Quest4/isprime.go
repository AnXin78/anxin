package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	for a := 2; a < nb; a++ {
		if nb%a == 0 {
			return false
		}
	}
	return true
}
