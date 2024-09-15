package main

import (
	"github.com/akshanshgusain/january"
	"januaryApp/data"
	"januaryApp/handlers"
	"januaryApp/middleware"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// init january
	jan := &january.January{}

	err = jan.New(path)
	if err != nil {
		log.Fatal(err)
	}

	jan.AppName = "jan_app"

	mid := &middleware.Middleware{
		App: jan,
	}

	// init handlers
	h := &handlers.Handlers{
		App: jan,
	}

	app := &application{
		App:        jan,
		Handlers:   h,
		Middleware: mid,
	}

	// init Routes
	app.App.Routes = app.routes()

	// init Models
	app.Models = data.New(app.App.DB.Pool)
	h.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
