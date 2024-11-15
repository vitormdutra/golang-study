package main

import "fmt"

func closure() func() {
	text := "Dentro da funcao closure"

	funcao := func() {
		fmt.Println(text)
	}

	return funcao
}

func main() {
	text := "Dentro da funcao main"
	fmt.Println(text)

	funcaoNova := closure()
	funcaoNova()
}
