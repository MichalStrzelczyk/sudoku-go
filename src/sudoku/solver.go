package main

import (
	"fmt"
	"mapping"
	"buffer"
	"decorator"	
)

func main() {
	emptyBoard:= [9][9]int8{
		{0,0,0,6,0,8,9,1,0},
		{0,9,0,4,3,2,6,8,7},
		{0,6,3,9,0,0,2,0,4},
		{9,0,0,3,0,4,0,0,2},
		{3,1,0,0,0,0,0,7,9},
		{0,7,0,0,0,9,5,0,0},
		{0,0,1,0,9,6,3,0,0},
		{0,0,0,0,0,1,0,2,0},
		{0,0,6,7,4,0,0,8,1}}

	board := mapping.MapAsFlatList(emptyBoard)	
	loopGlobal := 0
	loopRow := 0
	maxPerRow := 10

	

	originalBoard := board
	tmpBuffer := buffer.New()

	for i:=0; i<81; i++ {
		if board[i] != 0 {
			continue			
		}

		tmpBuffer.Fill()
		rowNumber := i / 9		
		colNumber := i % 9
		EliminateHorizontalNumbers(rowNumber, board, &tmpBuffer)						
		EliminateVerticalNumbers(colNumber, board, &tmpBuffer)	
		EliminateSquareNumbers(rowNumber, colNumber, board, &tmpBuffer)	

		result := tmpBuffer.GetRandomAvaiableNumber()
		
		if result == 0 {
			loopRow++
			loopGlobal++
						
			fmt.Println("Powrot: ", loopRow, loopGlobal, i)
			fmt.Println()

			decorator.Print(board)
			

			//return for row
			indexMin := rowNumber * 9
			CopyValuesFromIndex(indexMin, board, originalBoard)

			//return for row -1
			if loopRow == maxPerRow {
				loopRow = 0
				indexMin := (rowNumber) * 9	
				CopyValuesFromIndex(indexMin, board, originalBoard)
			}
			
			i = indexMin-1

			fmt.Println("Powrot na i: ", i)

			//fmt.Println()
			//decorator.Print(board)
			continue
		}else{
			board[i] = result			
		}			
	}		

	decorator.Print(board)
}

func CopyValuesFromIndex(index int, board [81]int8, originalBoard [81]int8) {	
	for i:=0; i<9; i++ {
		board[index + i] = originalBoard[index + i]		
	}

	fmt.Println("Index: ", index)
	fmt.Println(originalBoard)
	fmt.Println(board)

}

func EliminateHorizontalNumbers(rowNumber int, board [81]int8, tmpBuffer *buffer.Buffer) {
	index := rowNumber * 9

	tmpBuffer.UseNumber(board[index + 0])
	tmpBuffer.UseNumber(board[index + 1])
	tmpBuffer.UseNumber(board[index + 2])
	tmpBuffer.UseNumber(board[index + 3])
	tmpBuffer.UseNumber(board[index + 4])
	tmpBuffer.UseNumber(board[index + 5])
	tmpBuffer.UseNumber(board[index + 6])
	tmpBuffer.UseNumber(board[index + 7])
	tmpBuffer.UseNumber(board[index + 8])
}

func EliminateVerticalNumbers(colNumber int, board [81]int8, tmpBuffer *buffer.Buffer) {
	tmpBuffer.UseNumber(board[colNumber + 0 * 9])
	tmpBuffer.UseNumber(board[colNumber + 1 * 9])
	tmpBuffer.UseNumber(board[colNumber + 2 * 9])
	tmpBuffer.UseNumber(board[colNumber + 3 * 9])
	tmpBuffer.UseNumber(board[colNumber + 4 * 9])
	tmpBuffer.UseNumber(board[colNumber + 5 * 9])
	tmpBuffer.UseNumber(board[colNumber + 6 * 9])
	tmpBuffer.UseNumber(board[colNumber + 7 * 9])
	tmpBuffer.UseNumber(board[colNumber + 8 * 9])	
}

func EliminateSquareNumbers(rowNumber int, colNumber int, board [81]int8, tmpBuffer *buffer.Buffer) {
	squareRowNr := rowNumber % 3
	squareColNr := colNumber % 3
	position := squareColNr * 27 + squareRowNr * 3
	
	tmpBuffer.UseNumber(board[position + 0])
	tmpBuffer.UseNumber(board[position + 1])	
	tmpBuffer.UseNumber(board[position + 2])	
	
	tmpBuffer.UseNumber(board[position + 9])	
	tmpBuffer.UseNumber(board[position + 10])	
	tmpBuffer.UseNumber(board[position + 11])	
	
	tmpBuffer.UseNumber(board[position + 18])	
	tmpBuffer.UseNumber(board[position + 19])	
	tmpBuffer.UseNumber(board[position + 20])	
}
