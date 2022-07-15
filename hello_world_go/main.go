package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello World", "1")
	fmt.Println("Hello World", 1, 2)
	/*
		Funções que retornam mais de uma informação devem ter variaveis para receber esses valores.
		Caso não for necessario usar alguma variavel deve-se ser substituido pelo caractere "_".
	*/
	numeroBytes, _ := fmt.Println("Hello World")

	fmt.Println("bytes:", numeroBytes)

	/* Tipos*/
	x := 1
	y := "hello"
	z := true

	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(y), reflect.TypeOf(z))
}
