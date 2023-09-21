package main

import (
	"github.com/joho/godotenv"
	"os"
)

func setup() {
	err := godotenv.Load()
	if err != nil {
		exitGracefully(err)
	}
	path, err := os.Getwd()
	if err != nil {
		exitGracefully(err)
	}

	j.RootPath = path
	j.DB.DataType = os.Getenv("DATABASE_TYPE")
}
