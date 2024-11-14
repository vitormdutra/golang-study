package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("arqivo struct")

	var u usuario
	fmt.Println(u)
	u.nome = "Vitor"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua tal", 2}

	usuario2 := usuario{"Vitor", 21, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 21}
	fmt.Println(usuario3)
}
