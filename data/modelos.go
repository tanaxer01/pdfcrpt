package data

import "time"

type DataReceived struct {
	Password	string		`json:"password"`
	IP			string		`json:"ip"`
	SO			string		`json:"so"`
	Time		time.Time	`json:"timestamp"`
}
