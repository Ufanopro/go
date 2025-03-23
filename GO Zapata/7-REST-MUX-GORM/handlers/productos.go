package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/db"
	"restapi/dto"
	"strconv"

	modelos "restapi/models"

	"github.com/gorilla/mux"
)

// ProductosGET maneja la solicitud para obtener la lista de productos
func ProductosGET(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	database := db.GetDB()

	var productos []modelos.Producto
	if err := database.Preload("Categoria").Order("id desc").Find(&productos).Error; err != nil {
		http.Error(response, "Error al obtener los productos", http.StatusInternalServerError)
		return
	}

	var productosDTO []dto.ProductoDTO
	for _, p := range productos {
		productosDTO = append(productosDTO, dto.ProductoDTO{
			Nombre:      p.Nombre,
			Precio:      p.Precio,
			Stock:       p.Stock,
			Descripcion: p.Descripcion,
			CategoriaID: p.CategoriaID,
		})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(productosDTO)
}

// ProductosGETID maneja la solicitud para obtener un producto por ID
func ProductosGETID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(response, "ID inválido", http.StatusBadRequest)
		return
	}

	database := db.GetDB()

	var producto modelos.Producto
	if err := database.Preload("Categoria").First(&producto, id).Error; err != nil {
		http.Error(response, "Producto no encontrado", http.StatusNotFound)
		return
	}

	productoDTO := dto.ProductoDTO{
		Nombre:      producto.Nombre,
		Precio:      producto.Precio,
		Stock:       producto.Stock,
		Descripcion: producto.Descripcion,
		CategoriaID: producto.CategoriaID,
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(productoDTO)
}

// ProductosPOST maneja la solicitud para crear un nuevo producto
func ProductosPOST(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var productoDTO dto.ProductoDTO
	if err := json.NewDecoder(request.Body).Decode(&productoDTO); err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "JSON inválido"}`, http.StatusBadRequest)
		return
	}

	if productoDTO.Nombre == "" || productoDTO.Precio <= 0 || productoDTO.Stock < 0 {
		http.Error(response, `{"estado": "error", "mensaje": "Datos inválidos"}`, http.StatusBadRequest)
		return
	}

	producto := modelos.Producto{
		Nombre:      productoDTO.Nombre,
		Slug:        generarSlug(productoDTO.Nombre),
		Precio:      productoDTO.Precio,
		Stock:       productoDTO.Stock,
		Descripcion: productoDTO.Descripcion,
		CategoriaID: productoDTO.CategoriaID,
	}

	database := db.GetDB()
	if err := database.Create(&producto).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Error al guardar en la base de datos"}`, http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(map[string]string{
		"estado":  "ok",
		"mensaje": "Se creó el producto exitosamente",
	})
}

// ProductosPUT maneja la solicitud para actualizar un producto
func ProductosPUT(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "ID inválido"}`, http.StatusBadRequest)
		return
	}

	var productoDTO dto.ProductoDTO
	if err := json.NewDecoder(request.Body).Decode(&productoDTO); err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "JSON inválido"}`, http.StatusBadRequest)
		return
	}

	if productoDTO.Nombre == "" || productoDTO.Precio <= 0 || productoDTO.Stock < 0 {
		http.Error(response, `{"estado": "error", "mensaje": "Datos inválidos"}`, http.StatusBadRequest)
		return
	}

	database := db.GetDB()

	var producto modelos.Producto
	if err := database.First(&producto, id).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Producto no encontrado"}`, http.StatusNotFound)
		return
	}

	producto.Nombre = productoDTO.Nombre
	producto.Slug = generarSlug(productoDTO.Nombre)
	producto.Precio = productoDTO.Precio
	producto.Stock = productoDTO.Stock
	producto.Descripcion = productoDTO.Descripcion
	producto.CategoriaID = productoDTO.CategoriaID

	if err := database.Save(&producto).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Error al actualizar en la base de datos"}`, http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]string{
		"estado":  "ok",
		"mensaje": "Producto actualizado exitosamente",
	})
}

// ProductosDEL maneja la solicitud para eliminar un producto
func ProductosDEL(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "ID inválido"}`, http.StatusBadRequest)
		return
	}
	database := db.GetDB()

	var producto modelos.Producto
	if err := database.First(&producto, id).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Producto no encontrado"}`, http.StatusNotFound)
		return
	}

	if err := database.Delete(&producto).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Error al eliminar en la base de datos"}`, http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]string{
		"estado":  "ok",
		"mensaje": "Producto eliminado exitosamente",
	})
}
