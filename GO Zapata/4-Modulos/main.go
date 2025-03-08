package main

import (
	fechas "4-Modulos/modules"
	"fmt"
)

/*
// TIME
func main() {
	// 1. Obtener la hora actual
	now := time.Now()
	fmt.Println("Hora actual:", now)

	// 2. Formatear la fecha y hora
	formatted := now.Format("2006-01-02 15:04:05")
	fmt.Println("Fecha y hora formateada:", formatted)

	// 3. Agregar tiempo (sumar 1 hora)
	later := now.Add(time.Hour)
	fmt.Println("Dentro de 1 hora:", later)

	// 4. Calcular la diferencia entre dos tiempos
	past := now.Add(-48 * time.Hour) // Hace 2 días
	diff := now.Sub(past)
	fmt.Println("Diferencia de tiempo:", diff)

	// 5. Pausar la ejecución (dormir 2 segundos)
	fmt.Println("Esperando 2 segundos...")
	time.Sleep(2 * time.Second)
	fmt.Println("Continuando ejecución")
}
// STRINGS
func main() {
	texto := "Hola, Mundo"

	// 1. Convertir a minúsculas
	minusculas := strings.ToLower(texto)
	fmt.Println("Minúsculas:", minusculas)

	// 2. Convertir a mayúsculas
	mayusculas := strings.ToUpper(texto)
	fmt.Println("Mayúsculas:", mayusculas)

	// 3. Verificar si contiene una subcadena
	contiene := strings.Contains(texto, "Mundo")
	fmt.Println("Contiene 'Mundo'?", contiene)

	// 4. Reemplazar una palabra
	reemplazado := strings.ReplaceAll(texto, "Mundo", "Go")
	fmt.Println("Reemplazado:", reemplazado)

	// 5. Dividir una cadena en partes
	partes := strings.Split(texto, ", ")
	fmt.Println("Partes:", partes)
}

// MATH y RAND
func main() {
	// Crear un generador de números aleatorios independiente
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// 1. Generar un número aleatorio entre 0 y 99
	numAleatorio := r.Intn(100)
	fmt.Println("Número aleatorio entre 0 y 99:", numAleatorio)

	// 2. Calcular la raíz cuadrada
	numero := 25.0
	raizCuadrada := math.Sqrt(numero)
	fmt.Println("Raíz cuadrada de", numero, ":", raizCuadrada)

	// 3. Obtener el valor absoluto
	valor := -42.5
	absValor := math.Abs(valor)
	fmt.Println("Valor absoluto de", valor, ":", absValor)

	// 4. Redondear un número
	numRedondeo := 3.7
	redondeado := math.Round(numRedondeo)
	fmt.Println("Redondeo de", numRedondeo, ":", redondeado)

	// 5. Calcular el máximo entre dos números
	a, b := 15.4, 22.1
	maximo := math.Max(a, b)
	fmt.Println("Máximo entre", a, "y", b, ":", maximo)
}

// GENERADOR CONTRASEÑAS
// Definir conjuntos de caracteres
const (
	upper     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower     = "abcdefghijklmnopqrstuvwxyz"
	numberSet = "0123456789"
	special   = "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
)

// Generar una contraseña aleatoria con longitud especificada
func generatePassword(length int) string {
	if length < 4 {
		return "La contraseña debe tener al menos 4 caracteres"
	}

	// Crear un generador de números aleatorios independiente
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// Asegurar al menos un carácter de cada tipo
	password := []byte{
		upper[r.Intn(len(upper))],
		lower[r.Intn(len(lower))],
		numberSet[r.Intn(len(numberSet))],
		special[r.Intn(len(special))],
	}

	// Combinar todos los conjuntos
	allChars := upper + lower + numberSet + special

	// Completar la contraseña con caracteres aleatorios
	for i := 4; i < length; i++ {
		password = append(password, allChars[r.Intn(len(allChars))])
	}

	// Mezclar la contraseña para evitar patrones predecibles
	r.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}

func main() {
	length := 12 // Longitud deseada de la contraseña
	password := generatePassword(length)
	fmt.Println("Contraseña generada:", password)
}

// OS
os.Getenv(): Recupera el valor de una variable de entorno.
os.Mkdir(): Crea un directorio.
os.Remove(): Elimina un archivo o directorio.
os.Open(): Abre un archivo.
os.Create(): Crea un archivo, o lo abre si ya existe.
os.Stat(): Obtiene información sobre un archivo (tamaño, nombre, etc.).
os.Getwd(): Obtiene el directorio de trabajo actual.
os.Exit(): Termina el programa con un código de estado (usualmente 0 para éxito, cualquier número distinto de 0 para error).

// LOG
// SIMPLE
func main() {
	// Registrar un mensaje simple de log
	log.Println("Este es un mensaje de log simple.")
}
// ERROR SIMPLE
func main() {
	// Crear un archivo de log
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("No se pudo abrir el archivo de log:", err)
	}
	defer file.Close()

	// Establecer el archivo como el destino del log
	log.SetOutput(file)

	// Registrar un mensaje de error
	log.Println("Este es un mensaje de error.")

	// Registrar un error fatal (el programa se detendrá después de esto)
	log.Fatal("Error fatal ocurrido.")
}
// LOG PERSONALIZADO
func main() {
	// Registrar un mensaje de log con un prefijo y un formato
	log.SetPrefix("INFO: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("Este es un mensaje con formato.")
}

// LOG CON FECHA
func main() {
	// Establecer flags para que el log incluya la fecha y la hora
	log.SetFlags(log.Ldate | log.Ltime)

	log.Println("Este mensaje de log incluye la fecha y la hora.")
}

// LOG personalizado con múltiples destinos
func main() {
	// Crear un archivo de log
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("No se pudo abrir el archivo de log:", err)
	}
	defer file.Close()

	// Crear un logger personalizado con múltiples destinos
	multiLog := log.New(file, "CUSTOM_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Registrar un mensaje en el archivo
	multiLog.Println("Este mensaje se graba en un archivo.")

	// También registrar un mensaje en la consola
	log.Println("Este mensaje se muestra en la consola.")
}
// LOG personalizado con múltiples destinos
func logInfo(message string) {
	log.SetPrefix("INFO: 🌟 ")
	log.Println(message)
}

func logWarn(message string) {
	log.SetPrefix("WARN: ⚠️ ")
	log.Println(message)
}

func logError(message string) {
	log.SetPrefix("ERROR: ❌ ")
	log.Println(message)
}

func main() {
	// Registrar mensajes con diferentes niveles y emojis
	logInfo("Este es un mensaje de información.")
	logWarn("Este es un mensaje de advertencia.")
	logError("Este es un mensaje de error.")
}
*/
// MODULO
func main() {
	// Ejemplo de uso
	daysList := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}

	for _, day := range daysList {
		dayNumber, err := fechas.GetDayNumber(day)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("El día '%s' corresponde al número %d de la semana.\n", day, dayNumber)
		}
	}
}
