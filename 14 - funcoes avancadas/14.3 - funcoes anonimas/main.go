package main

import "fmt"

func main() {
	fmt.Println("funcoes anonimas")

	func() {
		fmt.Println("Ola mundo")
	}()

	retorno := func(texto string) string {
		//fmt.Println(texto)
		return fmt.Sprintf("Recebido -> %s", texto)
	}("passando a informacao para funcao")

	fmt.Println(retorno)
}
