package auth

import (
	"Streamvibe/db"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

var userID int

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

type Genero struct {
	GeneroID int
	Genero   string
}

type Catalogo struct {
	peliculaSeleccionada   *Pelicula
	peliculasSeleccionadas []Pelicula
}

func (c *Catalogo) SetPeliculaSeleccionada(pelicula *Pelicula) {
	c.peliculaSeleccionada = pelicula
}

func (c *Catalogo) GetPeliculaSeleccionada() *Pelicula {
	return c.peliculaSeleccionada
}

var catalogo Catalogo

func NewUser() {
	var nombre, email, contraseña string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese nombre: ")
	scanner.Scan()
	nombre = scanner.Text()
	fmt.Print("Ingrese email: ")
	scanner.Scan()
	email = scanner.Text()
	fmt.Print("Ingrese contraseña: ")
	scanner.Scan()
	contraseña = scanner.Text()

	_, err := db.DB.Exec("INSERT INTO Clientes (Nombre, Email, Contraseña) VALUES (@p1, @p2, @p3)", nombre, email, contraseña)
	if err != nil {
		log.Fatal("Error al crear el usuario:", err)
	}
	fmt.Println("Usuario creado exitosamente")
}

func Login() int {
	var email, contraseña string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese email: ")
	scanner.Scan()
	email = scanner.Text()
	fmt.Print("Ingrese contraseña: ")
	scanner.Scan()
	contraseña = scanner.Text()

	var clienteID int
	err := db.DB.QueryRow("SELECT ClienteID FROM Clientes WHERE Email = @p1 AND Contraseña = @p2", email, contraseña).Scan(&clienteID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Usuario y/o contraseña incorrecto")
		} else {
			log.Fatal("Error al iniciar sesión:", err)
		}
		return 0
	}
	fmt.Printf("Inicio de sesión exitoso. ¡BIENVENIDO A STREAMVIBE!! ClienteID: %d\n", clienteID)
	return clienteID
}

func UserMenu() {
	for {
		fmt.Println("Selecciona la Opción deseada")
		fmt.Println("1. Catalogo de Películas\n2. Cerrar Sesión")
		fmt.Print("Ingrese su elección: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Println("Opción no válida:", err)
			continue
		}

		switch choice {
		case 1:
			ShowCatalog()
		case 2:
			fmt.Println("Sesión cerrada. ¡Hasta luego!")
			userID = 0
			return
		default:
			fmt.Println("Opción no válida")
		}
	}
}

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
		fmt.Print("Ingrese el número de la o las películas que deseas seleccionar (Ingresa N para terminar la selección): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "N" {
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

		// Preguntar al usuario si desea agregar más películas
		fmt.Print("¿Desea agregar otra película? (Sí: S / No: N): ")
		scanner.Scan()
		input = scanner.Text()
		if input == "N" {
			break
		}
	}

	// Mostrar la lista de películas seleccionadas
	fmt.Println("Películas seleccionadas:")
	for i, pelicula := range catalogo.peliculasSeleccionadas {
		fmt.Printf("%d. %s\n", i+1, pelicula.Titulo)
	}
}
