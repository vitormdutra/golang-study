package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float64
}

type estudante struct {
	pessoa    //jeito que eh feito heranca
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Teste")

	p1 := pessoa{"joao", "pedro", 20, 1.68}
	fmt.Println(p1)

	e1 := estudante{p1, "engenharia", "USP"}
	fmt.Println(e1.nome)

}
