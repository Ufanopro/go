package router

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	modelos "webapp/models"
	"webapp/utilidades"

	"github.com/jung-kurt/gofpdf"

	qrcode "github.com/skip2/go-qrcode"
	excelize "github.com/xuri/excelize/v2"
)

func Utiles(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/recursos/recursos.html", utilidades.Front))
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

func PDF(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/pdf/pdf.html", utilidades.Front))
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

func PDFCreate(response http.ResponseWriter, request *http.Request) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {

	}
	fmt.Println("Se creo el documento")
	http.Redirect(response, request, "/recursos/pdf", http.StatusSeeOther)
}

func XLS(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/xls/xls.html", utilidades.Front))
	template.Execute(response, nil)
	/*
		fmt.Fprintln(response, "Hola WebAPP")

		template, err := template.ParseFiles("templates/index/index.html", "templates/front/front.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(response, nil)
		}

	*/ //excel
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.SetCellValue("Sheet1", "A1", "id")
	f.SetCellValue("Sheet1", "B1", "nombre")
	f.SetCellValue("Sheet1", "C1", "correo")
	f.SetActiveSheet(index)
	//le agregamos más datos al excel
	cliente1 := modelos.Clientes{modelos.Cliente{1, "César Cancino", "info@tamila.cl", ""}, modelos.Cliente{2, "Juan Pérez", "juanito@tamila.cl", ""}}
	contador := 2 //ésto es un int
	i := 0
	for _, service := range cliente1 {

		fila := strconv.Itoa(contador)

		f.SetCellValue("Sheet1", "A"+fila, service.Id)
		f.SetCellValue("Sheet1", "B"+fila, service.Nombre)
		f.SetCellValue("Sheet1", "C"+fila, service.Correo)
		contador++
		i++
	}
	//ahora construimos el documento excel
	time := strings.Split(time.Now().String(), " ")
	nombre := string(time[4][6:14]) + ".xlsx"
	if err := f.SaveAs("./" + nombre); err != nil {
		fmt.Println(err)
	}
	//retornamos

	template.Execute(response, nil)

}

const qrDirectory = "./static/downloads/"

func QR(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/qr/qr.html"))

	// Crear directorio si no existe
	if err := os.MkdirAll(qrDirectory, os.ModePerm); err != nil {
		log.Println("Error al crear el directorio:", err)
		http.Error(response, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	// Nombre del archivo basado en timestamp para evitar colisiones
	filename := fmt.Sprintf("qr_%d.png", time.Now().Unix())
	filePath := filepath.Join(qrDirectory, filename)

	// Generar QR y guardarlo en el directorio
	if err := qrcode.WriteFile("https://www.cesarcancino.com/cursos-en-udemy/", qrcode.High, 256, filePath); err != nil {
		log.Println("Error al generar el código QR:", err)
		http.Error(response, "Error al generar el código QR", http.StatusInternalServerError)
		return
	}

	// Pasar la ruta relativa de la imagen al template
	template.Execute(response, map[string]string{"QRPath": "/static/downloads/" + filename})
}

func Email(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/email/email.html", utilidades.Front))

	err := utilidades.EnviarEmail("ufano.pro@gmail.com", "Prueba Acumbamail", "Este es un email de prueba enviado desde Go.")
	if err != nil {
		fmt.Println("Error:", err)
	}
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
