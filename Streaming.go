package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB
var userID int

type Pelicula struct {
	Titulo          string
	Genero          string
	Director        string
	Duracion        int
	AnioLanzamiento int
	Descripcion     string
	Disponible      bool
}

func main() {
	var err error

	// Conectar a la base de datos SQL Server

	db, err = sql.Open("sqlserver", "server=DESKTOP-05LFD9C;port=1433;database=Servicio")
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	// Interfaz de usuario
	for {
		fmt.Println("Selecciona la Opción deseada")
		fmt.Println("1. Crear Nuevo Usuario\n2. Iniciar Sesión\n3. Salir")
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
			newUser()
		case 2:
			userID = login()
			if userID != 0 {
				// Usuario inició sesión correctamente, mostrar menú adicional
				userMenu()
			}
		case 3:
			fmt.Println("Gracias por usar StreamVibe. ¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida")
		}
	}
}

func newUser() {
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

	// Insertar en la base de datos
	_, err := db.Exec("INSERT INTO Clientes (Nombre, Email, Contraseña) VALUES (@p1, @p2, @p3)", nombre, email, contraseña)
	if err != nil {
		log.Fatal("Error al crear el usuario:", err)
	}
	fmt.Println("Usuario creado exitosamente")
}

func login() int {
	var email, contraseña string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese email: ")
	scanner.Scan()
	email = scanner.Text()
	fmt.Print("Ingrese contraseña: ")
	scanner.Scan()
	contraseña = scanner.Text()

	// Verificar en la base de datos
	var clienteID int
	err := db.QueryRow("SELECT ClienteID FROM Clientes WHERE Email = @p1 AND Contraseña = @p2", email, contraseña).Scan(&clienteID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Usuario y/o contraseña incorrecto")
		} else {
			log.Fatal("Error al iniciar sesión:", err)
		}
		return 0
	}
	fmt.Printf("Inicio de sesión exitoso. ¡¡BIENVENIDO A STREAMVIBE!! ClienteID: %d\n", clienteID)
	return clienteID
}

func userMenu() {
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
			showCatalog()
		case 2:
			fmt.Println("Sesión cerrada. ¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida")
		}
	}
}

func showCatalog() {
	// Mostrar el catálogo de películas desde la base de datos (tabla Peliculas)
	rows, err := db.Query("SELECT Titulo, Genero, Director, Duracion, AnioLanzamiento, Descripcion, Disponible FROM Peliculas")
	if err != nil {
		log.Fatal("Error al obtener el catálogo de películas:", err)
	}
	defer rows.Close()

	peliculas := make([]Pelicula, 0)

	fmt.Println("Catálogo de Películas:")
	for rows.Next() {
		var pelicula Pelicula
		if err := rows.Scan(&pelicula.Titulo, &pelicula.Genero, &pelicula.Director, &pelicula.Duracion, &pelicula.AnioLanzamiento, &pelicula.Descripcion, &pelicula.Disponible); err != nil {
			log.Fatal("Error al escanear películas:", err)
		}
		peliculas = append(peliculas, pelicula)
	}

	// Mostrar detalles de las películas
	for i, pelicula := range peliculas {
		fmt.Printf("%d. %s (%d)\n", i+1, pelicula.Titulo, pelicula.AnioLanzamiento)
		fmt.Printf("   Género: %s\n", pelicula.Genero)
		fmt.Printf("   Director: %s\n", pelicula.Director)
		fmt.Printf("   Duración: %d minutos\n", pelicula.Duracion)
		fmt.Printf("   Descripción: %s\n", pelicula.Descripcion)
		fmt.Printf("   Disponible: %t\n", pelicula.Disponible)
		fmt.Println()
	}
}
