package app

import "go-foodshop/src/pkg/database"

type App struct {
	dbContext database.Database
}

func NewApplication(dbContext database.Database) error {
	app := &App{}
	app.dbContext = dbContext
	return nil
}
