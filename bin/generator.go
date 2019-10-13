package main

import (
	"decorator"
	"fmt"
	"sudoku"
	"time"
)

func main() {
	var board [9][9]int8

	start := time.Now()
	sudoku.New(board)
	solvedBoard := sudoku.Solve()
	decorator.Print(solvedBoard)
	fmt.Printf("\nSolved in %6.6f sek\n", time.Since(start).Seconds())
}
