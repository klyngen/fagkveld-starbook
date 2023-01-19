package main

import (
	"database/sql"
	"os"

	"github.com/klyngen/fagkveld-starbook/packages/backend/starbook-auth/migration"
	"github.com/klyngen/fagkveld-starbook/packages/backend/starbook-auth/presentation"
	"github.com/klyngen/fagkveld-starbook/packages/backend/starbook-auth/repository"
	_ "github.com/lib/pq"
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

	api := presentation.NewApi(repository, presentation.AuthenticationConfig{
		Domain: "https://sts.windows.net/76749190-4427-4b08-a3e4-161767dd1b73",
	})

	api.Serve("1337")
}
