package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Configura las rutas para cada API
	r.GET("/api1", handleAPI1)
	r.GET("/api2", handleAPI2)
	r.GET("/api3", handleAPI3)
	r.GET("/api4", handleAPI4)

	// Inicia el servidor en el puerto 8080
	r.Run(":8080")
}

func handleAPI1(c *gin.Context) {
	// Lógica para conectarse y obtener contenido de la API 1
	// Reproduce el contenido multimedia
}

func handleAPI2(c *gin.Context) {
	// Lógica para conectarse y obtener contenido de la API 2
	// Reproduce el contenido multimedia
}

func handleAPI3(c *gin.Context) {
	// Lógica para conectarse y obtener contenido de la API 3
	// Reproduce el contenido multimedia
}

func handleAPI4(c *gin.Context) {
	// Lógica para conectarse y obtener contenido de la API 4
	// Reproduce el contenido multimedia
}
