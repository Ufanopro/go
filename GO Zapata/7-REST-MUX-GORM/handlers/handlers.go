package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"restapi/dto"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type ResponseGenerico struct {
	Estado  string
	Mensaje string
}

func EjemploGET(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(ResponseGenerico{"ok", "Metodo GET"})
	fmt.Fprintln(response, string(output))
}

func EjemploParams(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	response.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(ResponseGenerico{"ok", "Hola con parametros | id = " + vars["id"]})
	fmt.Fprintln(response, string(output))
}

func EjemploPOST(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var categoria dto.CategoriaDto
	err := json.NewDecoder(request.Body).Decode((&categoria))
	if err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(respuesta)
	}
	respuesta := map[string]string{
		"estado":        "error",
		"mensaje":       "Ocurrió un error inesperado",
		"nombre":        categoria.Nombre,
		"authotization": request.Header.Get("authotization"),
	}
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(respuesta)
}

/*
	func EjemploPOST(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		respuesta := map[string]string{
			"estado":  "ok",
			"mensaje": "Metodo POST2",
		}
		//response.WriteHeader(201)
		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(respuesta)
	}

	func EjemploPOST(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		output, _ := json.Marshal(ResponseGenerico{"ok", "Hola desde POST"})
		fmt.Fprintln(response, string(output))
	}
*/
func EjemploPUT(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	response.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(ResponseGenerico{"ok", "Hola con parametros en PUT | id = " + vars["id"]})
	fmt.Fprintln(response, string(output))
}

func EjemploDEL(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	response.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(ResponseGenerico{"ok", "Hola con parametros en DELETE | id = " + vars["id"]})
	fmt.Fprintln(response, string(output))
}

func EjemploQuery(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(ResponseGenerico{"ok", "query string | id=" + request.URL.Query().Get("id")})
	fmt.Fprintln(response, string(output))
}

func EjemploUpload(response http.ResponseWriter, request *http.Request) {
	// Verifica que la solicitud sea tipo multipart
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

	// Verifica que handler no sea nil antes de acceder a Filename
	if handler == nil {
		http.Error(response, "No se proporcionó un archivo", http.StatusBadRequest)
		return
	}

	// Obtiene la extensión del archivo de forma segura
	extension := filepath.Ext(handler.Filename)
	if extension == "" {
		http.Error(response, "Formato de archivo inválido", http.StatusBadRequest)
		return
	}

	// Genera el nombre del archivo con la hora actual
	foto := time.Now().Format("150405") + extension

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

// Permite visualizar una imagen subida
func VerFoto(response http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")
	folder := request.URL.Query().Get("folder")

	// Validaciones
	if len(file) == 0 || len(folder) == 0 {
		http.Error(response, "Parámetros inválidos", http.StatusBadRequest)
		return
	}

	// Construye la ruta del archivo
	rutaArchivo := filepath.Join("public/uploads", folder, file)

	// Verifica que el archivo exista
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

	// Detecta el tipo de contenido
	contentType := "image/jpeg"
	if strings.HasSuffix(file, ".png") {
		contentType = "image/png"
	}

	// Configura los headers y envía la imagen
	response.Header().Set("Content-Type", contentType)
	response.WriteHeader(http.StatusOK)
	io.Copy(response, OpenFile)
}
