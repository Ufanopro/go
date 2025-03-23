package conexion

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

// Conectar establece la conexión a la base de datos y la asigna a la variable global db.
func Conectar() (*sql.DB, error) {
	if db != nil {
		// Verificar si la conexión sigue abierta
		if err := db.Ping(); err == nil {
			return db, nil // La conexión sigue activa, se reutiliza
		}
		// Si la conexión está cerrada, reinicializar
		db = nil
	}

	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("⚠️ Error al cargar el archivo .env")
	}

	// Crear la cadena de conexión
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Conectar a la base de datos
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("❌ Error al conectar a la base de datos: %v", err)
	}

	// Verificar la conexión
	if err = db.Ping(); err != nil {
		db.Close()
		db = nil
		return nil, fmt.Errorf("❌ No se pudo conectar a la base de datos: %v", err)
	}

	return db, nil
}

// CerrarDB cierra la conexión a la base de datos si está abierta.
func CerrarDB() error {
	if db != nil {
		err := db.Close()
		db = nil
		return err
	}
	return nil
}

// Query ejecuta una consulta SQL y devuelve los resultados.
func Query(sqlQuery string, args ...interface{}) (*sql.Rows, error) {
	if db == nil {
		if _, err := Conectar(); err != nil {
			return nil, err
		}
	}
	return db.Query(sqlQuery, args...)
}
