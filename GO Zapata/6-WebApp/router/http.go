package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
	"webapp/models"
	"webapp/utilidades"
)

// var Token string = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6MzYsImlhdCI6MTY4MzE1NTM1NiwiZXhwIjoxNjg1NzQ3MzU2fQ.eEZZHIqiM5FpR8ZwK3jPd-qT367epSK5qjoHU9f7r1I"
// METODO GET
func ClientesGET(response http.ResponseWriter, request *http.Request) {
	// Parseamos la plantilla
	tmpl := template.Must(template.ParseFiles("templates/http/http_cliente.html", utilidades.Front))

	// Creamos el cliente HTTP
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8070/v1/clientes", nil)
	if err != nil {
		http.Error(response, "Error creando la solicitud", http.StatusInternalServerError)
		fmt.Println("Error creando la solicitud:", err)
		return
	}

	// Realizamos la solicitud
	reg, err := client.Do(req)
	if err != nil {
		http.Error(response, "Error en la petición al servicio", http.StatusInternalServerError)
		fmt.Println("Error en la petición:", err)
		return
	}
	defer reg.Body.Close()

	// Verificamos que el código de estado sea 200 OK
	if reg.StatusCode != http.StatusOK {
		http.Error(response, "Error en la respuesta del servicio", reg.StatusCode)
		fmt.Println("Error en la respuesta del servicio:", reg.Status)
		return
	}

	// Decodificamos la respuesta JSON en una lista de clientes
	var clientes models.Clientes
	err = json.NewDecoder(reg.Body).Decode(&clientes)
	if err != nil {
		http.Error(response, "Error procesando la respuesta JSON", http.StatusInternalServerError)
		fmt.Println("Error procesando la respuesta JSON:", err)
		return
	}

	// Creamos la estructura ClienteHttp con los datos obtenidos
	data := models.ClienteHttp{
		Css:     "success",
		Mensaje: "Datos cargados correctamente",
		Datos:   clientes,
	}

	// Renderizamos la plantilla con los datos cargados
	err = tmpl.Execute(response, data)
	if err != nil {
		http.Error(response, "Error ejecutando la plantilla", http.StatusInternalServerError)
		fmt.Println("Error ejecutando la plantilla:", err)
	}
}

// // METODO POST
// ClienteForm maneja la creación de un cliente a través de un formulario HTML
func ClienteForm(response http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("templates/http/ClientePost.html", utilidades.Front)
	if err != nil {
		http.Error(response, "Error cargando la plantilla", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(response, nil)
}

// ClientePOST maneja la creación de clientes con el REST
func ClientePOST(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Procesar los datos del formulario
	err := request.ParseForm()
	if err != nil {
		http.Error(response, "Error procesando el formulario", http.StatusBadRequest)
		return
	}

	// Crear un cliente sin ID
	clienteSinID := struct {
		Nombre   string `json:"nombre"`
		Correo   string `json:"correo"`
		Telefono string `json:"telefono"`
	}{
		Nombre:   request.FormValue("nombre"),
		Correo:   request.FormValue("correo"),
		Telefono: request.FormValue("telefono"),
	}

	// Convertir a JSON sin ID
	jsonData, err := json.MarshalIndent(clienteSinID, "", "  ")
	if err != nil {
		http.Error(response, "Error serializando JSON", http.StatusInternalServerError)
		return
	}

	// Imprimir en consola
	fmt.Println("Datos enviados en JSON:")
	fmt.Println(string(jsonData))

	// Enviar JSON a la API REST
	ClientePOSTREST(jsonData, response)
}

// ClientePOSTREST envía el JSON a la API REST y gestiona la respuesta
func ClientePOSTREST(jsonData []byte, response http.ResponseWriter) {
	url := "http://localhost:8070/v1/clientes"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(response, "Error creando la petición", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(response, "Error enviando la petición", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Leer respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(response, "Error leyendo la respuesta del backend", http.StatusInternalServerError)
		return
	}

	// Verificar código de estado
	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Cliente dado de alta correctamente")
		fmt.Println("JSON enviado:", string(jsonData))
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("Cliente creado correctamente"))
	} else {
		fmt.Println("Error al dar de alta al cliente:", string(body))
		http.Error(response, "Error al dar de alta al cliente", resp.StatusCode)
	}
}

// METODO UPDATE
// ClienteFormUpdate maneja la vista del formulario para actualizar un cliente.
func ClienteFormUpdate(response http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("templates/http/ClientePUT.html", utilidades.Front)
	if err != nil {
		http.Error(response, "Error cargando la plantilla", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(response, nil)
}

// ClienteUpdate maneja la actualización de clientes mediante una solicitud PUT a la API REST.
func ClienteUpdate(response http.ResponseWriter, request *http.Request) {
	// Convertimos POST a PUT si es necesario
	if request.Method == http.MethodPost && request.FormValue("_method") == "PUT" {
		request.Method = http.MethodPut
	}

	// Extraer ID del formulario
	id := request.FormValue("id")
	if id == "" {
		http.Error(response, "ID del cliente es requerido", http.StatusBadRequest)
		return
	}

	// Extraer datos del formulario
	cliente, err := extraerDatosCliente(request)
	if err != nil {
		http.Error(response, "Error al procesar los datos del formulario", http.StatusBadRequest)
		return
	}

	// Convertir a JSON
	jsonData, err := json.Marshal(cliente)
	if err != nil {
		http.Error(response, "Error convirtiendo a JSON", http.StatusInternalServerError)
		return
	}

	// Llamar a la API REST con el ID en la URL
	statusCode, err := ClienteUpdateREST(id, jsonData)
	if err != nil {
		http.Error(response, err.Error(), statusCode)
		return
	}

	// Responder con éxito
	response.WriteHeader(http.StatusOK)
	fmt.Fprintln(response, "Cliente actualizado correctamente")
}

// ClienteUpdateREST envía una solicitud PUT a la API REST para actualizar un cliente.
func ClienteUpdateREST(id string, jsonData []byte) (int, error) {
	// Aseguramos que la URL tenga el ID correcto
	fmt.Println("JSON enviado:", string(jsonData))
	apiURL := fmt.Sprintf("http://localhost:8070/v1/clientes/%s", id) // ✅ Ahora la URL incluye el ID

	req, err := http.NewRequest(http.MethodPut, apiURL, bytes.NewBuffer(jsonData)) // ✅ Método PUT
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Error creando la solicitud: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Error enviando la solicitud: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf("Error al actualizar el cliente, código: %d", resp.StatusCode)
	}

	return http.StatusOK, nil
}

// extraerDatosCliente extrae y valida los datos del formulario para crear un objeto Cliente.
func extraerDatosCliente(request *http.Request) (models.Cliente, error) {
	if err := request.ParseForm(); err != nil {
		return models.Cliente{}, fmt.Errorf("error al procesar el formulario: %w", err)
	}

	idStr := strings.TrimSpace(request.FormValue("id"))
	nombre := strings.TrimSpace(request.FormValue("nombre"))
	correo := strings.TrimSpace(request.FormValue("correo"))
	telefono := strings.TrimSpace(request.FormValue("telefono"))

	if idStr == "" || nombre == "" || correo == "" || telefono == "" {
		return models.Cliente{}, errors.New("todos los campos son obligatorios")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return models.Cliente{}, fmt.Errorf("ID inválido: %w", err)
	}

	return models.Cliente{
		Id:       id,
		Nombre:   nombre,
		Correo:   correo,
		Telefono: telefono,
	}, nil
}

//METODO DELETE
