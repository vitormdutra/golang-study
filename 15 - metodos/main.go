package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Println("Dentro do metodo salvar")
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"marcus", 24}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"vitor", 24}
	fmt.Println(usuario2)
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	fmt.Println(usuario2.idade)
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}
