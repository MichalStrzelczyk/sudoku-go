package buffer

import (
	"math/rand"
	"time"
)

type Buffer struct {
	List [9]int8
}

/* Public methods */
func New() Buffer {
	var bufTmp = Buffer{}
	bufTmp.Fill()

	return bufTmp
}

/* Object methods s*/
func (buffer *Buffer) Clear() {
	buffer.List = [9]int8{0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func (buffer *Buffer) Fill() {
	buffer.List = [9]int8{1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func (buffer *Buffer) IsEmpty() bool {
	for i := 0; i < 9; i++ {
		if buffer.List[i] != 0 {
			return false
		}
	}

	return true
}

func (buffer *Buffer) IsOnlyOnePossibility() bool {
	var cnt = 0
	for i := 0; i < 9; i++ {
		if buffer.List[i] != 0 {
			cnt++
		}
	}

	return cnt == 1
}

func (buffer *Buffer) UseNumber(number int8) {
	if number >= 1 && number <= 9 {
		buffer.List[number-1] = 0
	}
}

func (buffer *Buffer) GetAvaiableNumbers() []int8 {
	var tmpBuffer []int8
	for i := 0; i < 9; i++ {
		if buffer.List[i] != 0 {
			tmpBuffer = append(tmpBuffer, buffer.List[i])
		}
	}

	return tmpBuffer
}

func (buffer *Buffer) GetRandomAvaiableNumber() int8 {
	return GetRandomValueFromList(buffer.GetAvaiableNumbers())
}

func GetRandomValueFromList(list []int8) int8 {
	if len(list) == 0 {
		return 0
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(list))

	return list[index]
}
