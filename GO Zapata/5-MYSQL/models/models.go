package models

type Cliente struct {
	ID       int    `json:"id" db:"id"`
	Nombre   string `json:"nombre" db:"nombre"`
	Correo   string `json:"correo" db:"correo"`
	Telefono string `json:"telefono" db:"telefono"`
	Fecha    string `json:"fecha" db:"fecha"`
}

type Clientes []Cliente
