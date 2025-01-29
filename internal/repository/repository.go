package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"log"
)

var Database *sql.DB

func Repository() {
	// Connecting to the database
	connStr := "user=(here you describe your user name for the database) " +
		"password=(here you describe your password for the database) dbname=(here you specify " +
		"the name of your database) sslmode=(The sslmode parameter in the PostgreSQL " +
		"database connection string determines how the client application will establish " +
		"an SSL (Secure Sockets Layer) connection with the server. " +
		"we will use 'disable')"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	Database = db

	// Applying migrations
	log.Println("Starting migrations...")
	if err := goose.Up(db, "./migrations/database"); err != nil {
		log.Fatal(err)
	}
	log.Println("Migrations completed successfully.")
}
