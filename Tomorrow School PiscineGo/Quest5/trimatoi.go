package piscine

func TrimAtoi(s string) int {
	slice := []int{}
	counter := 0
	minus := 0

	for _, digit := range s {
		if digit >= '0' && digit <= '9' {
			a := int(digit - '0')
			slice = append(slice, a)
			counter++
		} else if digit == '-' && counter == 0 {
			minus = 1
		}
	}
	answer := 0
	for _, i := range slice {
		answer = answer*10 + i
	}
	if minus == 1 {
		answer = -answer
	}
	return (answer)
}
