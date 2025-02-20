package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album representa datos sobre un álbum musical.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// slice de albums para almacenar datos iniciales de álbumes.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responde con la lista de todos los álbumes como JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums agrega un álbum desde JSON recibido en el cuerpo de la solicitud.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Llama a BindJSON para vincular el JSON recibido a
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Agrega el nuevo álbum a la lista.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID encuentra el álbum cuyo valor de ID coincide con el parámetro id
// enviado por el cliente, y luego devuelve ese álbum como respuesta.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Recorre la lista de álbumes, buscando
	// un álbum cuyo valor de ID coincida con el parámetro.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
}

// putAlbum actualiza un álbum existente según su ID.
func putAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album

	// Vincula el JSON recibido al objeto updatedAlbum.
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	// Busca y actualiza el álbum por ID.
	for i, a := range albums {
		if a.ID == id {
			updatedAlbum.ID = id // Asegura que el ID se mantenga igual
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
}

// deleteAlbum elimina un álbum según su ID.
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	// Busca y elimina el álbum por ID.
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "álbum eliminado"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
}

func main() {
	/*Crea una nueva instancia de Gin
	router := gin.Default()

	// Define la ruta GET para el endpoint "/"
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "¡Hola Mundo!",
		})
	})

	// Inicia el servidor en el puerto 8080
	router.Run("localhost:8080")
	*/
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", putAlbum)       // Actualizar álbum
	router.DELETE("/albums/:id", deleteAlbum) // Eliminar álbum

	router.Run("localhost:8080")
}
