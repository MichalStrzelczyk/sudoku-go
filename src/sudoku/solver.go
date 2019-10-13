package sudoku

import (
	"buffer"
	"fmt"
	"math"
	"sort"
)

// Private
var originalBoard [9][9]int8
var maxCyclesPerRow = 500
var cycles = 0
var Board [9][9]int8
var TmpBuffer = buffer.New()

type CellData struct {
	X, Y, ChosenNumber int8
	AvaiableNumbers    []int8
}

type CellCollection []CellData

var cellCollection CellCollection
var history CellCollection

func (a CellCollection) Len() int { return len(a) }
func (a CellCollection) Less(i, j int) bool {
	return len(a[i].AvaiableNumbers) < len(a[j].AvaiableNumbers)
}
func (a CellCollection) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (collection CellCollection) HasCertainties() bool {
	for _, cell := range collection {
		if len(cell.AvaiableNumbers) == 1 {
			return true
		}
	}

	return false
}

func New(boardToSolved [9][9]int8) {
	Board = boardToSolved
	originalBoard = boardToSolved
}

func Solve() [9][9]int8 {
	i := 0
	needWork := true
	for needWork {
		buildPossibilities()
		fillCertainties()

		if !cellCollection.HasCertainties() {
			needWork = false
		}

		/*
			buildPossibilities()
			fillBoard()

			needWork = false
			return Board

			if isEmpty() {
				fmt.Println(history)
				fmt.Println(i)
				return Board
				needWork = false
			}*/
		/*
			if cellCollection.HasCertainties() {
				fillBoard()
			} else {
				fmt.Println(i)
				return Board
				needWork = false
			}
		*/
		i++
	}

	fmt.Println(i, cellCollection)
	return Board

}

func fillCertainties() {
	for _, ceilData := range cellCollection {
		if len(ceilData.AvaiableNumbers) == 1 {
			ceilData.ChosenNumber = ceilData.AvaiableNumbers[0]
			Board[ceilData.Y][ceilData.X] = ceilData.ChosenNumber
		}
	}
}

func isEmpty() bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if Board[x][y] == 0 {
				return false
			}
		}
	}

	return true
}

func buildPossibilities() [9][9]int8 {
	cellCollection = CellCollection{}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if Board[y][x] == 0 {
				EliminateNumbers(int8(x), int8(y))
				ceilData := CellData{int8(x), int8(y), 0, TmpBuffer.GetAvaiableNumbers()}
				cellCollection = append(cellCollection, ceilData)
			}
		}
	}

	// Sort results
	sort.Sort(CellCollection(cellCollection))

	return Board
}

func EliminateNumbers(x int8, y int8) {
	TmpBuffer.Fill()

	for i := 0; i < 9; i++ {
		// Row
		TmpBuffer.UseNumber(Board[y][i])
		// Col
		TmpBuffer.UseNumber(Board[i][x])
	}

	// Square
	fromX := int8((math.Floor(float64(x / 3))) * 3)
	fromY := int8((math.Floor(float64(y / 3))) * 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			y := fromY + int8(i)
			x := fromX + int8(j)
			TmpBuffer.UseNumber(Board[y][x])
		}
	}
}
