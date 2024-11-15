package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
	//fmt.Println(numeros)
}

func escrever(texto string, numeros ...int) {
	for _, numeros := range numeros {
		fmt.Println(texto, numeros)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6)
	fmt.Println(totalDaSoma)

	escrever("Ola mundo", 123, 425, 6, 7, 87)
}
