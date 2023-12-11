package db

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlserver", "server=DESKTOP-05LFD9C;port=1433;database=Servicio")
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
}

func Close() {
	if err := DB.Close(); err != nil {
		log.Fatal("Error al cerrar la conexión de la base de datos:", err)
	}
}
