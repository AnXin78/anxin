package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	for {
		IsPrime := true
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				IsPrime = false
				break
			}
		}
		if IsPrime {
			return nb
		}
		nb++

	}
}
