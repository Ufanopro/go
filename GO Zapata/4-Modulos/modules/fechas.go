package fechas

import (
	"errors"
	"strings"
)

// Función que convierte el día de la semana en su número correspondiente
func GetDayNumber(day string) (int, error) {
	// Convertir el día a minúsculas para evitar problemas con mayúsculas/minúsculas
	day = strings.ToLower(day)

	// Mapear los días de la semana a sus números correspondientes
	days := map[string]int{
		"lunes":     1,
		"martes":    2,
		"miércoles": 3,
		"jueves":    4,
		"viernes":   5,
		"sábado":    6,
		"domingo":   7,
	}

	// Verificar si el día existe en el mapa
	if num, exists := days[day]; exists {
		return num, nil
	}
	return 0, errors.New("día de la semana inválido")
}
