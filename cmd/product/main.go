package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-foodshop/cmd/product/config"
	"go-foodshop/src/product/app"
	"go.uber.org/automaxprocs/maxprocs"
	"log/slog"
	"os"
)

func main() {
	_, err := maxprocs.Set()

	if err != nil {
		slog.Error("failed set max procs", err)
	}
	fmt.Println("Hello World")

	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	if err != nil {
		panic(err)
	}

	a, err := app.InitApp(config)

	if err != nil {
		panic(err)
	}
	err = a.StartApplication()
	if err != nil {
		panic(err)
	}
}
