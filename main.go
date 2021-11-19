package main

import (
	handle "github.com/gorilla/handlers"
	"pdfcrpt/handlers"
	"net/http"

	"log"
)

func main() {
	log.Println("[+] Escuchando en el puerto 8080")

	headersOk := handle.AllowedHeaders([]string{"X-Requested-With"})
    originsOk := handle.AllowedOrigins([]string{"*"})
    methodsOk := handle.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	r := handlers.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", handle.CORS(headersOk, originsOk, methodsOk)(r)))

}
