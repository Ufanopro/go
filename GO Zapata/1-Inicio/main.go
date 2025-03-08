package main

import (
	"fmt"
	"reflect"
)

/*
// HOLA MUNDO
func main() {
	fmt.Println("Hola GO Zapata!")
}

// VARIABLES Y CONSTANTES
const MiConstante = "Valor de mi constante propia"

func main() {
	//declaración por inferencia
	var nombre string = "César"
	fmt.Println(nombre)
	//declaración rápida o corta
	nombre2 := "Juan"
	fmt.Println(nombre2)
	fmt.Printf("El valor de MiConstante es: %s \n", MiConstante)
}
// TIPOS DATOS
func main() {
	var string1 string = "texto"
	fmt.Println(string1)
	textoGrande := `Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Do`
	fmt.Println("-----------------------------------------------")
	fmt.Println(textoGrande)
	var estado bool = false
	fmt.Println("-----------------------------------------------")
	fmt.Println(estado)
	var flotante32 float32 = 32.33
	fmt.Println("-----------------------------------------------")
	fmt.Println(flotante32)
	var flotante64 float64 = 64.33
	fmt.Println("-----------------------------------------------")
	fmt.Println(flotante64)
	var entero int = 1234
	fmt.Println("-----------------------------------------------")
	fmt.Println(entero)
	var entero_int8 int8 = 123 //-128 a 127
	fmt.Println("-----------------------------------------------")
	fmt.Println(entero_int8)
	var entero_int16 int16 = 123 //-2^15 a 2^15 -1
	fmt.Println("-----------------------------------------------")
	fmt.Println(entero_int16)
	var entero_int32 int32 = 45611 //-2^31 a 2^31 -1
	fmt.Println("-----------------------------------------------")
	fmt.Println(entero_int32)
	var entero_int64 int64 = 78910 //-2^63 a 2^63 -1
	fmt.Println("-----------------------------------------------")
	fmt.Println(entero_int64)
	var entero_uint8 uint8 = 233 //0 a 255
	fmt.Println("-----------------------------------------------")
	fmt.Println(entero_uint8)
	var entero_uint16 uint16 = 12451 //0 a 2^16 -1
	fmt.Println("-----------------------------------------------")
	fmt.Println(entero_uint16)
	var entero_uint32 uint32 = 789123235 //0 a 2^32 -1
	fmt.Println("-----------------------------------------------")
	fmt.Println(entero_uint32)
	var entero_uint64 uint64 = 7891234566321 //0 a 2^64 -1
	fmt.Println("-----------------------------------------------")
	fmt.Println(entero_uint64)
}

// REFLECT y TYPEOF
// import "reflect"
func main() {
	string1 := true
	fmt.Println(reflect.TypeOf(string1))
}

// PUNTEROS
func main() {
	var estado bool = true
	color := "rojo"
	fmt.Println(color, &color)
	fmt.Println(estado, &estado)
}

// CONDICIONALES
func main() {
	//operadores de comparación

	//x == y | Es igual x igual a y ?
	//x != y | Es x diferente de y ?
	//x < y | Es x menor que y?
	//x <= y | Es x menor o igual que y?
	//x > y | Es x mayor que y?
	//x >= y | Es x mayor o igual que y?
	edad := 11
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("eres menor de edad")
	}
	fmt.Println("-----------------------------------------")
	color := "azul"
	if color == "rojo" {
		fmt.Println("Es rojo como la sangre de los araucanos")
	} else if color == "blanco" {
		fmt.Println("Es blanco como la nieve")
	} else if color == "azul" {
		fmt.Println("Es azul como el cielo")
	} else {
		fmt.Println("No tiene color patrio")
	}

	//operadores lógicos && (and) || (or) ! (not)
	if color == "azul" && edad == 11 {
		fmt.Println("-----------------------------------------")
		fmt.Println("color es azul y edad 11")
	}
	//declarar variable en una condición
	if variable := 2; variable == 1 {
		fmt.Println("-----------------------------------------")
		fmt.Println("variable es igual a 2")
	}
	//switch case
	switch color {
	case "rojo":
		fmt.Println("-----------------------------------------")
		fmt.Println("Es rojo como la sangre de los araucanos")
		break
	case "azul":
		fmt.Println("-----------------------------------------")
		fmt.Println("Es azul como el cielo sss")
		break
	case "blanco":
		fmt.Println("-----------------------------------------")
		fmt.Println("Es blanco como la nieve")
		break
	default:
		fmt.Println("-----------------------------------------")
		fmt.Println("No tiene color patrio")
		break
	}
}

// ITERACIONES
func main() {
	fmt.Println("------------WHILE ----------------")
	contador := 0
	for contador < 5 {
		fmt.Println(contador)
		contador++
	}
	fmt.Println("------------DO WHILE ----------------")
	contador = 0
	for ok := true; ok; ok = (contador < 5) {
		fmt.Println(contador)
		contador++
	}
	fmt.Println("------------DO WHILE BREAK ----------------")
	contador = 0
	for {
		fmt.Println(contador)
		contador++
		if contador >= 5 {
			break
		}
	}
	// Bucle que imprime los números del 0 al 4
	fmt.Println("------------FOR ----------------")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("------------FOR EACH----------------")
	// Definición de un slice de cadenas
	frutas := []string{"manzana", "banana", "cereza"}

	// Bucle que itera sobre cada elemento del slice
	for indice, fruta := range frutas {
		fmt.Printf("Índice: %d, Fruta: %s\n", indice, fruta)
	}
}

// ARRAYS y SLICES
func main() {
	//arreglo (array)
	var paises [5]string
	paises[0] = "Chile"
	paises[1] = "Perú"
	paises[2] = "Bolivia"
	paises[3] = "Venezuela"
	paises[4] = "España"
	fmt.Println(paises)
	fmt.Println(paises[2])
	fmt.Println("El largo del arreglo es ", len(paises))
	fmt.Printf("El largo es %v \n", len(paises))
	fmt.Println(reflect.TypeOf(paises))
	//slice
	fmt.Println("------------------------------------------------")
	var paises2 = make([]string, 5)
	paises2[0] = "México"
	paises2[1] = "Ecuador"
	paises2[2] = "Argentina"
	paises2[3] = "Uruguay"
	paises2[4] = "España"
	fmt.Println(paises2)
	fmt.Println(paises2[2])
	fmt.Println("El largo del arreglo es ", len(paises2))
	fmt.Printf("El largo es %v \n", len(paises2))
	fmt.Println(reflect.TypeOf(paises2))
	fmt.Println("------------------------------------------------")
	// agregar un elemento al slice
	paises2 = append(paises2, "Noruega")
	fmt.Println(paises2)
	fmt.Println(paises2[2])
	fmt.Println("El largo del arreglo es ", len(paises2))
	fmt.Printf("El largo es %v \n", len(paises2))
	fmt.Println(reflect.TypeOf(paises2))
	//eliminar un elemento
	paises2 = append(paises[:2], paises2[2+1:]...)
	fmt.Println("------------------------------------------------")
	fmt.Println(paises2)
	fmt.Println(paises2[1])
	fmt.Println("El largo del arreglo es ", len(paises2))
	fmt.Printf("El largo es %v \n", len(paises2))
	fmt.Println(reflect.TypeOf(paises2))
}
*/
// MAP
func main() {
	//{"id":1, "nombre":"chile"}
	paises := make(map[string]int)
	paises["argentina"] = 4000000
	paises["españa"] = 46000000
	paises["brasil"] = 190000000
	paises["uruguay"] = 3400000
	paises["chile"] = 20000000
	fmt.Println(paises)
	fmt.Println(reflect.TypeOf(paises))
	fmt.Println(paises["chile"])
	fmt.Println("------------------------------------------------")
	paises2 := map[int]string{
		1:  "Chile",
		2:  "Perú",
		3:  "Brasil",
		4:  "México",
		5:  "Venezuela",
		6:  "Argentina",
		7:  "España",
		8:  "Dinamarca",
		11: "USA",
	}
	fmt.Println(paises2)
	fmt.Println(paises2[1])
	//veamos si existe algún valor en el map
	fmt.Println("------------------------------------------------")
	pais, existe := paises2[11]
	if existe {
		fmt.Println("Si existe el pais:", pais)
	} else {
		fmt.Println("NO existe el pais:", pais)
	}
	//eliminar un elemento de un país
	fmt.Println("------------------------------------------------")
	delete(paises2, 1)
	fmt.Println(paises2)
	//recorrer un map con for
	fmt.Println("------------------------------------------------")
	for id, valor := range paises2 {
		fmt.Printf("ID: %v | Nombre: %v \n", id, valor)
	}
	fmt.Println("------------------------------------------------")
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Cualquier mensaje  ",
	}
	fmt.Println(respuesta)
	fmt.Println("estado=", respuesta["estado"])
}
