package Slice

import "fmt"

func LengthSLice(){
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	// o slice e mais flexivel que o array
	// o slice nao tem tamanho fixo

	fmt.Println(len(slice))
}