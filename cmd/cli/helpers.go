package main

import (
	"fmt"
	"github.com/fatih/color"
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

func getDSN() string {
	dbType := j.DB.DataType

	if dbType == "pgx" {
		dbType = "postgres"
	}

	if dbType == "postgres" {
		var dsn string
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		}
		return dsn
	}
	return "mysql://" + j.BuildDSN()
}

func showHelp() {
	color.Yellow(`Available commands: 
	help					-	show the help commands
	version					-	print application version
	migrate					-	run all up migrations that have not been run previously
	migrate down			-	reverses the most recent migration
	migrate reset			-	runs all down migrations in reverse order, and then all up migrations
	make migration <name>	-	creates two new up and down migrations in the migrations folder`)
}

func exitGracefully(e error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}
	if e != nil {
		color.Red("Error: %v\n", e)
	}
	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished!")
	}
	os.Exit(0)
}
