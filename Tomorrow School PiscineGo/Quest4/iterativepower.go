package piscine

func IterativePower(nb int, power int) int {
	if nb > 20 {
		return 0
	}
	if power < 0 {
		return 0
	} else {
		result := 1
		for i := 0; i < power; i++ {
			result *= nb
		}
		return result
	}
}
