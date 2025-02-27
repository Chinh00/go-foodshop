package main

import (
	"fmt"
	config2 "go-foodshop/cmd/api/config"
	app2 "go-foodshop/src/app"
)

func main() {
	fmt.Println("Hello World")

	config, err := config2.NewConfig()
	if err != nil {
		panic(err)
	}
	app, err := app2.InitApp(config)
	if err != nil {
		panic(err)
	}
	err = app.StartApplication()
	if err != nil {
		panic(err)
	}
}
