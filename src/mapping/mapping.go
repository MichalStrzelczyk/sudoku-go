package mapping

type Board struct {
	list [81]int8
}

func MapFromMatrix(matrix [9][9]int8) Board {
	var board Board
	a := 0
	for i:= 0; i<9; i++ {
		for j:= 0; j<9; j++ {
			board.list[a] = matrix[i][j]
			a++
		}
	}

	return board
}