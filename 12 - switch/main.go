package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terca-feira"
	case 4:
		return "quarta-feira"
	case 5:
		return "quinta-feira"
	case 6:
		return "sexta-feira"
	case 7:
		return "Sabado"
	default:
		return "mumero invalido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terca-feira"
	case numero == 4:
		diaDaSemana = "quarta-feira"
	case numero == 5:
		diaDaSemana = "quinta-feira"
	case numero == 6:
		diaDaSemana = "sexta-feira"
	case numero == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "mumero invalido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("switch")

	dia := diaDaSemana(3)
	fmt.Println(dia)
}
