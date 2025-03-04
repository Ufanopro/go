package db

import (
	"fmt"
	"log"
	"os"

	"albaranes/sqlc/generated"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB

// Queries combina las consultas generadas por sqlc con la conexión a la BD
type Queries struct {
	*generated.Queries
	db *sqlx.DB
}

// NewQueries crea una nueva instancia de Queries
func NewQueries(db *sqlx.DB) *Queries {
	return &Queries{
		Queries: generated.New(db.DB), // Inicializa sqlc
		db:      db,
	}
}

// InitDB inicializa la conexión a la base de datos
func InitDB() (*sqlx.DB, error) {
	// Intentar cargar el archivo .env
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️ No se pudo cargar el archivo .env, usando valores de entorno por defecto")
	}

	// Obtener las variables de entorno
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construir la cadena de conexión DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Conectar a la base de datos
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("❌ Error conectando a la base de datos: %v", err)
	}

	log.Println("✅ Conexión exitosa a la base de datos")
	DB = db
	return db, nil
}

// CloseDB cierra la conexión a la base de datos
func CloseDB(db *sqlx.DB) {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Println("⚠️ Error al cerrar la conexión a la base de datos:", err)
		} else {
			log.Println("✅ Conexión a la base de datos cerrada correctamente")
		}
	}
}

// GetDB retorna la instancia de la base de datos
func GetDB() *sqlx.DB {
	return DB
}
