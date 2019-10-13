package sudoku

import (
	"buffer"
	"math"
	"sort"
)

// Types
type CellData struct {
	X, Y, ChosenNumber int8
	AvaiableNumbers    []int8
}

type CellCollection []CellData

// Private properties
var originalBoard [9][9]int8
var cellCollection CellCollection
var history CellCollection

// Public properties
var Board [9][9]int8
var TmpBuffer = buffer.New()

// Function for sorting
func (a CellCollection) Len() int {
	return len(a)
}

func (a CellCollection) Less(i, j int) bool {
	return len(a[i].AvaiableNumbers) < len(a[j].AvaiableNumbers)
}

func (a CellCollection) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (collection CellCollection) HasCertainties() bool {
	for _, cell := range collection {
		if len(cell.AvaiableNumbers) == 1 {
			return true
		}
	}

	return false
}

// Constructor
func New(boardToSolved [9][9]int8) {
	Board = boardToSolved
	originalBoard = boardToSolved
}

// Main function
func Solve() [9][9]int8 {
	needWork := true
	for needWork {
		buildPossibilities()
		fillCertainties()

		if !cellCollection.HasCertainties() {
			fillCeil()
			if isEmpty() {
				needWork = false
			}
		}
	}

	return Board
}

func fillCeil() {
	if len(cellCollection) > 0 {

		if len(cellCollection[0].AvaiableNumbers) == 0 {
			Board = originalBoard
			return
		}

		cellCollection[0].ChosenNumber = buffer.GetRandomValueFromList(cellCollection[0].AvaiableNumbers)
		Board[cellCollection[0].Y][cellCollection[0].X] = cellCollection[0].ChosenNumber
		history = append(history, cellCollection[0])
	}
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
				eliminateNumbers(int8(x), int8(y))
				ceilData := CellData{int8(x), int8(y), 0, TmpBuffer.GetAvaiableNumbers()}
				cellCollection = append(cellCollection, ceilData)
			}
		}
	}

	// Sort results
	sort.Sort(CellCollection(cellCollection))

	return Board
}

func eliminateNumbers(x int8, y int8) {
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
