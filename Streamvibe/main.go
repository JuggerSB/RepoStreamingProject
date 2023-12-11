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

func main() {
	db.Init()
	defer db.Close()

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
			auth.NewUser()
		case 2:
			userID := auth.Login()
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
