package mapping

func MapAsFlatList(matrix [9][9]int8) [81]int8 {
	var board [81]int8
	a := 0
	for i:= 0; i<9; i++ {
		for j:= 0; j<9; j++ {
			board[a] = matrix[i][j]
			a++
		}
	}

	return board
}