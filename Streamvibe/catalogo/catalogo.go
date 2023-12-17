package catalogo

import (
	"Streamvibe/db"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Pelicula struct {
	PeliculaID      int
	Titulo          string
	Director        string
	Duracion        int
	AnioLanzamiento int
	Descripcion     string
	Disponible      bool
	Genero          string
}

// Crear Lista de Reproduccion:
// Definimos la estructura de un género
type Genero struct {
	GeneroID int
	Genero   string
}

// Catalogo mantiene la información sobre las películas seleccionadas
type Catalogo struct {
	peliculaSeleccionada   *Pelicula
	peliculasSeleccionadas []Pelicula // Slice con la lista de películas seleccionadas
}

// SetPeliculaSeleccionada establece la película actualmente seleccionada en el catálogo
func (c *Catalogo) SetPeliculaSeleccionada(pelicula *Pelicula) {
	c.peliculaSeleccionada = pelicula
}

// GetPeliculaSeleccionada nos permite obtener la o las películas actualmente seleccionadas del catálogo
func (c *Catalogo) GetPeliculaSeleccionada() *Pelicula {
	return c.peliculaSeleccionada
}

// Declaramos la variable catalogo
var catalogo Catalogo

func ShowCatalog() {
	rows, err := db.DB.Query("SELECT p.PeliculaID, p.Titulo, p.Director, p.Duracion, p.AnioLanzamiento, p.Descripcion, p.Disponible, g.Genero FROM Peliculas p INNER JOIN Genero g ON p.GeneroID = g.GeneroID")
	if err != nil {
		log.Fatal("Error al obtener el catálogo de películas:", err)
	}
	defer rows.Close()

	var peliculas []Pelicula

	fmt.Println("Catálogo de Películas:")
	for rows.Next() {
		var pelicula Pelicula
		if err := rows.Scan(&pelicula.PeliculaID, &pelicula.Titulo, &pelicula.Director, &pelicula.Duracion, &pelicula.AnioLanzamiento, &pelicula.Descripcion, &pelicula.Disponible, &pelicula.Genero); err != nil {
			log.Fatal("Error al escanear películas:", err)
		}
		peliculas = append(peliculas, pelicula)
	}

	for i, pelicula := range peliculas {
		fmt.Printf("%d. %s (%d)\n", i+1, pelicula.Titulo, pelicula.AnioLanzamiento)
		fmt.Printf("   Director: %s\n", pelicula.Director)
		fmt.Printf("   Duración: %d minutos\n", pelicula.Duracion)
		fmt.Printf("   Descripción: %s\n", pelicula.Descripcion)
		fmt.Printf("   Disponible: %t\n", pelicula.Disponible)
		fmt.Printf("   Género: %s\n", pelicula.Genero)
		fmt.Println()
	}

	// Inicializar la lista de películas seleccionadas
	catalogo.peliculasSeleccionadas = nil
	// Solicitar al usuario que seleccione películas por su número
	for {
		fmt.Print("Ingrese el número de la película que deseas seleccionar (0 para terminar la selección): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "0" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil || num < 1 || num > len(peliculas) {
			fmt.Println("Número de película no válido.")
			continue
		}

		pelicula := peliculas[num-1]

		// Agregar la película seleccionada al slice
		catalogo.peliculasSeleccionadas = append(catalogo.peliculasSeleccionadas, pelicula)
		fmt.Printf("Has seleccionado la película '%s'.\n", pelicula.Titulo)
	}

	// Mostrar la lista de películas seleccionadas
	fmt.Println("Películas seleccionadas:")
	for i, pelicula := range catalogo.peliculasSeleccionadas {
		fmt.Printf("%d. %s\n", i+1, pelicula.Titulo)
	}
}
