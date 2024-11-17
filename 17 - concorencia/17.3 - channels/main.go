package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Ola mundo", canal)

	// for {
	// 	mensagem, aberto := <-canal
	// 	fmt.Println(aberto)
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//fmt.Println(texto)
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
