package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	variavel2 := "variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "teste"
		variavel4 string = "teste2"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "teste3", "teste4"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	variavel6, variavel5 = variavel5, variavel6
	fmt.Println(variavel5, variavel6)
}
