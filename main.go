package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	fmt.Println("Star Server!")
	http.HandleFunc("/", homePage)

	fmt.Println("Servidor rodando: Porta: 8080")
	http.ListenAndServe(":8080", nil)

}

// http://localhost:8080/
