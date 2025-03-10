package utilidades

import (
	"regexp"
)

// ValidateEmail verifica si el correo es válido y devuelve un mensaje si hay error
func ValidateEmail(email string) (bool, string) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return false, "El correo electrónico no es válido."
	}
	return true, ""
}

// ValidatePassword verifica si la contraseña cumple los requisitos y devuelve un mensaje de error
func ValidatePassword(password string) (bool, string) {
	if len(password) < 8 {
		return false, "La contraseña debe tener al menos 8 caracteres."
	}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[@$!%*?&]`).MatchString(password)

	if !hasUpper {
		return false, "La contraseña debe contener al menos una letra mayúscula."
	}
	if !hasLower {
		return false, "La contraseña debe contener al menos una letra minúscula."
	}
	if !hasDigit {
		return false, "La contraseña debe contener al menos un número."
	}
	if !hasSpecial {
		return false, "La contraseña debe contener al menos un carácter especial (@ $ ! % * ? &)."
	}

	return true, ""
}
