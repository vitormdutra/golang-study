package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	fmt.Println("retorno nomeado")

	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
