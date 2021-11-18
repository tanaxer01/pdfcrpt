package handlers

import (
	"io"
	"log"

	"net/http"
	"encoding/json"

	"pdfcrpt/data"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    // A very simple health check.
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    // In the future we could report back on the status of our DB, or our cache
    // (e.g. Redis) by performing a simple PING, and include them in the response.
    io.WriteString(w, `{"alive": true}`)
}


func reciveInfo(w http.ResponseWriter, r *http.Request) {
	var datos data.DataReceived
	err := json.NewDecoder(r.Body).Decode(&datos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(datos)

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	io.WriteString(w, `al is god`)
}
