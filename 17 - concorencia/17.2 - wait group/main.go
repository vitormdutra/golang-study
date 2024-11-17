package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Ola mundo")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em GO")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // esperar a execucao dos goroutines
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
