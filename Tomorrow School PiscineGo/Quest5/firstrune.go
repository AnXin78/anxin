package piscine

func FirstRune(s string) rune {
	for _, r := range s {
		return r // как только найдена первая руна — сразу возвращаем
	}
	return 0 // если строка пустая
}
