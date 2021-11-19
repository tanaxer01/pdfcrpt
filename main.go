package main

import (
	"pdfcrpt/handlers"
	"net/http"

	"fmt"
)

func main() {
	fmt.Println("[+] Escuchando en el puerto 8080")

	r := handlers.NewRouter()
	http.ListenAndServe(":8080", r)

}
