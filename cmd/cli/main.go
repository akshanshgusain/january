package main

import (
	"errors"
	"github.com/akshanshgusain/january"
	"github.com/fatih/color"
	"os"
)

const version = "1.0.0"

var j january.January

func main() {
	var message string
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}

	setup(arg1, arg2)

	switch arg1 {
	case "help":
		showHelp()
	case "new":
		if arg2 == "" {
			exitGracefully(errors.New("new requires an application name"))
		}
		doNew(arg2)
	case "version":
		color.Cyan("Application version: " + version)
	case "make":
		if arg2 == "" {
			exitGracefully(errors.New("make requires a subcommand: (migration | model | handler)"))
		}
		err = doMake(arg2, arg3)
		if err != nil {
			exitGracefully(err)
		}
	case "migrate":
		err = doMigrate(arg2, arg3)
		if err != nil {
			exitGracefully(err)
		}
		message = "Migrations complete!"
	default:
		showHelp()
	}

	exitGracefully(nil, message)
}

func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string

	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) > 2 {
			arg2 = os.Args[2]
		}

		if len(os.Args) > 3 {
			arg3 = os.Args[3]
		}
	} else {
		color.Red("Error: command required")
		showHelp()
		return "", "", "", errors.New("command required")
	}

	return arg1, arg2, arg3, nil
}
