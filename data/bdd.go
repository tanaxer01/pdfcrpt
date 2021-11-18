package data

import (
	"github.com/mattn/go-sqlite3"
)


func GenBDD() {
	db, err := sql.Open("sqlite3", "./info.db")
	if err!= nil {
		log.Fatal(err)
	}
}
