package models

type Todo struct {
	Id       int    `json:"id"`
	Nama    string `json:"nama"`
	Complete bool   `json:"complete"`
}
