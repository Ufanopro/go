package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/db"
	"restapi/dto"
	"strconv"
	"strings"

	modelos "restapi/models"

	"github.com/gorilla/mux"
)

// CategoriasGET maneja la solicitud para obtener la lista de categorías
func CategoriasGET(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	// Se obtiene la instancia de la base de datos
	database := db.GetDB()

	// Se crea un slice de categorías
	var categorias []modelos.Categoria

	// Se obtienen los datos ordenados por ID descendente
	if err := database.Order("id desc").Find(&categorias).Error; err != nil {
		http.Error(response, "Error al obtener las categorías", http.StatusInternalServerError)
		return
	}

	// Se transforman los datos a DTO antes de enviarlos
	var categoriasDTO []dto.CategoriaDto
	for _, c := range categorias {
		categoriasDTO = append(categoriasDTO, dto.CategoriaDto{Nombre: c.Nombre})
	}

	// Se responde con los datos en formato JSON
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(categoriasDTO)
}

// CategoriasGETID maneja la solicitud para obtener una categoría por ID
func CategoriasGETID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	// Obtener el ID desde los parámetros de la URL
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(response, "ID inválido", http.StatusBadRequest)
		return
	}

	// Obtener la instancia de la base de datos
	database := db.GetDB()

	// Buscar la categoría en la BD
	var categoria modelos.Categoria
	if err := database.First(&categoria, id).Error; err != nil {
		http.Error(response, "Categoría no encontrada", http.StatusNotFound)
		return
	}

	// Convertir la categoría a DTO antes de enviarla
	categoriaDTO := dto.CategoriaDto{Nombre: categoria.Nombre}

	// Responder con la categoría en formato JSON
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(categoriaDTO)
}

// CategoriaPOST maneja la solicitud para crear una nueva categoría
func CategoriasPOST(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	// Decodificar el cuerpo de la solicitud
	var categoriaDTO dto.CategoriaDto
	if err := json.NewDecoder(request.Body).Decode(&categoriaDTO); err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "JSON inválido"}`, http.StatusBadRequest)
		return
	}

	// Validar que el nombre no esté vacío
	if categoriaDTO.Nombre == "" {
		http.Error(response, `{"estado": "error", "mensaje": "El nombre es obligatorio"}`, http.StatusBadRequest)
		return
	}

	// Crear la nueva categoría
	nuevaCategoria := modelos.Categoria{
		Nombre: categoriaDTO.Nombre,
		Slug:   generarSlug(categoriaDTO.Nombre),
	}

	// Obtener la instancia de la base de datos
	database := db.GetDB()
	if err := database.Create(&nuevaCategoria).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Error al guardar en la base de datos"}`, http.StatusInternalServerError)
		return
	}

	// Respuesta de éxito
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(map[string]string{
		"estado":  "ok",
		"mensaje": "Se creó el registro exitosamente",
	})
}

// generarSlug genera un slug simple a partir del nombre
func generarSlug(nombre string) string {
	return strings.ToLower(strings.ReplaceAll(nombre, " ", "-"))
}

// CategoriasPUT maneja la solicitud para actualizar una categoría por ID
func CategoriasPUT(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	// Verificar que el método es PUT
	if request.Method != http.MethodPut {
		http.Error(response, `{"estado": "error", "mensaje": "Método no permitido"}`, http.StatusMethodNotAllowed)
		return
	}

	// Obtener el ID desde los parámetros de la URL
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "ID inválido"}`, http.StatusBadRequest)
		return
	}

	// Decodificar el cuerpo de la solicitud
	var categoriaDTO dto.CategoriaDto
	if err := json.NewDecoder(request.Body).Decode(&categoriaDTO); err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "JSON inválido"}`, http.StatusBadRequest)
		return
	}

	// Validar que el nombre no esté vacío
	if categoriaDTO.Nombre == "" {
		http.Error(response, `{"estado": "error", "mensaje": "El nombre es obligatorio"}`, http.StatusBadRequest)
		return
	}

	// Obtener la instancia de la base de datos
	database := db.GetDB()

	// Buscar la categoría en la BD
	var categoria modelos.Categoria
	if err := database.First(&categoria, id).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Categoría no encontrada"}`, http.StatusNotFound)
		return
	}

	// Actualizar los valores
	categoria.Nombre = categoriaDTO.Nombre
	categoria.Slug = generarSlug(categoriaDTO.Nombre)

	// Guardar los cambios en la BD
	if err := database.Save(&categoria).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Error al actualizar en la base de datos"}`, http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]string{
		"estado":  "ok",
		"mensaje": "Registro actualizado exitosamente",
	})
}

// CategoriasDEL maneja la solicitud para eliminar una categoría por ID
func CategoriasDEL(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	// Obtener el ID desde los parámetros de la URL
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "ID inválido"}`, http.StatusBadRequest)
		return
	}

	// Obtener la instancia de la base de datos
	database := db.GetDB()

	// Buscar la categoría en la BD
	var categoria modelos.Categoria
	if err := database.First(&categoria, id).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Categoría no encontrada"}`, http.StatusNotFound)
		return
	}

	// Eliminar la categoría
	if err := database.Delete(&categoria).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Error al eliminar en la base de datos"}`, http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]string{
		"estado":  "ok",
		"mensaje": "Registro eliminado exitosamente",
	})
}
