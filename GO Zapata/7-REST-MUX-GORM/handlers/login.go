package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi/db"
	"restapi/dto"
	"restapi/util"

	modelos "restapi/models"
	validaciones "restapi/util"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	log.Println("Iniciando proceso de login")

	var registro dto.LoginDto
	if err := json.NewDecoder(request.Body).Decode(&registro); err != nil {
		log.Printf("Error al decodificar JSON: %v", err)
		enviarRespuesta(response, http.StatusBadRequest, "Ocurrió un error inesperado")
		return
	}

	log.Printf("Correo recibido: %s", registro.Correo)

	if !validarCorreo(registro.Correo, response) {
		log.Println("Error en la validación del correo")
		return
	}

	if !validarPassword(registro.Password, response) {
		log.Println("Error en la validación de la contraseña")
		return
	}

	var usuario modelos.Usuario
	if !buscarUsuarioPorCorreo(registro.Correo, &usuario, response) {
		log.Println("Usuario no encontrado en la base de datos")
		return
	}

	log.Println("Usuario encontrado, verificando contraseña")

	if !compararPassword(usuario.Password, registro.Password) {
		log.Println("Contraseña incorrecta")
		enviarRespuesta(response, http.StatusUnauthorized, "Las credenciales ingresadas son inválidas")
		return
	}

	log.Println("Contraseña correcta, generando token")

	token, err := util.GenerarJWT(usuario)
	if err != nil {
		log.Printf("Error al generar el token JWT: %v", err)
		enviarRespuesta(response, http.StatusInternalServerError, "Error al generar el token")
		return
	}

	log.Println("Token generado correctamente, enviando respuesta")

	enviarRespuestaExitoso(response, usuario.Nombre, token)
}

func validarCorreo(correo string, response http.ResponseWriter) bool {
	if correo == "" {
		log.Println("Correo vacío")
		enviarRespuesta(response, http.StatusBadRequest, "El E-Mail es obligatorio")
		return false
	}
	if validaciones.RegexCorreo.FindStringSubmatch(correo) == nil {
		log.Printf("Correo inválido: %s", correo)
		enviarRespuesta(response, http.StatusBadRequest, "El E-Mail ingresado no es válido")
		return false
	}
	log.Println("Correo válido")
	return true
}

func validarPassword(password string, response http.ResponseWriter) bool {
	if !validaciones.ValidarPassword(password) {
		log.Println("Contraseña no cumple con los requisitos")
		enviarRespuesta(response, http.StatusBadRequest, "La contraseña debe tener al menos 1 número, una mayúscula y un largo entre 6 y 20 caracteres")
		return false
	}
	log.Println("Contraseña válida")
	return true
}

func buscarUsuarioPorCorreo(correo string, usuario *modelos.Usuario, response http.ResponseWriter) bool {
	log.Printf("Buscando usuario con correo: %s", correo)
	if db.GetDB().Where("correo = ?", correo).First(usuario).RowsAffected == 0 {
		log.Println("Usuario no encontrado")
		enviarRespuesta(response, http.StatusUnauthorized, "Las credenciales ingresadas son inválidas")
		return false
	}
	log.Println("Usuario encontrado")
	return true
}

func compararPassword(hashPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
	if err != nil {
		log.Println("Error al comparar contraseñas:", err)
		return false
	}
	log.Println("Contraseña correcta")
	return true
}

func enviarRespuesta(response http.ResponseWriter, statusCode int, mensaje string) {
	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(map[string]string{
		"estado":  "error",
		"mensaje": mensaje,
	})
}

func enviarRespuestaExitoso(response http.ResponseWriter, nombre, token string) {
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(dto.LoginRespuestaDto{
		Nombre: nombre,
		Token:  token,
	})
}

func LoginReg(response http.ResponseWriter, request *http.Request) {
	var registro dto.UsuarioDto
	response.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(request.Body).Decode(&registro); err != nil {
		enviarRespuesta(response, http.StatusBadRequest, "Ocurrió un error inesperado")
		return
	}

	if err := validarUsuario(registro); err != "" {
		enviarRespuesta(response, http.StatusBadRequest, err)
		return
	}

	dbInstance := db.GetDB()
	var usuario modelos.Usuario

	if err := dbInstance.Where("correo = ?", registro.Correo).First(&usuario).Error; err == nil {
		enviarRespuesta(response, http.StatusBadRequest, "El E-Mail "+registro.Correo+" ya está siendo usado por otro usuario")
		return
	} else if err != gorm.ErrRecordNotFound {
		enviarRespuesta(response, http.StatusInternalServerError, "Error al verificar el usuario")
		return
	}

	// Crear usuario
	if err := crearUsuario(dbInstance, registro); err != nil {
		enviarRespuesta(response, http.StatusInternalServerError, "Error al crear el usuario")
		return
	}
	enviarRespuesta(response, http.StatusCreated, "Se creó el registro exitosamente")
}

func LoginRegPUT(response http.ResponseWriter, request *http.Request) {
	var registro dto.UsuarioDto
	response.Header().Set("Content-Type", "application/json")

	// Decodificar el JSON del request
	if err := json.NewDecoder(request.Body).Decode(&registro); err != nil {
		enviarRespuesta(response, http.StatusBadRequest, "Ocurrió un error inesperado")
		return
	}

	// Validar usuario (excepto la contraseña si no se envía)
	if len(registro.Password) > 0 {
		if err := validarUsuario(registro); err != "" {
			enviarRespuesta(response, http.StatusBadRequest, err)
			return
		}
	}

	dbInstance := db.GetDB()
	var usuario modelos.Usuario

	// Buscar usuario por correo
	if err := dbInstance.Where("correo = ?", registro.Correo).First(&usuario).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			enviarRespuesta(response, http.StatusNotFound, "El usuario no existe")
		} else {
			enviarRespuesta(response, http.StatusInternalServerError, "Error al buscar el usuario")
		}
		return
	}

	// Actualizar los datos del usuario
	usuario.Nombre = registro.Nombre
	usuario.Telefono = registro.Telefono
	usuario.PerfilID = registro.PerfilID

	// Si se proporciona una nueva contraseña, se actualiza
	if len(registro.Password) > 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(registro.Password), bcrypt.DefaultCost)
		if err != nil {
			enviarRespuesta(response, http.StatusInternalServerError, "Error al cifrar la nueva contraseña")
			return
		}
		usuario.Password = string(hash)
	}

	// Guardar cambios en la base de datos (forzar actualización de todos los campos)
	if err := dbInstance.Model(&usuario).Updates(usuario).Error; err != nil {
		enviarRespuesta(response, http.StatusInternalServerError, "Error al actualizar el usuario")
		return
	}

	enviarRespuesta(response, http.StatusOK, "Usuario actualizado correctamente")
}

func validarUsuario(registro dto.UsuarioDto) string {
	if len(registro.Nombre) == 0 {
		return "El Nombre es obligatorio"
	}
	if len(registro.Correo) == 0 {
		return "El E-Mail es obligatorio"
	}
	if validaciones.RegexCorreo.FindStringSubmatch(registro.Correo) == nil {
		return "El E-Mail ingresado no es válido"
	}
	if !validaciones.ValidarPassword(registro.Password) {
		return "La contraseña debe tener al menos 1 número, una mayúscula y un largo entre 6 y 20 caracteres"
	}
	return ""
}

func crearUsuario(dbInstance *gorm.DB, registro dto.UsuarioDto) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(registro.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usuario := modelos.Usuario{
		Nombre:   registro.Nombre,
		Correo:   registro.Correo,
		Telefono: registro.Telefono,
		PerfilID: registro.PerfilID,
		Password: string(hash),
		Fecha:    time.Now(),
	}
	return dbInstance.Create(&usuario).Error
}

// Rutas protegidas
func LoginSec(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Recurso protegido",
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(respuesta)
}
