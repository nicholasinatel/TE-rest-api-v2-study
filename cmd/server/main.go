package main

import (
	"context"
	"errors"
	"fmt"
	"nicholasinatel/TE-rest-api-v2/internal/comment"
	"nicholasinatel/TE-rest-api-v2/internal/database"

	"github.com/golang-migrate/migrate/v4"
)

// SDG - To Him and only Him, all the glory
const SDG = "Soli Deo Gloria"

func init() {
	fmt.Println(SDG)
}

// Run - is going to be responsible for
// the instantiation and staretup of our go application
func Run() error {
	fmt.Println("Starting up our application")

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	if err := db.MigrateDB(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("could not run up migrations: %w", err)
		}
	}

	fmt.Println("successfully connected, pinged and migrated database")

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(
		context.Background(),
		"1",
	))

	return nil
}

func main() {
	fmt.Println("Hello from Docker within VSCode and Wsl2")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
