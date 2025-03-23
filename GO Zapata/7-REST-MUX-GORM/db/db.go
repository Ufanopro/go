package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

// InitDB carga las variables de entorno y establece la conexión con la BD
func InitDB() {
	once.Do(func() {
		// Carga las variables del archivo .env
		if err := godotenv.Load(); err != nil {
			log.Println("Advertencia: No se pudo cargar el archivo .env, se usarán variables de entorno")
		}

		// Obtiene las credenciales de la BD
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		database := os.Getenv("DB_NAME")

		// Verifica que no haya valores vacíos
		if user == "" || password == "" || host == "" || port == "" || database == "" {
			log.Fatal("Error: Faltan variables de entorno para la conexión a la BD")
		}

		// Construcción del DSN (Data Source Name)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, database)

		// Intenta abrir la conexión
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Error de conexión a la BD:", err)
		}

		log.Println("✅ Conexión a la BD exitosa")
		dbInstance = db
	})
}

// GetDB devuelve la instancia de la BD
func GetDB() *gorm.DB {
	if dbInstance == nil {
		InitDB()
	}
	return dbInstance
}
