package buffer

import (
	"math/rand"
)

type Buffer struct {
	list [9]int8
}

func Create() Buffer {
	var buffer Buffer
	Fill(&buffer)

	return buffer
}

func Clear(buffer *Buffer) {
	buffer.list = [9]int8{0,0,0,0,0,0,0,0,0}	
}
		
func Fill(buffer *Buffer){
	buffer.list = [9]int8{1,2,3,4,5,6,7,8,9}		
}

func IsEmpty(buffer *Buffer) bool {
	for i:=0; i<9; i++ {
		if buffer.list[i] != 0 {
			return false
		}
	}

	return true
}
int
func UseNumber(buffer *Buffer, number int8) {
	buffer.list[number-1] = 0
}

func GetCommonNumbers(buffer1 Buffer, buffer2 Buffer, buffer3 Buffer) []int8 {
	var tmpBuffer []int8
	for i:=0; i<9; i++ {
		if buffer1.list[i] != 0 && buffer2.list[i] != 0 && buffer3.list[i] != 0 {			
			tmpBuffer = append(tmpBuffer, buffer1.list[i])
		}
	}

	return tmpBuffer
}

func GetRandomAvaiableNumber(buffer Buffer) int8 {
	var tmpBuffer []int8
	for i:=0; i<9; i++ {
		if buffer.list[i] != 0 {
			tmpBuffer = append(tmpBuffer, buffer.list[i])
		}
	}
	
	return GetRandomValueFromList(tmpBuffer)
}

func GetRandomValueFromList(list []int8) int8 {
	if len(list) == 0 {
		return 0
	}

	var index = rand.Intn(len(list))

	return list[index]
}
