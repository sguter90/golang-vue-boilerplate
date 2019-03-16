package app

import (
	"errors"
	"fmt"
	"github.com/rubenv/sql-migrate"
	"github.com/sguter90/golang-vue-boilerplate/internal/app/repository"
	"strings"
)

func (app *App) Migrate(params []string) error {
	migrateDirection := migrate.Up

	action := "help"
	if len(params) > 0 {
		action = strings.ToLower(params[0])
	}

	switch action {
	case "up":
		migrateDirection = migrate.Up
	case "down":
		migrateDirection = migrate.Down
	case "help":
		output := `
    Usage: 
        migrate up      Migrates the database to the most recent version available
        migrate down   Undo a database migration
			`
		fmt.Println(output)
		return nil
	default:
		return errors.New(fmt.Sprintf("Invalid action '%s'", action))
	}
	if action == "down" {
		migrateDirection = migrate.Down
	}

	dbx, dialect, err := repository.ConnectDb()
	if err != nil {
		return err
	}

	migrationFolder := &migrate.FileMigrationSource{
		Dir: "migration/" + dialect,
	}

	db := dbx.DB
	numberApplied, migErr := migrate.Exec(db, dialect, migrationFolder, migrateDirection)
	if migErr != nil {
		return migErr
	}

	fmt.Printf("Applied %d migration(s)\n", numberApplied)

	return nil
}
