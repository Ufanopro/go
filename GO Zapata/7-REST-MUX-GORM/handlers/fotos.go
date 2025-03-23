package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"restapi/db"
	"restapi/dto"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	modelos "restapi/models"
)

// FotosGET maneja la solicitud para obtener la lista de fotos
func FotosGET(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	database := db.GetDB()

	var fotos []modelos.Foto
	if err := database.Order("id desc").Find(&fotos).Error; err != nil {
		http.Error(response, "Error al obtener las fotos", http.StatusInternalServerError)
		return
	}

	var fotosDTO []dto.FotoDTO
	for _, f := range fotos {
		// Validar que el tipo sea "productos" o "categorias" antes de agregar al DTO
		if f.Tipo == "productos" || f.Tipo == "categorias" {
			fotosDTO = append(fotosDTO, dto.FotoDTO{
				Tipo:         f.Tipo,
				Nombre:       f.Nombre,
				ReferenciaID: f.ReferenciaID,
			})
		}
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(fotosDTO)
}

// FotosGETByReferenciaID maneja la solicitud para obtener fotos por referencia_id
func FotosGETByReferenciaID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	database := db.GetDB()

	vars := mux.Vars(request)
	referenciaID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(response, "ID de referencia inválido", http.StatusBadRequest)
		return
	}

	var fotos []modelos.Foto
	if err := database.Where("referencia_id = ?", referenciaID).Find(&fotos).Error; err != nil {
		http.Error(response, "Error al obtener las fotos", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(fotos)
}

// FotosGETByTipo maneja la solicitud para obtener fotos por tipo
func FotosGETByTipo(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	database := db.GetDB()

	vars := mux.Vars(request)
	tipo := vars["tipo"]

	var fotos []modelos.Foto
	if err := database.Where("tipo = ?", tipo).Find(&fotos).Error; err != nil {
		http.Error(response, "Error al obtener las fotos", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(fotos)
}

// FotosPOST crea una nueva foto
func FotosPOST(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var fotoDTO dto.FotoDTO
	if err := json.NewDecoder(request.Body).Decode(&fotoDTO); err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "JSON inválido"}`, http.StatusBadRequest)
		return
	}

	if fotoDTO.Tipo == "" || fotoDTO.Nombre == "" || fotoDTO.ReferenciaID <= 0 {
		http.Error(response, `{"estado": "error", "mensaje": "Datos inválidos"}`, http.StatusBadRequest)
		return
	}

	foto := modelos.Foto{
		Tipo:         fotoDTO.Tipo,
		Nombre:       fotoDTO.Nombre,
		ReferenciaID: fotoDTO.ReferenciaID,
	}
	database := db.GetDB()
	if err := database.Create(&foto).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Error al guardar en la base de datos"}`, http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(map[string]string{
		"estado":  "ok",
		"mensaje": "Foto creada exitosamente",
	})
}

// FotosPUT actualiza una foto existente
func FotosPUT(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "ID inválido"}`, http.StatusBadRequest)
		return
	}

	var fotoDTO dto.FotoDTO
	if err := json.NewDecoder(request.Body).Decode(&fotoDTO); err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "JSON inválido"}`, http.StatusBadRequest)
		return
	}

	if fotoDTO.Tipo == "" || fotoDTO.Nombre == "" || fotoDTO.ReferenciaID <= 0 {
		http.Error(response, `{"estado": "error", "mensaje": "Datos inválidos"}`, http.StatusBadRequest)
		return
	}
	database := db.GetDB()

	var foto modelos.Foto
	if err := database.First(&foto, id).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Foto no encontrada"}`, http.StatusNotFound)
		return
	}

	foto.Tipo = fotoDTO.Tipo
	foto.Nombre = fotoDTO.Nombre
	foto.ReferenciaID = fotoDTO.ReferenciaID

	if err := database.Save(&foto).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Error al actualizar en la base de datos"}`, http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]string{
		"estado":  "ok",
		"mensaje": "Foto actualizada exitosamente",
	})
}

// FotosDEL elimina una foto
func FotosDEL(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "ID inválido"}`, http.StatusBadRequest)
		return
	}
	database := db.GetDB()

	var foto modelos.Foto
	if err := database.First(&foto, id).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Foto no encontrada"}`, http.StatusNotFound)
		return
	}

	if err := database.Delete(&foto).Error; err != nil {
		http.Error(response, `{"estado": "error", "mensaje": "Error al eliminar en la base de datos"}`, http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]string{
		"estado":  "ok",
		"mensaje": "Foto eliminada exitosamente",
	})
}

func FotosUpload(response http.ResponseWriter, request *http.Request) {
	// Verifica que la solicitud sea multipart
	err := request.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(response, "Error al procesar el formulario: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Obtiene el archivo
	file, handler, err := request.FormFile("foto")
	if err != nil {
		http.Error(response, "Error al obtener el archivo: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Verifica que el archivo tiene una extensión válida
	extension := filepath.Ext(handler.Filename)
	if extension == "" {
		http.Error(response, "Formato de archivo inválido", http.StatusBadRequest)
		return
	}

	// Genera el nombre del archivo con la hora actual
	foto := fmt.Sprintf("%s%s", time.Now().Format("150405"), extension)

	// Ruta donde se guardará el archivo
	directorio := "public/uploads/fotos"
	if err := os.MkdirAll(directorio, os.ModePerm); err != nil {
		http.Error(response, "Error al crear directorio: "+err.Error(), http.StatusInternalServerError)
		return
	}
	rutaArchivo := filepath.Join(directorio, foto)

	// Crea el archivo en el servidor
	f, err := os.OpenFile(rutaArchivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(response, "Error al subir la imagen: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Copia el contenido del archivo
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(response, "Error al copiar la imagen: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta JSON
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Se creó el archivo exitosamente",
		"foto":    foto,
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(respuesta)
}

func FotoVer(response http.ResponseWriter, request *http.Request) {
	// Obtiene el nombre del archivo desde la URL
	file := request.URL.Query().Get("file")
	if file == "" {
		http.Error(response, "Parámetro 'file' es obligatorio", http.StatusBadRequest)
		return
	}

	// Construye la ruta del archivo
	rutaArchivo := filepath.Join("public/uploads/fotos", file)

	// Verifica que el archivo existe
	if _, err := os.Stat(rutaArchivo); os.IsNotExist(err) {
		http.Error(response, "Archivo no encontrado", http.StatusNotFound)
		return
	}

	// Abre el archivo
	OpenFile, err := os.Open(rutaArchivo)
	if err != nil {
		http.Error(response, "Error al abrir el archivo", http.StatusInternalServerError)
		return
	}
	defer OpenFile.Close()

	// Lee los primeros bytes para detectar el tipo MIME
	buf := make([]byte, 512)
	n, _ := OpenFile.Read(buf)
	contentType := http.DetectContentType(buf[:n])

	// Configura los headers y envía la imagen
	response.Header().Set("Content-Type", contentType)
	response.WriteHeader(http.StatusOK)
	OpenFile.Seek(0, 0) // Reinicia el puntero del archivo para la copia completa
	io.Copy(response, OpenFile)
}

func FotoDeLocal(response http.ResponseWriter, request *http.Request) {
	// Obtiene el nombre del archivo desde la URL
	file := request.URL.Query().Get("file")
	if file == "" {
		http.Error(response, "Parámetro 'file' es obligatorio", http.StatusBadRequest)
		return
	}

	// Construye la ruta del archivo
	rutaArchivo := filepath.Join("public/uploads/fotos", file)

	// Verifica si el archivo existe
	if _, err := os.Stat(rutaArchivo); os.IsNotExist(err) {
		http.Error(response, "Archivo no encontrado", http.StatusNotFound)
		return
	}

	// Intenta eliminar el archivo
	err := os.Remove(rutaArchivo)
	if err != nil {
		http.Error(response, "Error al eliminar el archivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta JSON
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Archivo eliminado correctamente",
		"foto":    file,
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(respuesta)
}
