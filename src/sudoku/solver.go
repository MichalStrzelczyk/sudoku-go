package main

import (
	"fmt"
	"mapping"
)

func main() {
/*	
	horizontalBuffer := buffer.Create()
	verticalBuffer := buffer.Create()
	squareBuffer := buffer.Create()
*/
	//var board [9][9]int8
	var emptyBoard [9][9]int8
	board := mapping.MapFromMatrix(emptyBoard)
	
	fmt.Println(board)	
}