package piscine

//import "github.com/01-edu/z01"

func AlphaCount(s string) int {
	count := 0
	for _, char := range s {
		// Проверяем, является ли символ буквой
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			count++
			//z01.PrintRune(char) // Выводим каждую найденную букву
		}
	}
	return count
}
