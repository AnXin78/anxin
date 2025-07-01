package piscine

//import "github.com/01-edu/z01"

func PrintComb2() {
	a1 := '0'
	for a1 <= '9' {
		a2 := '0'
		for a2 <= '9' {
			b1 := a1
			b2 := a2 + 1
			for b1 <= '9' {
				for b2 <= '9' {
					//z01.PrintRune(a1)
					//z01.PrintRune(a2)
					//z01.PrintRune(' ')
					//z01.PrintRune(b1)
					//z01.PrintRune(b2)
					if !(a1 == '9' && a2 == '8' && b1 == '9' && b2 == '9') {
						//z01.PrintRune(',')
						//z01.PrintRune(' ')
					}
					b2++
				}
				b1++
				b2 = '0'
			}
			a2++
		}
		a1++
	}
	//z01.PrintRune('\n')
}
