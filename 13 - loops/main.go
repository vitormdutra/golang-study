package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		fmt.Println("incrementando:", i)
		time.Sleep(time.Second)
		i++
	}

	fmt.Println(i)

	for j := 0; j < 10; j += 2 {
		fmt.Println("incrementando:", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"vitor", "joao", "davi"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "leonardo",
		"sobrenome": "silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("infinito")
	}
}
