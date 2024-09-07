package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strings"
)

func setup(arg1, arg2 string) {
	if arg1 != "new" && arg1 != "version" && arg1 != "help" {
		// irrespective of where the binary lives, if the current directory
		// has got the .env file it will work. current directory = root path
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
}

// re-formating the DSN for the golang-migration/migration tool
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
	// use spaces not tabs
	color.Magenta(`Available commands:

	help                  - show the help commands
	version               - print application version
	migrate up            - runs all up migrations that have not been run previously
	migrate down          - reverses the most recent migration
	migrate reset         - runs all down migrations in reverse order, and then all up migrations
	make migration <name> - creates two new up and down migrations in the migrations folder
	make auth             - creates and runs migrations for authentication tables, and creates models and middleware
	make handler <name>   - creates a stub handler in the handlers directory
	make model <name>     - creates a stub model in the data directory
	make session          - creates a tables in the database as a session store
	make key              - generates a 32 character random string
	make mail <name>      - create two starter mail templates in the mail directory
	`)
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
		color.Magenta(message)
	} else {
		color.Green("Finished!")
	}
	os.Exit(0)
}

func updateSourceFiles(path string, fi os.FileInfo, err error) error {
	// check for an error before doing anything else
	if err != nil {
		return err
	}

	// check if current file is directory
	if fi.IsDir() {
		return nil
	}

	// only check go files
	matched, err := filepath.Match("*.go", fi.Name())
	if err != nil {
		return err
	}

	// we have a matching file
	if matched {
		// read file contents
		read, err := os.ReadFile(path)
		if err != nil {
			exitGracefully(err)
		}

		newContents := strings.Replace(string(read), "januaryApp", appURL, -1)

		// write the changed file
		err = os.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			exitGracefully(err)
		}
	}

	return nil
}

func updateSource() {
	// walk entire project folder, including sub folders
	err := filepath.Walk(".", updateSourceFiles)
	if err != nil {
		exitGracefully(err)
	}
}
