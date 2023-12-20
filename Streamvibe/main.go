package main

import (
	"Streamvibe/auth"
	"Streamvibe/db"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Estructuras para deserializar la respuesta JSON
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type GenreList struct {
	Genres []Genre `json:"genres"`
}

func main() {

	//DB
	db.Init()
	defer db.Close()

	var userID int
	for {
		fmt.Println("Selecciona la Opción deseada")
		fmt.Println("1. Crear Nuevo Usuario")
		fmt.Println("2. Iniciar Sesión")
		fmt.Println("3. Salir")
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
			auth.NewUser()
		case 2:
			userID = auth.Login()
			if userID != 0 {
				auth.UserMenu()
			}
		case 3:
			fmt.Println("Gracias por usar StreamVibe. ¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida")
		}
	}
}
