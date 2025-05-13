package shared_infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

func RunMigrations() {
	migrator, err := migrate.New(
		"file://db/migrations",
		os.Getenv("DATABASE_URL"),
	)

	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot load migrate: %s", err.Error()))
	}

	log.Println("Running migrations...")

	up_err := migrator.Up()

	if up_err != nil {
		if up_err == migrate.ErrNoChange {
			log.Println("No migrations to apply.")
		} else {
			log.Fatalf("Migration failed: %v", up_err)
		}
	} else {
		log.Println("Migrations ran successfully")
	}
}

func CreateDatabase() *sql.DB {
	connection_string := os.Getenv("DATABASE_URL")

	database, err := sql.Open("postgres", connection_string)

	if err != nil {
		log.Fatal(err)
	}

	return database
}
