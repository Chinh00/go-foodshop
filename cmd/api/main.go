package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	config2 "go-foodshop/cmd/api/config"
	app2 "go-foodshop/src/app"
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

	config, err := config2.NewConfig()
	if err != nil {
		panic(err)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	app, err := app2.InitApp(config)
	//consumer := consumer.NewConsumer()
	//go func() {
	//	consumer.StartConsumer()
	//}()
	if err != nil {
		panic(err)
	}
	err = app.StartApplication()
	if err != nil {
		panic(err)
	}
}
