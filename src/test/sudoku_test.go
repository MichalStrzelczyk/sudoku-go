package test

import (
	"fmt"
	"sudoku"
	"testing"
)

func TestEliminateNumbers(t *testing.T) {
	board := [9][9]int8{
		{0, 0, 0, 6, 0, 8, 9, 1, 0},
		{0, 9, 0, 4, 3, 2, 6, 8, 7},
		{0, 6, 3, 9, 0, 0, 2, 5, 4},
		{9, 0, 0, 3, 0, 4, 0, 0, 2},
		{3, 1, 0, 0, 0, 0, 0, 7, 9},
		{0, 7, 0, 0, 0, 9, 5, 0, 0},
		{0, 0, 1, 0, 9, 6, 3, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 2, 0},
		{0, 0, 6, 7, 4, 0, 0, 8, 1}}

	sudoku.New(board)

	// Scenario 1
	sudoku.EliminateNumbers(8, 0)
	if sudoku.TmpBuffer.List != [9]int8{0, 0, 3, 0, 0, 0, 0, 0, 0} {
		t.Errorf("Incorrect buffer")
		fmt.Println("Given:", sudoku.TmpBuffer.List)
		fmt.Println("Expected: ", [9]int8{0, 0, 3, 0, 0, 0, 0, 0, 0})
	}

	// Scenario 2
	sudoku.EliminateNumbers(3, 7)
	if sudoku.TmpBuffer.List != [9]int8{0, 0, 0, 0, 5, 0, 0, 8, 0} {
		t.Errorf("Incorrect buffer")
		fmt.Println("Given:", sudoku.TmpBuffer.List)
		fmt.Println("Expected: ", [9]int8{0, 0, 0, 0, 5, 0, 0, 8, 0})
	}
}

func TestRestoreSquereData(t *testing.T) {
	board := [9][9]int8{
		{0, 0, 0, 6, 0, 8, 9, 1, 0},
		{0, 9, 0, 4, 3, 2, 6, 8, 7},
		{0, 6, 3, 9, 0, 0, 2, 5, 4},
		{9, 0, 0, 3, 0, 4, 0, 0, 2},
		{3, 1, 0, 0, 0, 0, 0, 7, 9},
		{0, 7, 0, 0, 0, 9, 5, 0, 0},
		{0, 0, 1, 0, 9, 6, 3, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 2, 0},
		{0, 0, 6, 7, 4, 0, 0, 8, 1}}

	sudoku.New(board)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sudoku.Board[i][j] = 9
		}
	}

	for i := 0; i < 9; i++ {
		sudoku.RestoreSquereData(int8(i))
	}

	if sudoku.Board != board {
		t.Errorf("Data in board wasn't restored correctly")
		fmt.Println("Given:    ", sudoku.Board)
		fmt.Println("Expected: ", board)
	}
}

func TestSolveSquare(t *testing.T) {
	board := [9][9]int8{
		{4, 2, 0, 6, 0, 8, 9, 1, 0},
		{5, 9, 1, 4, 3, 2, 6, 8, 7},
		{0, 6, 3, 9, 0, 0, 2, 5, 4},
		{9, 0, 0, 3, 0, 4, 0, 0, 2},
		{3, 1, 0, 0, 0, 0, 0, 7, 9},
		{0, 7, 0, 0, 0, 9, 5, 0, 0},
		{0, 0, 1, 0, 9, 6, 3, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 2, 0},
		{0, 0, 6, 7, 4, 0, 0, 8, 1}}

	solvedBoard := [9][9]int8{
		{4, 2, 7, 6, 0, 8, 9, 1, 0},
		{5, 9, 1, 4, 3, 2, 6, 8, 7},
		{8, 6, 3, 9, 0, 0, 2, 5, 4},
		{9, 0, 0, 3, 0, 4, 0, 0, 2},
		{3, 1, 0, 0, 0, 0, 0, 7, 9},
		{0, 7, 0, 0, 0, 9, 5, 0, 0},
		{0, 0, 1, 0, 9, 6, 3, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 2, 0},
		{0, 0, 6, 7, 4, 0, 0, 8, 1}}

	sudoku.New(board)
	sudoku.SolveSquare(0)

	if sudoku.Board != solvedBoard {
		t.Errorf("Data in board wasn't restored correctly")
		fmt.Println("Given:    ", sudoku.Board)
		fmt.Println("Expected: ", solvedBoard)
	}
}
