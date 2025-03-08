package router

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	param1 = "p1"
	param2 = "p2"
	param3 = "p3"
)

func Home(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/index/index.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
	/*
		template := template.Must(template.ParseFiles("templates/ejemplo/home.html", utilidades.Frontend))
		template.Execute(response, nil)

		fmt.Fprintln(response, "Hola WebAPP")
	*/
}

func About(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/about/about.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
	/*
		fmt.Fprintln(response, "About Us TakeUS")
	*/
}

// RUTA con URL
func URL(response http.ResponseWriter, request *http.Request) {
	// Capturar los parámetros dinámicos desde la URL
	vars := mux.Vars(request)
	param1 := vars["param1"] // Captura el primer parámetro
	param2 := vars["param2"] // Captura el segundo parámetro
	// Si param1 o param2 están vacíos, mostrar un mensaje
	if param1 == "" {
		param1 = "param1 no proporcionado"
	}
	if param2 == "" {
		param2 = "param2 no proporcionado"
	}

	// Imprimir los valores de los parámetros en la consola
	fmt.Printf("Valor de param1: %s\n", param1)
	fmt.Printf("Valor de param2: %s\n", param2)

	// Parsear la plantilla HTML
	tmpl, err := template.ParseFiles("templates/url/url.html")
	if err != nil {
		http.Error(response, "Error cargando la plantilla", http.StatusInternalServerError)
		return
	}

	// Datos a pasar a la plantilla
	data := map[string]string{
		"param1": param1,
		"param2": param2,
	}

	// Renderizar la plantilla con los valores dinámicos
	err = tmpl.Execute(response, data)
	if err != nil {
		http.Error(response, "Error renderizando la plantilla", http.StatusInternalServerError)
	}
}

// Ruta para manejar los parámetros de la query string
func Parametros(response http.ResponseWriter, request *http.Request) {
	// Obtener los parámetros de la query string
	queryParams := request.URL.Query()

	// Mapa para almacenar los valores de los parámetros
	data := map[string]string{
		"param1": queryParams.Get(param1),
		"param2": queryParams.Get(param2),
		"param3": queryParams.Get(param3),
	}

	// Imprimir los parámetros en la consola
	log.Println("Parámetros recibidos:", data)

	// Parsear la plantilla HTML
	tmpl, err := template.ParseFiles("templates/params/params.html")
	if err != nil {
		http.Error(response, "Error cargando la plantilla", http.StatusInternalServerError)
		log.Println("Error al cargar la plantilla:", err)
		return
	}

	// Renderizar la plantilla con los valores obtenidos
	err = tmpl.Execute(response, data)
	if err != nil {
		http.Error(response, "Error renderizando la plantilla", http.StatusInternalServerError)
		log.Println("Error al renderizar la plantilla:", err)
	}
}

type Habilidad struct {
	Nombre string
}

type Datos struct {
	Nombre      string
	Edad        int
	Perfil      int
	Habilidades []Habilidad
}

func Estructuras(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/estructuras/estructuras.html")
	habilidad1 := Habilidad{"manualidades"}
	habilidad2 := Habilidad{"programación"}
	habilidad3 := Habilidad{"idiomas"}
	habilidades := []Habilidad{habilidad1, habilidad2, habilidad3}
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, Datos{"César", 45, 1, habilidades})
	}
}

/*

eq	==	Igual a	a == b → false
ne	!=	Diferente de	a != b → true
lt	<	Menor que	a < b → true
le	<=	Menor o igual que	a <= b → true
gt	>	Mayor que	a > b → false
ge	>=	Mayor o igual que	a >= b → false

*/
