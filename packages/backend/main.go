package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"

	"github.com/klyngen/packages/backend/starbook-auth/migration"
	"github.com/klyngen/packages/backend/starbook-auth/presentation"
	"github.com/klyngen/packages/backend/starbook-auth/repository"
)

func main() {
	connectionString := os.Getenv("STARBOOK_DATABASE")
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic("Unable to connect to a database, making this app useless: " + err.Error())
	}

	err = migration.MigrateDatabase(db)

	if err != nil {
		panic("Unable to connect to migrate the database, making this app useless: " + err.Error())
	}

	repository := repository.NewRepository(db)

	api := presentation.NewApi(repository, presentation.AuthenticationConfig{})

	api.Serve("1337")
}
