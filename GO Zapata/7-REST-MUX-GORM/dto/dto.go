package dto

type CategoriaDto struct {
	Nombre string `json:"nombre"`
}
type ProductoDTO struct {
	Nombre      string `json:"nombre" binding:"required"`
	Precio      int    `json:"precio" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
	CategoriaID uint   `json:"categoria_id" binding:"required"`
}
type PerfilDto struct {
	Nombre string `json:"nombre"`
}
type UsuarioDto struct {
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
	Password string `json:"password"`
	PerfilID uint   `json:"perfil_id"`
}
type LoginDto struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}

type LoginRespuestaDto struct {
	Nombre string `json:"nombre"`
	Token  string `json:"token"`
}

type FotoDTO struct {
	Tipo         string `json:"tipo" binding:"required,oneof=productos categorias"`
	Nombre       string `json:"nombre" binding:"required"`
	ReferenciaID uint   `json:"referencia_id" binding:"required"`
}
