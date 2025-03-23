package utilidades

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gorilla/sessions"
	"github.com/jordan-wright/email"
)

// Ruta a la interfaz de usuario
var Front = "templates/front/front.html"

// Almacén de sesiones con clave segura desde variables de entorno
var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

// RetornarLogin obtiene el ID y nombre del usuario desde la sesión
func RetornarLogin(request *http.Request) (string, string) {
	session, _ := Store.Get(request, "session-name")

	tamilaID, tamilaNombre := "", ""
	if id, ok := session.Values["tamila_id"].(string); ok {
		tamilaID = id
	}
	if nombre, ok := session.Values["tamila_nombre"].(string); ok {
		tamilaNombre = nombre
	}
	return tamilaID, tamilaNombre
}

// RetornarMensajesFlash obtiene mensajes flash almacenados en la sesión
func RetornarMensajesFlash(response http.ResponseWriter, request *http.Request) (string, string) {
	session, _ := Store.Get(request, "flash-session")

	cssSesion, mensajeSesion := "", ""
	if flashes := session.Flashes("css"); len(flashes) > 0 {
		if css, ok := flashes[0].(string); ok {
			cssSesion = css
		}
	}
	if flashes := session.Flashes("mensaje"); len(flashes) > 0 {
		if mensaje, ok := flashes[0].(string); ok {
			mensajeSesion = mensaje
		}
	}

	_ = session.Save(request, response)
	return cssSesion, mensajeSesion
}

// CrearMensajesFlash almacena mensajes flash en la sesión
func CrearMensajesFlash(response http.ResponseWriter, request *http.Request, css string, mensaje string) {
	session, err := Store.Get(request, "flash-session")
	if err != nil {
		http.Error(response, "Error obteniendo la sesión", http.StatusInternalServerError)
		return
	}
	session.AddFlash(css, "css")
	session.AddFlash(mensaje, "mensaje")
	_ = session.Save(request, response)
}

// EnviarEmail envía un correo usando SMTP con configuración segura
func EnviarEmail(destinatario, asunto, cuerpo string) error {
	// Obtener credenciales SMTP desde variables de entorno
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	if smtpServer == "" || smtpPort == "" || username == "" || password == "" {
		return fmt.Errorf("❌ Configuración SMTP incompleta. Verifica las variables de entorno.")
	}

	// Crear el email
	e := email.NewEmail()
	e.From = username
	e.To = []string{destinatario}
	e.Subject = asunto
	e.Text = []byte(cuerpo)

	// Configurar autenticación SMTP
	auth := smtp.PlainAuth("", username, password, smtpServer)
	tlsConfig := &tls.Config{ServerName: smtpServer}

	// Enviar email con TLS
	err := e.SendWithTLS(smtpServer+":"+smtpPort, auth, tlsConfig)
	if err != nil {
		return fmt.Errorf("❌ Error enviando email: %w", err)
	}

	fmt.Println("✅ Correo enviado con éxito a", destinatario)
	return nil
}
