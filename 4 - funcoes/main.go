package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	fmt.Println("teste")

	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("teste de funcao")

	resultadoSoma, resultadoSub := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSub)

	resultadoSoma1, _ := calculosMatematicos(43, 55)
	fmt.Println(resultadoSoma1)
}
