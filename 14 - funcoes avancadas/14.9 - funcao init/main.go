package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando a funcao init")
	n = 10
}

func main() {
	fmt.Println("Executando a funcao main")
	fmt.Println(n)
}
