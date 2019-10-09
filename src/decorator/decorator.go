package decorator

import (
	"fmt"
)

func Print(matrix [81]int8) {
	fmt.Println(matrix[0:8])		
	fmt.Println(matrix[0+1*9:8+1*9])		
	fmt.Println(matrix[0+2*9:8+2*9])		

	fmt.Println(matrix[0+3*9:8+3*9])		
	fmt.Println(matrix[0+4*9:8+4*9])		
	fmt.Println(matrix[0+5*9:8+5*9])		

	fmt.Println(matrix[0+6*9:8+6*9])		
	fmt.Println(matrix[0+7*9:8+7*9])		
	fmt.Println(matrix[0+8*9:8+8*9])		
}