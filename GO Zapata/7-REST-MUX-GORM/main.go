package main

import (
	"fmt"
	"log"
	"net/http"

	"restapi/db"
	"restapi/handlers"
	"restapi/logic"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	//Conexion con BD
	// Inicializa la base de datos
	db.InitDB()

	// Obtiene la instancia y la usa
	database := db.GetDB()

	// Verifica conexión con Ping
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("Error obteniendo conexión de base de datos:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("No se pudo hacer ping a la BD:", err)
	}

	fmt.Println("🎉 Conexión a la BD verificada con éxito")

	mux := mux.NewRouter()
	prefijo := "/api/v1/"
	mux.HandleFunc(prefijo+"ejemplo", handlers.EjemploGET).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+}", handlers.EjemploParams).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo", handlers.EjemploPOST).Methods("POST")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+}", handlers.EjemploPUT).Methods("PUT")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+}", handlers.EjemploDEL).Methods("DELETE")
	mux.HandleFunc(prefijo+"query", handlers.EjemploQuery).Methods("GET")
	mux.HandleFunc(prefijo+"upload", handlers.EjemploUpload).Methods("POST")
	mux.HandleFunc(prefijo+"ver", handlers.VerFoto).Methods("GET")

	//GORM
	mux.HandleFunc(prefijo+"categorias", handlers.CategoriasGET).Methods("GET")
	mux.HandleFunc(prefijo+"categorias/{id:[0-9]+}", handlers.CategoriasGETID).Methods("GET")
	mux.HandleFunc(prefijo+"categorias", handlers.CategoriasPOST).Methods("POST")
	mux.HandleFunc(prefijo+"categorias/{id:[0-9]+}", handlers.CategoriasPUT).Methods(http.MethodPut)
	mux.HandleFunc(prefijo+"categorias/{id:[0-9]+}", handlers.CategoriasDEL).Methods(http.MethodDelete)

	// Rutas para Productos
	mux.HandleFunc(prefijo+"productos", handlers.ProductosGET).Methods("GET")
	mux.HandleFunc(prefijo+"productos/{id:[0-9]+}", handlers.ProductosGETID).Methods("GET")
	mux.HandleFunc(prefijo+"productos", handlers.ProductosPOST).Methods("POST")
	mux.HandleFunc(prefijo+"productos/{id:[0-9]+}", handlers.ProductosPUT).Methods(http.MethodPut)
	mux.HandleFunc(prefijo+"productos/{id:[0-9]+}", handlers.ProductosDEL).Methods(http.MethodDelete)

	//Rutas para fotos
	mux.HandleFunc(prefijo+"fotos", handlers.FotosGET).Methods("GET")
	mux.HandleFunc(prefijo+"fotos/referencia/{id:[0-9]+}", handlers.FotosGETByReferenciaID).Methods("GET")
	mux.HandleFunc(prefijo+"fotos/tipo/{tipo}", handlers.FotosGETByTipo).Methods("GET")
	mux.HandleFunc(prefijo+"fotos", handlers.FotosPOST).Methods("POST")
	mux.HandleFunc(prefijo+"fotos/{id:[0-9]+}", handlers.FotosPUT).Methods(http.MethodPut)
	mux.HandleFunc(prefijo+"fotos/{id:[0-9]+}", handlers.FotosDEL).Methods(http.MethodDelete)
	mux.HandleFunc(prefijo+"fotos/upload", handlers.FotosUpload).Methods("POST")
	mux.HandleFunc(prefijo+"fotos/ver", logic.ValidarJWT(handlers.FotoVer)).Methods("GET")
	mux.HandleFunc(prefijo+"fotos/ver", handlers.FotoDeLocal).Methods("DELETE")

	//Rutas Login
	mux.HandleFunc(prefijo+"seguridad/registro", handlers.LoginReg).Methods("POST")
	mux.HandleFunc(prefijo+"seguridad/registro", handlers.LoginRegPUT).Methods("PUT")
	mux.HandleFunc(prefijo+"seguridad/login", handlers.Login).Methods("POST")
	mux.HandleFunc(prefijo+"seguridad/protegido", logic.ValidarJWT(handlers.LoginSec)).Methods("GET")

	//CORS
	handler := cors.AllowAll().Handler(mux)
	//log.Fatal(http.ListenAndServe(":8084", mux))
	log.Fatal(http.ListenAndServe(":8084", handler))

}
