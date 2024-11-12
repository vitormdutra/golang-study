package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64
	var numero int = 1000000000
	numero2 := -10000000
	fmt.Println(numero)
	fmt.Println(numero2)

	// nao aceita numero negativo
	var numero3 uint = 100000
	fmt.Println(numero3)

	// alias
	//int32
	var numero4 rune = 10000011
	fmt.Println(numero4)

	//uint8
	var numero5 byte = 100
	fmt.Println(numero5)

	//float32 float64
	var numeroReal float64 = 1.235812381747814
	fmt.Println(numeroReal)

	var numeroReal1 float32 = 1.23582
	fmt.Println(numeroReal1)

	numeroReal2 := 1.23582
	fmt.Println(numeroReal2)

	//string
	var str string = "lalalla"
	str2 := "lalallalal"

	fmt.Println(str)
	fmt.Println(str2)

	// null
	var texto int
	fmt.Println(texto)

	//	bool
	var booleano bool = true
	fmt.Println(booleano)

	// erro
	var erro error
	fmt.Println(erro)
	var erro2 error = errors.New("erro interno")
	fmt.Println(erro2)
}
