package router

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
	"webapp/utilidades"
)

func GetFormulario(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/formularios/formulario.html", utilidades.Front))
	template.Execute(response, nil)
	/*
		fmt.Fprintln(response, "Hola WebAPP")

		template, err := template.ParseFiles("templates/index/index.html", "templates/front/front.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(response, nil)
		}

	*/
}

func PostFormulario(response http.ResponseWriter, request *http.Request) {
	// Asegurar que el método sea POST
	if request.Method != http.MethodPost {
		http.Error(response, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Parsear el formulario
	if err := request.ParseForm(); err != nil {
		http.Error(response, "Error al procesar el formulario", http.StatusInternalServerError)
		return
	}

	// Obtener los valores del formulario
	nombre := request.FormValue("nombre")
	correo := request.FormValue("correo")
	telefono := request.FormValue("telefono")
	password := request.FormValue("password")
	confirmPassword := request.FormValue("confirm_password")

	// Expresión regular para validar correo electrónico
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// Variable para acumular errores
	var errorMsj []string

	// Validaciones
	if nombre == "" {
		errorMsj = append(errorMsj, "El campo 'Nombre' es obligatorio.")
	}
	if correo == "" {
		errorMsj = append(errorMsj, "El campo 'Correo' es obligatorio.")
	} else if !emailRegex.MatchString(correo) {
		errorMsj = append(errorMsj, "El formato del correo no es válido.")
	}
	if telefono == "" {
		errorMsj = append(errorMsj, "El campo 'Teléfono' es obligatorio.")
	}
	if password == "" {
		errorMsj = append(errorMsj, "El campo 'Contraseña' es obligatorio.")
	}
	if confirmPassword == "" {
		errorMsj = append(errorMsj, "El campo 'Confirmar Contraseña' es obligatorio.")
	}
	if password != "" && confirmPassword != "" && password != confirmPassword {
		errorMsj = append(errorMsj, "Las contraseñas no coinciden.")
	}

	// Si hay errores, los mostramos
	if len(errorMsj) > 0 {
		http.Error(response, strings.Join(errorMsj, "\n"), http.StatusBadRequest)
		return
	}

	// Si no hay errores, mostrar mensaje de éxito
	fmt.Fprintln(response, "Formulario correcto")
	fmt.Fprintf(response, "Nombre: %s\n", nombre)
	fmt.Fprintf(response, "Correo: %s\n", correo)
	fmt.Fprintf(response, "Teléfono: %s\n", telefono)
	fmt.Fprintf(response, "Contraseña: %s\n", password)
}

func UploadFormulario(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/formularios/formulario_upload.html", utilidades.Front))
	template.Execute(response, nil)
	/*
		fmt.Fprintln(response, "Hola WebAPP")

		template, err := template.ParseFiles("templates/index/index.html", "templates/front/front.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(response, nil)
		}

	*/
}

func UploadFormularioPost(response http.ResponseWriter, request *http.Request) {
	// Asegurar que el método sea POST
	if request.Method != http.MethodPost {
		http.Error(response, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Parsear el formulario
	err := request.ParseMultipartForm(10 << 20) // 10MB máximo
	if err != nil {
		http.Error(response, "Error al procesar el formulario", http.StatusInternalServerError)
		return
	}

	// Obtener el archivo
	file, handler, err := request.FormFile("foto")
	if err != nil {
		http.Error(response, "No ha subido ninguna foto", http.StatusBadRequest)
		return
	}
	defer file.Close() // Cerrar el archivo después de usarlo

	// Extraer la extensión del archivo
	partes := strings.Split(handler.Filename, ".")
	if len(partes) < 2 {
		http.Error(response, "Formato de archivo no válido", http.StatusBadRequest)
		return
	}
	extension := partes[len(partes)-1]

	// Crear nombre único basado en timestamp
	nuevoNombre := fmt.Sprintf("%d.%s", time.Now().UnixNano(), extension)
	rutaArchivo := fmt.Sprintf("static/uploads/fotos/%s", nuevoNombre)

	// Crear carpeta si no existe
	os.MkdirAll("static/uploads/fotos", os.ModePerm)

	// Crear archivo en disco
	destino, err := os.Create(rutaArchivo)
	if err != nil {
		http.Error(response, "Error al guardar la imagen", http.StatusInternalServerError)
		return
	}
	defer destino.Close()

	// Copiar contenido
	_, err = io.Copy(destino, file)
	if err != nil {
		http.Error(response, "Error al copiar la imagen", http.StatusInternalServerError)
		return
	}

	// Redirigir a la página de carga con mensaje de éxito
	http.Redirect(response, request, "/formulario/upload", http.StatusSeeOther)
}
