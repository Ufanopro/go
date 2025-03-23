package modelos

import (
	"time"
)

type Categoria struct {
	Id     uint   `json:"id"`
	Nombre string `gorm:"type:varchar(100)" json:"nombre"`
	Slug   string `gorm:"type:varchar(100)" json:"slug"`
}

type Categorias []Categoria

type Producto struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre      string    `gorm:"type:varchar(100);not null" json:"nombre"`
	Slug        string    `gorm:"type:varchar(100);not null;unique" json:"slug"`
	Precio      int       `gorm:"not null" json:"precio"`
	Stock       int       `gorm:"not null" json:"stock"`
	Descripcion string    `gorm:"type:text;not null" json:"descripcion"`
	CategoriaID uint      `gorm:"not null" json:"categoria_id"`
	Categoria   Categoria `gorm:"foreignKey:CategoriaID" json:"categoria"`
}

type Productos []Producto

type ProductoFoto struct {
	Id         int      `json:"id"`
	Nombre     string   `gorm:"type:varchar(100)"  json:"nombre"`
	ProductoID int      `json:"producto_id"`
	Producto   Producto `json:"producto"`
}
type ProductoFotos []ProductoFoto

// ////////////////////////// usuarios y perfil
type Perfil struct {
	Id     uint   `json:"id"`
	Nombre string `gorm:"type:varchar(100)" json:"nombre"`
}

type Perfiles []Perfil
type Usuario struct {
	Id       uint      `json:"id"`
	PerfilID uint      `json:"perfil_id"`
	Perfil   Perfil    `json:"perfil"`
	Nombre   string    `gorm:"type:varchar(100)" json:"nombre"`
	Correo   string    `gorm:"type:varchar(100)" json:"correo"`
	Telefono string    `gorm:"type:varchar(50)" json:"telefono"`
	Password string    `gorm:"type:varchar(160)" json:"password"`
	Fecha    time.Time `json:"fecha"`
}

type Usuarios []Usuario

type Foto struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Tipo         string `gorm:"type:enum('productos', 'categorias');not null" json:"tipo"`
	Nombre       string `gorm:"type:varchar(255);not null" json:"nombre"`
	ReferenciaID uint   `gorm:"not null" json:"referencia_id"`
}

type Fotos []Foto
