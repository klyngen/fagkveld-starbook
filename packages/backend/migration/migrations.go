package migration

import (
	"database/sql"

	migrator "github.com/klyngen/mini-migrator"
)

var migrations = []migrator.Migration{
	{
		Name:        "Initial migration",
		Description: "Create the basic tables",
		Script: `
			CREATE TABLE person (
				id SERIAL PRIMARY KEY,
				user_id VARCHAR(50),
				personName VARCHAR(100),
				image VARCHAR(200));

			CREATE TABLE star (
				id SERIAL PRIMARY KEY,
				person_id INTEGER,
				description VARCHAR(500),
				user_id VARCHAR(50),
			FOREIGN KEY (person_id) REFERENCES person(id));
		`,
	},

	{
		Name:        "Column change",
		Description: "",
		Script:      `ALTER TABLE person ALTER COLUMN image TYPE VARCHAR(1000);`,
	},
}

func MigrateDatabase(db *sql.DB) error {
	var driver = migrator.PostgreSQLDriver
	driver.CreationQuery = `CREATE TABLE IF NOT EXISTS migrationTable (
		id INTEGER NOT NULL PRIMARY KEY,
		timestamp timestamp NOT NULL,
		description TEXT,
		name VARCHAR(50),
		hash VARCHAR(36),
		status INTEGER NOT NULL)`

	driver.WriteMigrationQuery = "INSERT INTO migrationTable  (id, timestamp, description, name, status, hash) VALUES($1, $2, $3, $4, $5, $6)"
	driver.UpdateMigrationStatusQuery = "UPDATE migrationTable SET status = $1 WHERE id = $2"

	migrator, err := migrator.NewMigrator(db, driver)

	if err != nil {
		return err
	}

	return migrator.MigrateDatabase(migrations)

}
