package models

type Usuario struct {
	ID       int64  `json:"id"`
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
