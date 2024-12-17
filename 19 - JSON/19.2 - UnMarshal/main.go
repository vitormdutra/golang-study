package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {
	cachorroEmJSON := `{"nome": "Rex", "raca": "Dalmata", "idade": 3}`

	var c cachorro

	fmt.Println(c)
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorro2EmJSON := `{"nome": "Toby", "raca": "SRD"}`

	c2 := make(map[string]string)

	fmt.Println(c2)
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
}
