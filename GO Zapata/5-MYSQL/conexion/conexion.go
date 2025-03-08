package conexion

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Cargar las variables de entorno desde el archivo .env
func LoadEnvVariables() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("⚠️ Error al cargar el archivo .env")
	}
	return nil
}

// Crear la cadena de conexión a la base de datos
func CreateDSN() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}

// Conectar a la base de datos
func ConnectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("❌ Error al conectar a la base de datos: %v", err)
	}
	return db, nil
}

// Verificar la conexión
func PingDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("❌ No se pudo conectar a la base de datos: %v", err)
	}
	return nil
}
