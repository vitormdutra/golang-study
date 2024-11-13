package main

import "fmt"

func main() {

	//aritmeticos
	soma := 1 + 2
	sub := 5 - 9
	divisao := 6 / 2
	multi := 5 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, sub, divisao, multi, restoDaDivisao)

	// tipos diferentes da erro
	var num1 int16 = 10
	var num2 int16 = 11

	soma1 := num1 + num2
	fmt.Println(soma1)

	// atribuicao

	var variavel string = "atribuido"
	variavel1 := "atribuido tbm"

	fmt.Println(variavel, variavel1)

	// operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	fmt.Println("--------------")
	// operadores logicos

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// operadores unarios
	nume := 10
	nume++
	nume += 15 // nume = nume + 15
	fmt.Println(nume)

	nume--
	nume -= 4 // nume = nume - 4
	fmt.Println(nume)

	nume /= 10 // nume = nume / 10
	nume *= 4  // nume = nume * 4
	nume %= 3  // nume = nume % 3

	fmt.Println(nume)

	var texto string
	if nume > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}

	fmt.Println(texto)
}
