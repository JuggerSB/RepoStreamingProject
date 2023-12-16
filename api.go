package main

import (
	"fmt"
	"log"

	"github.com/anaskhan96/go-tmdb"
)

func main() {
	// Configura tu clave de API de TMDb
	apiKey := "tu_clave_de_api"
	tmdbClient, err := tmdb.Init(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	// Realiza una consulta para obtener información sobre una película específica (por ejemplo, con ID 123)
	movieID := 123
	movie, err := tmdbClient.GetMovieDetails(movieID, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Imprime información sobre la película
	fmt.Printf("Título: %s\n", movie.Title)
	fmt.Printf("Sinopsis: %s\n", movie.Overview)

	// Puedes agregar lógica adicional para reproducir la película aquí
}
