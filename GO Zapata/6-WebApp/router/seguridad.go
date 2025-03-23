package router

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"webapp/conexion"
	"webapp/models"
	"webapp/utilidades"

	"golang.org/x/crypto/bcrypt"
)

func SeguridadRegistro(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/seguridad/registro.html", utilidades.Front))
	template.Execute(response, nil)
	/*
		fmt.Fprintln(response, "Hola WebAPP")

		template, err := template.ParseFiles("templates/index/index.html", "templates/front/front.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(response, nil)
		}

	*/
}

// SeguridadRegistro_post maneja la validación del formulario de registro
func SeguridadRegistroPost(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := request.ParseForm(); err != nil {
		http.Error(response, "Error al procesar el formulario", http.StatusInternalServerError)
		return
	}

	nombre := request.FormValue("nombre")
	correo := request.FormValue("correo")
	telefono := request.FormValue("telefono")
	password := request.FormValue("password")
	confirmPassword := request.FormValue("confirm_password")

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	var errores []string

	if nombre == "" {
		errores = append(errores, "El campo 'Nombre' es obligatorio.")
	}
	if correo == "" || !emailRegex.MatchString(correo) {
		errores = append(errores, "El correo no es válido o está vacío.")
	}
	if telefono == "" {
		errores = append(errores, "El campo 'Teléfono' es obligatorio.")
	}
	if password == "" || confirmPassword == "" || password != confirmPassword {
		errores = append(errores, "Las contraseñas no coinciden o están vacías.")
	}

	if len(errores) > 0 {
		log.Println("Errores en el formulario:")
		for _, err := range errores {
			log.Println("-", err)
		}
		http.Error(response, "Error en los datos enviados", http.StatusBadRequest)
		return
	}

	// Encriptar la contraseña antes de guardarla
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(response, "Error al encriptar la contraseña", http.StatusInternalServerError)
		return
	}

	// Conectar a la base de datos
	db, err := conexion.Conectar()
	if err != nil {
		http.Error(response, "Error al conectar a la base de datos", http.StatusInternalServerError)
		log.Println("❌ Error al conectar a la base de datos:", err)
		return
	}
	defer db.Close()

	// Query corregida sin la columna "fecha"
	sql := "INSERT INTO usuarios (nombre, correo, telefono, password) VALUES (?, ?, ?, ?)"
	res, err := db.Exec(sql, nombre, correo, telefono, passwordHash)
	if err != nil {
		http.Error(response, "Error al insertar en la base de datos", http.StatusInternalServerError)
		log.Println("❌ Error en la inserción:", err)
		return
	}

	// Verificar si realmente se insertó un registro
	filasAfectadas, _ := res.RowsAffected()
	log.Println("✅ Filas afectadas:", filasAfectadas)

	// Redirigir a la página principal
	http.Redirect(response, request, "/", http.StatusSeeOther)
}

// SeguridadLogin carga la página de login
// SeguridadLogin carga la página de login
func SeguridadLogin(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/seguridad/login.html", utilidades.Frontend))
	template.Execute(response, nil)
}

// SeguridadLoginPost maneja la validación del login
func SeguridadLoginPost(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := request.ParseForm(); err != nil {
		http.Error(response, "Error al procesar el formulario", http.StatusInternalServerError)
		return
	}

	correo := request.FormValue("correo")
	password := request.FormValue("password")

	// Validaciones
	mensaje := ""

	// Validar email y contraseña
	correo := request.FormValue("correo")
	password := request.FormValue("password")

	if len(correo) == 0 {
		mensaje += " . El campo E-Mail está vacío "
	} else if validaciones.Regex_correo.FindStringSubmatch(correo) == nil {
		mensaje += " . El E-Mail ingresado no es válido "
	}

	if !validaciones.ValidarPassword(password) {
		mensaje += " . La contraseña debe tener al menos 1 número, una mayúscula y un largo entre 6 y 20 caracteres "
	}

	// Si hay errores, redirigir con mensaje de error
	if mensaje != "" {
		utilidades.CrearMensajesFlash(response, request, "danger", mensaje)
		http.Redirect(response, request, "/seguridad/login", http.StatusSeeOther)
		return
	}

	// Conectamos a la BD
	conectar.Conectar()
	defer conectar.CerrarConexion() // Asegurar que la conexión se cierre al salir de la función

	// Consultar usuario en la BD
	sqlQuery := "SELECT id, nombre, correo, telefono, password FROM usuarios WHERE correo=?"
	var dato modelos.Usuario

	err := conectar.Db.QueryRow(sqlQuery, correo).Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono, &dato.Password)
	if err == sql.ErrNoRows {
		utilidades.CrearMensajesFlash(response, request, "danger", "Las credenciales son inválidas")
		http.Redirect(response, request, "/seguridad/login", http.StatusSeeOther)
		return
	} else if err != nil {
		fmt.Println("Error al consultar usuario:", err)
		utilidades.CrearMensajesFlash(response, request, "danger", "Ocurrió un error, intenta nuevamente")
		http.Redirect(response, request, "/seguridad/login", http.StatusSeeOther)
		return
	}


// SeguridadProtegida muestra una página protegida para usuarios autenticados
func SeguridadProtegida(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/seguridad/protegida.html", utilidades.Frontend))

	// Obtener mensajes flash y datos de sesión
	cssSesion, cssMensaje := utilidades.RetornarMensajesFlash(response, request)
	tamilaID, tamilaNombre := utilidades.RetornarLogin(request)

	data := map[string]string{
		"css":           cssSesion,
		"mensaje":       cssMensaje,
		"tamila_id":     tamilaID,
		"tamila_nombre": tamilaNombre,
	}

	template.Execute(response, data)
}

// SeguridadLogout cierra la sesión del usuario
func SeguridadLogout(response http.ResponseWriter, request *http.Request) {
	session, _ := utilidades.Store.Get(request, "session-name")
	session.Values["tamila_id"] = nil
	session.Values["tamila_nombre"] = nil

	err := session.Save(request, response)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	utilidades.CrearMensajesFlash(response, request, "primary", "Se ha cerrado tu sesión exitosamente")
	http.Redirect(response, request, "/seguridad/login", http.StatusSeeOther)
}
