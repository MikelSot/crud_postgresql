package storage

// link documentacion postgres para GO --> https://pkg.go.dev/github.com/lib/pq#section-documentation
import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	_ "github.com/lib/pq"
)

// variables globales  (estaran disponibles para todas las funciones y metodos de este paquete)
var (
	db *sql.DB  // Iimporta el paquete database/sql
	once sync.Once // a travez de estro es que se crea el SINGLETON
)

func NewPostgresDB() {
	once.Do(func() {
		var err error

		// lo que este dentro de esta funcion se va a ejecutar una sola VEZ
		db, err = sql.Open("postgres", "postgres://me-postgresql:cmd.08miguel@localhost:5432/go_db?sslmode=disable")
		if err != nil { // valida que los parametros sean los correctos
			log.Fatalf("no pudimos abrir la base de datos --> %v", err)
		}

		// db.Ping()  con este metodo es que vamos a validar recien si se hizo la conexion con la base de datos
		if err = db.Ping(); err != nil {
			log.Fatalf("no pudimos hacer ping --> %v", err)
		}

		fmt.Println("conextado a POSTGRESQL")
	})
}

// con esta funcion retorna una unica instancia de DB
func Pool() *sql.DB {
	return db
}