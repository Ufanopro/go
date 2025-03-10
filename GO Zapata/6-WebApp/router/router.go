package router

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"webapp/utilidades"

	"github.com/gorilla/mux"
)

const (
	param1 = "p1"
	param2 = "p2"
	param3 = "p3"
)

func Home(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/index/index.html", utilidades.Front))
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

func About(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/about/about.html", utilidades.Front))
	template.Execute(response, nil)
	/*
		template, err := template.ParseFiles("templates/about/about.html", "templates/front/front.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(response, nil)
		}
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
	tmpl, err := template.ParseFiles("templates/url/url.html", utilidades.Front)
	if err != nil {
		http.Error(response, "Error cargando la plantilla", http.StatusInternalServerError)
		return
	}

	// Datos a pasar a la plantilla
	data := map[string]string{
		"param1": param1,
		"param2": param2,
		"Texto":  "Este texto es harcodeado",
	}

	// Renderizar la plantilla con los valores dinámicos
	err = tmpl.Execute(response, data)
	if err != nil {
		http.Error(response, "Error renderizando la plantilla", http.StatusInternalServerError)
	}
}

// Ruta para manejar los parámetros de la query string
func Parametros(response http.ResponseWriter, request *http.Request) {
	// Parsear la plantilla HTML
	tmpl, err := template.ParseFiles("templates/params/params.html", utilidades.Front)
	if err != nil {
		http.Error(response, "Error cargando la plantilla", http.StatusInternalServerError)
		log.Println("Error al cargar la plantilla:", err)
		return
	}

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
	habilidad1 := Habilidad{"manualidades"}
	habilidad2 := Habilidad{"programación"}
	habilidad3 := Habilidad{"idiomas"}
	habilidades := []Habilidad{habilidad1, habilidad2, habilidad3}
	template := template.Must(template.ParseFiles("templates/estructuras/estructuras.html", utilidades.Front))
	template.Execute(response, Datos{"César", 45, 1, habilidades})
	/*
		template, err := template.ParseFiles("templates/estructuras/estructuras.html", utilidades.Front)
		if err != nil {
			panic(err)
		} else {
			template.Execute(response, Datos{"César", 45, 1, habilidades})
		}
	*/
}

func Pagina404(response http.ResponseWriter, request *http.Request) {
	// Configurar el código de estado HTTP 404
	response.WriteHeader(http.StatusNotFound)

	// Intentar cargar la plantilla
	template := template.Must(template.ParseFiles("templates/404/404.html", utilidades.Front))
	template.Execute(response, nil)
	/*
			tmpl, err := template.ParseFiles("templates/404/404.html")
		    if err != nil {
		        // En caso de error, mostrar un mensaje simple
		        http.Error(response, "Página no encontrada", http.StatusNotFound)
		        return
		    }

		    // Renderizar la plantilla
		    err = tmpl.Execute(response, nil)
		    if err != nil {
		        http.Error(response, "Error renderizando la página 404", http.StatusInternalServerError)
		    }
	*/
}

/*

eq	==	Igual a	a == b → false
ne	!=	Diferente de	a != b → true
lt	<	Menor que	a < b → true
le	<=	Menor o igual que	a <= b → true
gt	>	Mayor que	a > b → false
ge	>=	Mayor o igual que	a >= b → false

*/
