package handlers

import (
	"clientes/conexion"
	"clientes/models"
	"database/sql"
	"fmt"
	"log"
)

// AllClientes obtiene todos los clientes ordenados por ID descendente
func AllClientes() ([]models.Cliente, error) {
	db, err := conexion.ConnectDB(conexion.CreateDSN())
	if err != nil {
		log.Printf("Error conectando a la base de datos: %v", err)
		return nil, err
	}
	defer db.Close()

	query := "SELECT id, nombre, correo, telefono, fecha FROM clientes ORDER BY id ASC"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error ejecutando la consulta: %v", err)
		return nil, err
	}
	defer rows.Close()

	var clientes []models.Cliente
	for rows.Next() {
		var c models.Cliente
		err := rows.Scan(&c.ID, &c.Nombre, &c.Correo, &c.Telefono, &c.Fecha)
		if err != nil {
			log.Printf("Error escaneando fila: %v", err)
			return nil, err
		}
		clientes = append(clientes, c)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error en las filas: %v", err)
		return nil, err
	}

	return clientes, nil
}

// ByIdCliente obtiene un cliente por su ID
func ByIdCliente(id int) (*models.Cliente, error) {
	db, err := conexion.ConnectDB(conexion.CreateDSN())
	if err != nil {
		log.Printf("Error conectando a la base de datos: %v", err)
		return nil, err
	}
	defer db.Close()

	query := "SELECT id, nombre, correo, telefono, fecha FROM clientes WHERE id = ?"
	row := db.QueryRow(query, id)

	var c models.Cliente
	err = row.Scan(&c.ID, &c.Nombre, &c.Correo, &c.Telefono, &c.Fecha)
	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontró ningún cliente con el ID proporcionado
			return nil, fmt.Errorf("cliente con ID %d no encontrado", id)
		}
		log.Printf("Error escaneando fila: %v", err)
		return nil, err
	}

	return &c, nil
}

// Crear Cliente
func CrearCliente(nombre, correo, telefono, fecha string) (*models.Cliente, error) {
	db, err := conexion.ConnectDB(conexion.CreateDSN())
	if err != nil {
		log.Printf("Error conectando a la base de datos: %v", err)
		return nil, err
	}
	defer db.Close()

	// Query para insertar el cliente
	insertQuery := `
		INSERT INTO clientes (nombre, correo, telefono, fecha)
		VALUES (?, ?, ?, ?)
	`
	result, err := db.Exec(insertQuery, nombre, correo, telefono, fecha)
	if err != nil {
		log.Printf("Error insertando cliente: %v", err)
		return nil, err
	}

	// Obtener el ID del cliente recién creado
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error obteniendo el ID del cliente: %v", err)
		return nil, err
	}

	// Query para obtener el cliente recién creado
	selectQuery := `
		SELECT id, nombre, correo, telefono, fecha
		FROM clientes
		WHERE id = ?
	`
	row := db.QueryRow(selectQuery, id)

	var c models.Cliente
	err = row.Scan(&c.ID, &c.Nombre, &c.Correo, &c.Telefono, &c.Fecha)
	if err != nil {
		log.Printf("Error escaneando fila: %v", err)
		return nil, err
	}

	return &c, nil
}

// Modifica cliente por su ID
func UpdateCliente(id int, nombre, correo, telefono, fecha string) (*models.Cliente, error) {
	db, err := conexion.ConnectDB(conexion.CreateDSN())
	if err != nil {
		log.Printf("Error conectando a la base de datos: %v", err)
		return nil, err
	}
	defer db.Close()

	// Query para actualizar el cliente
	updateQuery := `
		UPDATE clientes
		SET nombre = ?, correo = ?, telefono = ?, fecha = ?
		WHERE id = ?
	`
	_, err = db.Exec(updateQuery, nombre, correo, telefono, fecha, id)
	if err != nil {
		log.Printf("Error actualizando cliente: %v", err)
		return nil, err
	}

	// Query para obtener el cliente actualizado
	selectQuery := `
		SELECT id, nombre, correo, telefono, fecha
		FROM clientes
		WHERE id = ?
	`
	row := db.QueryRow(selectQuery, id)

	var c models.Cliente
	err = row.Scan(&c.ID, &c.Nombre, &c.Correo, &c.Telefono, &c.Fecha)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cliente con ID %d no encontrado", id)
		}
		log.Printf("Error escaneando fila: %v", err)
		return nil, err
	}

	return &c, nil
}

// Eliminar Cliente por Id
func DeleteCliente(id int) error {
	db, err := conexion.ConnectDB(conexion.CreateDSN())
	if err != nil {
		log.Printf("Error conectando a la base de datos: %v", err)
		return err
	}
	defer db.Close()

	// Query para eliminar el cliente
	query := "DELETE FROM clientes WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error eliminando cliente: %v", err)
		return err
	}

	// Verificar si se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error obteniendo filas afectadas: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("cliente con ID %d no encontrado", id)
	}

	return nil
}

// Mostrar Cliente
func MostrarCliente(c models.Cliente) {
	// Diseño ASCII superior
	fmt.Println("╔══════════════════════════════════════════════════╗")
	fmt.Println("║                     CLIENTE                      ║")
	fmt.Println("╠══════════════════════════════════════════════════╣")

	// Información del cliente
	fmt.Printf("║ ID: %-45d ║\n", c.ID)
	fmt.Printf("║ Nombre: %-40s ║\n", c.Nombre)
	fmt.Printf("║ Correo: %-40s ║\n", c.Correo)
	fmt.Printf("║ Teléfono: %-38s ║\n", c.Telefono)
	fmt.Printf("║ Fecha: %-40s ║\n", c.Fecha)

	// Diseño ASCII inferior
	fmt.Println("╚══════════════════════════════════════════════════╝")
	fmt.Println() // Espacio entre clientes
}
