package main

import (
	"albaranes/db"
	"albaranes/handlers"
	"albaranes/sqlc/generated"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("🔄 Iniciando aplicación...")

	// Inicializar la base de datos
	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatalf("❌ Error al inicializar la base de datos: %v", err)
	}
	defer db.CloseDB(dbConn)

	log.Println("✅ Conexión a la base de datos establecida")

	// Crear las queries
	queries := generated.New(dbConn)

	// Crear el handler de usuario
	usuarioHandler := handlers.NewUsuarioHandler(queries)

	// Configurar el router
	r := mux.NewRouter()
	r.HandleFunc("/usuarios", usuarioHandler.CreateUsuario).Methods("POST")
	r.HandleFunc("/usuarios", usuarioHandler.ListUsuarios).Methods("GET")
	r.HandleFunc("/usuarios/{usuario}", usuarioHandler.GetUsuario).Methods("GET")
	r.HandleFunc("/usuarios/{usuario}", usuarioHandler.UpdateUsuario).Methods("PUT")
	r.HandleFunc("/usuarios/{usuario}", usuarioHandler.DeleteUsuario).Methods("DELETE")

	// Iniciar el servidor
	log.Println("🚀 Servidor iniciado en :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("❌ Error al iniciar el servidor: %v", err)
	}
}
