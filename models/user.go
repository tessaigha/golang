package models

type User struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
