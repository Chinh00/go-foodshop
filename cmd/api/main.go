package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	config2 "go-foodshop/cmd/api/config"
	app2 "go-foodshop/src/app"
	"go-foodshop/src/pkg/database"
)

func main() {
	fmt.Println("Hello World")
	e := echo.New()
	config, err := config2.NewConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}
	db := &database.SqlDatabase{}
	err = app2.NewApplication(db)

	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	e.Start(address)
}
