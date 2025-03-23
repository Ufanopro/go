package util

import (
	"os"
	modelos "restapi/models"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerarJWT(usuario modelos.Usuario) (string, error) {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic("Error loading .env file")
	}
	miClave := []byte(os.Getenv("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"correo":         usuario.Correo,
		"nombre":         usuario.Nombre,
		"generado_desde": "https://www.cds-holdings.com",
		"id":             usuario.Id,
		"iat":            time.Now().Unix(),
		"exp":            time.Now().Add(time.Hour * 24).Unix(), //24 HORAS
	})
	tokenString, err := token.SignedString(miClave)
	return tokenString, err
}
