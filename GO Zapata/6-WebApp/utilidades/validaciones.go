package utilidades

import (
	"errors"
	"regexp"
	"unicode"
)

// Expresiones regulares compiladas solo una vez
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// ValidateEmail verifica si el correo es válido y devuelve un error si no lo es
func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("❌ El correo electrónico no puede estar vacío.")
	}
	if !emailRegex.MatchString(email) {
		return errors.New("❌ El correo electrónico no es válido. Debe tener un formato correcto.")
	}
	return nil
}

// ValidatePassword verifica si la contraseña cumple con los requisitos y devuelve un error si no lo hace
func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("❌ La contraseña no puede estar vacía.")
	}
	if len(password) < 8 {
		return errors.New("❌ La contraseña debe tener al menos 8 caracteres.")
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	specialChars := "@$!%*?&"

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case containsRune(specialChars, char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("❌ La contraseña debe contener al menos una letra mayúscula.")
	}
	if !hasLower {
		return errors.New("❌ La contraseña debe contener al menos una letra minúscula.")
	}
	if !hasNumber {
		return errors.New("❌ La contraseña debe contener al menos un número.")
	}
	if !hasSpecial {
		return errors.New("❌ La contraseña debe contener al menos un carácter especial (@ $ ! % * ? &).")
	}

	return nil
}

// containsRune verifica si un carácter está en una cadena
func containsRune(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}
