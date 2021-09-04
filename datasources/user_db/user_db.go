package user_db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	postgres_user   = "postgres"
	postgres_pass   = "postgres"
	postgres_ip     = "localhost"
	postgres_dbname = "bookstore"
)

var (
	Client *sql.DB
	// user   = os.Getenv(postgres_user)
	// pass   = os.Getenv(postgres_pass)
	// ip     = os.Getenv(postgres_ip)
	// db     = os.Getenv(postgres_dbname)
)

func init() {
	var err error

	connStr := "postgres://" + postgres_user + ":" + postgres_pass + "@" + postgres_ip + "/" + postgres_dbname + "?sslmode=disable"

	Client, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	// this will be printed in the terminal, notifying that we successfully connected to our database
	log.Println("you are now connected to the database.")
}
