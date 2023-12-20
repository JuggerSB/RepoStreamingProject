package tmdbapi

import (
	"log"

	"github.com/ryanbradynd05/go-tmdb"
)

var tmdbAPI *tmdb.TMDb

// InitTMDB inicializa el cliente de TMDB con la clave API dada.
func InitTMDB(apiKey string) {
	config := tmdb.Config{
		APIKey: apiKey,
	}
	tmdbAPI = tmdb.Init(config)
}

// GetMovieInfo obtiene los detalles de una película por su ID.
func GetMovieInfo(movieID int) *tmdb.Movie {
	movie, err := tmdbAPI.GetMovieInfo(movieID, nil)
	if err != nil {
		log.Println("Error al obtener detalles de película:", err)
		return nil
	}
	return movie
}
