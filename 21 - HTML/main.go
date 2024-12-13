package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{
		"Joao",
		"joao.pedro@gmail.com",
	}

	templates.ExecuteTemplate(w, "home.html", u)
}

type usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("Escutando na porta 1234")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
