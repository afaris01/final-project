package models

type Todo struct {
	Id       int    `json:"id"`
	Nama    string `json:"nama"`
	Selesai bool   `json:"selesai"`
}
