package main

import (
	"decorator"
	"fmt"
	"sudoku"
	"time"
)

func main() {

	board := [9][9]int8{
		{4, 0, 0, 0, 9, 0, 2, 0, 8},
		{6, 0, 9, 0, 0, 0, 0, 5, 0},
		{0, 0, 3, 4, 0, 0, 6, 0, 9},
		{5, 0, 7, 0, 0, 8, 0, 0, 1},
		{3, 0, 8, 7, 0, 6, 5, 0, 2},
		{9, 0, 0, 5, 0, 0, 3, 0, 7},
		{8, 0, 5, 0, 0, 9, 7, 0, 0},
		{0, 9, 0, 0, 0, 0, 8, 0, 5},
		{7, 0, 2, 0, 6, 0, 0, 0, 3}}

	/*
		board := [9][9]int8{
			{3, 0, 2, 6, 0, 0, 9, 0, 1},
			{0, 0, 0, 9, 1, 0, 0, 0, 2},
			{0, 9, 0, 0, 5, 4, 0, 0, 8},
			{0, 2, 0, 0, 4, 5, 8, 1, 7},
			{8, 5, 0, 7, 0, 0, 3, 0, 0},
			{4, 0, 0, 0, 0, 0, 2, 6, 5},
			{6, 0, 5, 0, 0, 9, 0, 2, 0},
			{0, 3, 0, 0, 0, 2, 5, 0, 0},
			{0, 0, 9, 5, 0, 8, 0, 4, 6}}
	*/

	start := time.Now()
	sudoku.New(board)
	solvedBoard := sudoku.Solve()
	decorator.Print(solvedBoard)
	fmt.Printf("\nSolved in %6.6f sek\n", time.Since(start).Seconds())
}