package main

import "fmt"

func funcao1() {
	fmt.Println("Executando funcao 1")
}

func funcao2() {
	fmt.Println("Executando funcao 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculda, resultado sera retornado")
	fmt.Println("Entrando na funcao para ver se o aluno esta aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	defer funcao1()
	// defer = adiar
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 9))
}
