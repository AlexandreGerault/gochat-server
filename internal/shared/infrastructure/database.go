package shared_infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	upErr := migrator.Up()

	if upErr != nil {
		if upErr == migrate.ErrNoChange {
			log.Println("No migrations to apply.")
		} else {
			log.Fatalf("Migration failed: %v", upErr)
		}
	} else {
		log.Println("Migrations ran successfully")
	}
}
