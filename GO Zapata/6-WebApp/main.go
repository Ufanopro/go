package main

//go install github.com/air-verse/air@latest

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
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "127.0.0.1" // Valor por defecto
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Valor por defecto
	}
	addr := fmt.Sprintf("%s:%s", host, port)

	// Configurar el enrutador
	r := mux.NewRouter()                                     // Correcta inicialización del enrutador
	r.HandleFunc("/", router.Home)                           // raiz APP
	r.HandleFunc("/about", router.About)                     // About us
	r.HandleFunc("/url/{param1:.*}/{param2:.*}", router.URL) // datos atraves URL
	r.HandleFunc("/params/", router.Parametros)              // Ruta con parámetros de query string
	r.HandleFunc("/estructuras", router.Estructuras)         // Estructuras

	// Servir archivos estáticos
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./static/"))))

	// Crear y configurar el servidor
	server := &http.Server{
		Addr:         addr,
		Handler:      r, // Asignar el enrutador
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Iniciar el servidor
	fmt.Printf("Servidor corriendo en http://%s\n", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}

// MODO BASICO
/*
func main() {
	// Cargar variables desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando valores por defecto")
	}

	// Obtener host y puerto de las variables de entorno
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "127.0.0.1" // Valor por defecto
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Valor por defecto
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	// Configurar el enrutador
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	fmt.Printf("Servidor corriendo en http://%s\n", addr)

	// Iniciar el servidor
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}
*/
