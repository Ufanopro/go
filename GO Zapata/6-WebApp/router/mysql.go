package router

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"webapp/conexion"
	"webapp/models"
	"webapp/utilidades"

	"github.com/gorilla/mux"
)

func AllClientes(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/mysql/home.html", utilidades.Front))
	//Conexion DB
	conexion.Conectar()
	sql := "SELECT id, nombre, correo, telefono FROM clientes"
	clientes := models.Clientes{}
	datos, err := conexion.Query(sql)
	if err != nil {
		fmt.Println((err))
	}
	defer conexion.CerrarDB()
	for datos.Next() {
		dato := models.Cliente{}
		datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		clientes = append(clientes, dato)
	}

	/*
		fmt.Fprintln(response, "Hola WebAPP")

		template, err := template.ParseFiles("templates/index/index.html", "templates/front/front.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(response, nil)
		}

	*/
	// Renderizar plantilla con los datos
	err = template.Execute(response, map[string]interface{}{
		"Clientes": clientes,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(response, "Error al renderizar la plantilla", http.StatusInternalServerError)
	}
}

func ClienteView(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/mysql/cliente.html", utilidades.Front))
	template.Execute(response, nil)
}

func ClientePost(response http.ResponseWriter, request *http.Request) {
	fmt.Println("🔍 ClientePost ejecutándose...") // Verificación

	if request.Method != http.MethodPost {
		http.Error(response, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtener valores del formulario
	nombre := request.FormValue("nombre")
	correo := request.FormValue("correo")
	telefono := request.FormValue("telefono")

	fmt.Println("📋 Datos recibidos:", nombre, correo, telefono) // Verificación

	// Validación básica
	if nombre == "" || correo == "" || telefono == "" {
		http.Error(response, "Todos los campos son obligatorios", http.StatusBadRequest)
		fmt.Println("⚠️ Datos inválidos")
		return
	}

	// Conectar a la base de datos
	db, err := conexion.Conectar()
	if err != nil {
		http.Error(response, "Error al conectar a la base de datos", http.StatusInternalServerError)
		fmt.Println("❌ Error al conectar a la base de datos:", err)
		return
	}
	//defer db.Close() // Cerramos la conexión solo para esta solicitud

	// Query de inserción
	sql := "INSERT INTO clientes (nombre, correo, telefono) VALUES (?, ?, ?)"
	res, err := db.Exec(sql, nombre, correo, telefono)
	if err != nil {
		http.Error(response, "Error al insertar en la base de datos", http.StatusInternalServerError)
		fmt.Println("❌ Error en la inserción:", err)
		return
	}

	// Verificar si realmente se insertó un registro
	filasAfectadas, _ := res.RowsAffected()
	fmt.Println("✅ Filas afectadas:", filasAfectadas)

	// Redirigir a la lista de clientes
	http.Redirect(response, request, "/mysql", http.StatusSeeOther)
}

func ClienteActualizarView(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/mysql/editar.html", utilidades.Front))

	// Obtener ID del cliente desde la URL
	vars := mux.Vars(request)
	id := vars["id"]

	// Conectar a la base de datos
	db, err := conexion.Conectar()
	if err != nil {
		http.Error(response, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Buscar datos del cliente
	var cliente models.Cliente
	query := "SELECT id, nombre, correo, telefono FROM clientes WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&cliente.Id, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)

	if err != nil {
		http.Error(response, "Cliente no encontrado", http.StatusNotFound)
		return
	}

	// Renderizar la plantilla con los datos del cliente
	tmpl.Execute(response, cliente)
}

func ClienteActualizar(response http.ResponseWriter, request *http.Request) {
	log.Println("Iniciando actualización de cliente...")

	if request.Method != http.MethodPost {
		log.Println("Método no permitido:", request.Method)
		http.Error(response, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtener ID desde la URL
	vars := mux.Vars(request)
	id := vars["id"]
	log.Println("ID del cliente a actualizar:", id)

	// Parsear los datos del formulario
	err := request.ParseForm()
	if err != nil {
		log.Println("Error al procesar el formulario:", err)
		http.Error(response, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	nombre := request.FormValue("nombre")
	correo := request.FormValue("correo")
	telefono := request.FormValue("telefono")

	log.Println("Datos recibidos - Nombre:", nombre, ", Correo:", correo, ", Teléfono:", telefono)

	// Conectar a la base de datos
	db, err := conexion.Conectar()
	if err != nil {
		log.Println("Error al conectar con la base de datos:", err)
		http.Error(response, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}

	log.Println("Conexión a la base de datos establecida correctamente.")

	// Verificar si la conexión sigue abierta
	if err = db.Ping(); err != nil {
		log.Println("La conexión a la base de datos se cerró inesperadamente:", err)
		http.Error(response, "Error con la base de datos", http.StatusInternalServerError)
		return
	}

	// Actualizar datos del cliente
	query := "UPDATE clientes SET nombre = ?, correo = ?, telefono = ? WHERE id = ?"
	result, err := db.Exec(query, nombre, correo, telefono, id)
	if err != nil {
		log.Println("Error al actualizar el cliente:", err)
		http.Error(response, "Error al actualizar el cliente", http.StatusInternalServerError)
		return
	}

	// Verificar si se actualizó algún registro
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("No se encontró el cliente con ID:", id)
		http.Error(response, "Cliente no encontrado", http.StatusNotFound)
		return
	}

	log.Println("Cliente actualizado correctamente, ID:", id)

	// Redirigir a la lista de clientes
	http.Redirect(response, request, "/mysql", http.StatusSeeOther)
}

// DELETE
func ClienteDelView(response http.ResponseWriter, request *http.Request) {
	log.Println("ClienteDelView: Solicitud recibida") // Log informativo

	tmpl := template.Must(template.ParseFiles("templates/mysql/borrar.html", utilidades.Front))

	vars := mux.Vars(request)
	id, ok := vars["id"]
	if !ok {
		http.Error(response, "ID no proporcionado", http.StatusBadRequest)
		return
	}

	// Conectar a la base de datos
	db, err := conexion.Conectar()
	if err != nil {
		log.Println("Error al conectar con la base de datos")
		http.Error(response, "Error interno", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var cliente models.Cliente
	query := "SELECT id, nombre, correo, telefono FROM clientes WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&cliente.Id, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)
	if err != nil {
		http.Error(response, "Cliente no encontrado", http.StatusNotFound)
		return
	}

	// Renderizar la plantilla con los datos del cliente
	tmpl.Execute(response, cliente)
}

func ClienteDel(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(request)
	id := vars["id"]

	db, err := conexion.Conectar()
	if err != nil {
		log.Println("Error al conectar con la base de datos")
		http.Error(response, "Error interno", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := "DELETE FROM clientes WHERE id = ?"
	_, err = db.Exec(query, id)
	if err != nil {
		http.Error(response, "Error al eliminar el cliente", http.StatusInternalServerError)
		return
	}

	log.Println("Cliente eliminado con ID:", id) // Log de éxito
	http.Redirect(response, request, "/mysql", http.StatusSeeOther)
}
