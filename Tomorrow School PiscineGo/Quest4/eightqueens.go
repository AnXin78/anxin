package main

//import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(0, &board)
}

func solve(col int, board *[8]int) {
	if col == 8 {
		printBoard(*board)
		return
	}

	for row := 1; row <= 8; row++ {
		if isSafe(col, row, *board) {
			board[col] = row
			solve(col+1, board)
		}
	}
}

func isSafe(col, row int, board [8]int) bool {
	for prevCol := 0; prevCol < col; prevCol++ {
		prevRow := board[prevCol]

		// Check same row
		if prevRow == row {
			return false
		}

		// Check diagonals
		if abs(prevRow-row) == abs(prevCol-col) {
			return false
		}
	}
	return true
}

func printBoard(board [8]int) {
	for i := 0; i < 8; i++ {
		//z01.PrintRune(rune(board[i] + '0'))
	}
	//z01.PrintRune('\n')
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
