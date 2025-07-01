package piscine

//import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	var comb [10]int
	for i := 0; i < n; i++ {
		comb[i] = i
	}

	for {
		// Печать текущей комбинации
		for i := 0; i < n; i++ {
			//z01.PrintRune(rune(comb[i]) + '0')
		}

		// Проверка: это последняя комбинация?
		if comb[0] == 10-n {
			break
		}

		//z01.PrintRune(',')
		//z01.PrintRune(' ')

		// Генерация следующей комбинации
		i := n - 1
		for i >= 0 && comb[i] == 9-(n-1-i) {
			i--
		}
		comb[i]++
		for j := i + 1; j < n; j++ {
			comb[j] = comb[j-1] + 1
		}
	}

	// Добавляем \n в самом конце
	//z01.PrintRune('\n')
}
