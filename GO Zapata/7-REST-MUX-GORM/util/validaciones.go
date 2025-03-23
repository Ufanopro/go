package util

import (
	"regexp"
	"unicode"
)

var RegexCorreo = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// ValidarCorreo verifica si el correo tiene un formato válido.
func ValidarCorreo(correo string) bool {
	return RegexCorreo.MatchString(correo)
}

// ValidarPassword verifica que la contraseña tenga al menos 6 caracteres, máximo 20, una mayúscula y un número
func ValidarPassword(password string) bool {
	// Verifica la longitud
	if len(password) < 6 || len(password) > 20 {
		return false
	}

	// Bandera para mayúscula y número
	hasUpper := false
	hasNumber := false

	// Recorrer la contraseña para verificar condiciones
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasNumber = true
		}
	}

	// Retorna true solo si cumple ambas condiciones
	return hasUpper && hasNumber
}
