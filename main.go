package main

import (
	"pdfcrpt/handlers"
	"net/http"
)

func main() {
	r := handlers.NewRouter()
	http.ListenAndServe(":8080", r)

}
