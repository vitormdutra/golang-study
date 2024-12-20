package main

import (
	"crud/servidor"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// CRUD - CREATE, READ, UPDATE, DELETE
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods("POST")
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods("GET")
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods("GET")
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods("DELETE")

	fmt.Println("Escutando na porta 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
