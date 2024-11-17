package main

import (
	"fmt"
	"time"
)

func main() {
	// concorencia != paralelismo
	go escrever("Ola mundo!") // goroutine
	escrever("Programando em GO!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
