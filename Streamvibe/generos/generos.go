package generos

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GenreList struct {
	Genres []Genre `json:"genres"`
}

func ShowGenres() {
	apiKey := "52abf4732494d35e674f7b2345c0486f"
	url := fmt.Sprintf("https://api.themoviedb.org/3/genre/movie/list?api_key=%s&language=es", apiKey)

	req, err := http.NewRequest("GET", url, nil)
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

	var genreList GenreList
	err = json.Unmarshal(body, &genreList)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Lista de Géneros:")
	for _, genre := range genreList.Genres {
		fmt.Printf("- ID: %d, Género: %s\n", genre.ID, genre.Name)
	}
}
