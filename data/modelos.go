package data

type DataReceived struct {
	Id			int
	Password	string		`json:"password"`
	IP			string
	SO			string		`json:"so"`
}
