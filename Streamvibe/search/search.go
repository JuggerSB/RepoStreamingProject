package search

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Movie struct {
	Title       string `json:"title"`
	Overview    string `json:"overview"`
	ReleaseDate string `json:"release_date"`
}

// MovieSearchResponse representa la respuesta de la búsqueda de películas de TMDB.
type MovieSearchResponse struct {
	Results []Movie `json:"results"`
}

func Searchtmdb() {
	// Solicitar al usuario el término de búsqueda
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese el término de búsqueda: ")
	scanner.Scan()
	searchTerm := scanner.Text()

	apiKey := "52abf4732494d35e674f7b2345c0486f"
	escapedSearchTerm := url.QueryEscape(searchTerm)
	searchURL := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=%s&query=%s&include_adult=false&language=es-ES&page=1", apiKey, escapedSearchTerm)

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response MovieSearchResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("Error al deserializar la respuesta: ", err)
	}

	for _, movie := range response.Results {
		fmt.Printf("Título: %s\n", movie.Title)
		fmt.Printf("Fecha de lanzamiento: %s\n", movie.ReleaseDate)
		fmt.Printf("Descripción: %s\n\n", movie.Overview)

	}
}

// BUSQUEDA EN YOUTUBE
// Estructura para almacenar la información relevante de cada video e imprimir solo eso
type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	// Puedes agregar más campos según sea necesario
}

func SearchYoutube() {
	ctx := context.Background()
	apiKey := "AIzaSyC8LiUPsvPqJQow-VfgzueO9DchLHwr8Hk" // Reemplaza con tu clave API real

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error al crear el servicio de YouTube: %v", err)
	}

	// Solicitar al usuario el término de búsqueda
	fmt.Print("Ingrese el término de búsqueda: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchTerm := scanner.Text()

	call := service.Search.List([]string{"snippet"}).Q(searchTerm).MaxResults(5)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error al realizar la solicitud de búsqueda: %v", err)
	}

	// Iterar sobre los resultados y mostrar solo titulo y descripcion
	for _, item := range response.Items {
		fmt.Printf("Título: %s\n", item.Snippet.Title)
		fmt.Printf("Descripción: %s\n", item.Snippet.Description)

		// Crear y mostrar el enlace al video de YouTube
		if item.Id.Kind == "youtube#video" {
			videoUrl := fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id.VideoId)
			fmt.Printf("Ver en YouTube: %s\n", videoUrl)
		}

		fmt.Println()
	}
}
