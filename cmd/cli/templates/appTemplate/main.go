package main

import (
	"github.com/akshanshgusain/january"
	"januaryApp/data"
	"januaryApp/handlers"
	"januaryApp/middleware"
)

type application struct {
	App        *january.January
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	j := initApplication()
	j.App.RunServer()
}
