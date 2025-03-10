package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"webapp/router"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando valores por defecto")
	}

	// Obtener host y puerto de las variables de entorno
	host := getEnv("SERVER_HOST", "127.0.0.1")
	port := getEnv("SERVER_PORT", "8080")
	addr := fmt.Sprintf("%s:%s", host, port)

	// Configurar el enrutador
	r := mux.NewRouter()
	r.HandleFunc("/", router.Home)
	r.HandleFunc("/about", router.About)
	r.HandleFunc("/url/{param1}/{param2}", router.URL)
	r.HandleFunc("/params/", router.Parametros)
	r.HandleFunc("/estructuras", router.Estructuras)
	r.HandleFunc("/formulario", router.GetFormulario)
	r.HandleFunc("/formularios-post", router.PostFormulario).Methods("POST")

	// Servir archivos estáticos
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./static/"))))

	// Configurar error 404
	r.NotFoundHandler = http.HandlerFunc(router.Pagina404)

	// Crear y configurar el servidor
	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Iniciar el servidor
	fmt.Printf("Servidor corriendo en http://%s\n", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}

// getEnv obtiene una variable de entorno o devuelve un valor por defecto
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
