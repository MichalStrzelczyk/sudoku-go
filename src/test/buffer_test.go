package test

import (
	"testing"
	"buffer"	
)

func TestNew(t *testing.T) {
	buffer := buffer.New()

	if buffer.List != [9]int8{1,2,3,4,5,6,7,8,9} {
		t.Errorf("Object buffer is inccorect")
	}		
}

func TestClear(t *testing.T) {
	buffer := buffer.New()
	buffer.Clear()

	if buffer.List != [9]int8{0,0,0,0,0,0,0,0,0} {
		t.Errorf("buffer.Clear not working properly")
	}		
}

func TestFill(t *testing.T) {
	buffer := buffer.New()
	buffer.Clear()
	buffer.Fill()

	if buffer.List != [9]int8{1,2,3,4,5,6,7,8,9} {
		t.Errorf("buffer.Fill not working properly.")
	}		
}

func TestIsEmpty(t *testing.T) {
	buffer := buffer.New()
	buffer.Clear()
	
	if buffer.IsEmpty() == false {
		t.Errorf("buffer.IsEmpty not working properly. Buffer is empty")
	}	
	
	buffer.Fill()
	if buffer.IsEmpty() != false {
		t.Errorf("buffer.IsEmpty not working properly. Buffer is full")
	}	
}

func TestUseNumber(t *testing.T) {
	buffer := buffer.New()
	
	buffer.UseNumber(1)
	buffer.UseNumber(5)
	buffer.UseNumber(6)
	buffer.UseNumber(9)
	
	if buffer.List != [9]int8{0,2,3,4,0,0,7,8,0} {
		t.Errorf("buffer.UseNumber not working properly. Buffer should look like 0,2,3,4,0,0,7,8,0")
	}		
}

func TestGetRandomAvaiableNumberScenario1(t *testing.T) {
	buffer := buffer.New()
	
	buffer.UseNumber(1)Available
	buffer.UseNumber(2)
	buffer.UseNumber(3)
	buffer.UseNumber(4)
	buffer.UseNumber(5)
	buffer.UseNumber(6)
	buffer.UseNumber(7)
	buffer.UseNumber(8)
	
	if buffer.GetRandomAvaiableNumber() != 9 {
		t.Errorf("buffer.GetRandomAvaiableNumber not working properly. Available number is 9")
	}		
}

func TestGetRandomAvaiableNumberScenario2(t *testing.T) {
	buffer := buffer.New()
	
	buffer.UseNumber(1)
	buffer.UseNumber(2)

	buffer.UseNumber(4)
	buffer.UseNumber(5)
	buffer.UseNumber(6)
	buffer.UseNumber(7)
	buffer.UseNumber(8)
	
	number := buffer.GetRandomAvaiableNumber()
	if !(number == 3 || number == 9) {
		t.Errorf("buffer.GetRandomAvaiableNumber not working properly. Shuffled number: %d Available: [3,9]", number)
	}		
}