package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola mundo"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pagina de usuarios"))
}

func main() {
	// http protocolo de comunicacao

	// cliente (faz requesicao) - servidor (processa requisicao e envia resposta)

	// request -response

	// rotas
	// URI - identificado do recurso
	// metodo - GET, POST, PUT, DELETE
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":1234", nil))
}
