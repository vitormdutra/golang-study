package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("escrendo algo aqui")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("teste@gmail.com")
	fmt.Println(erro)
}
