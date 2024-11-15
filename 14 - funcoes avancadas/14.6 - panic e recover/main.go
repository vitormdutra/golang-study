package main

import "fmt"

func recuperarExecuca() {
	fmt.Println("tentativa de recuperar a execucao")

	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecuca()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("Media exatamente 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 8))
	fmt.Println("Pos execucao")

}
