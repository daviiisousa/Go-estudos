package main

import "fmt"

// se vc usa const o seu valor nao pode ser alterado
const nome string = "davi"

func main(){
	var text string = "ola"
	// quando eu nao passo o tipo ele assume o valor que a variavel recebe
	var number = 1
	// posso declarar so o tipo de variavel mas sem atribuir valor
	var boolean bool
	// declaracao curta ja nasce com valor e tipo
	anotherText := "ola"

	fmt.Println(text, number, boolean, anotherText, nome)
}