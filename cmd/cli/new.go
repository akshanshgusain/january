package main

import (
	"github.com/fatih/color"
	"log"
	"strings"
)

func doNew(appName string) {
	appName = strings.ToLower(appName)

	//TODO: sanitise the application name
	if strings.Contains(appName, "/") {
		exploded := strings.SplitAfter(appName, "/")
		appName = exploded[(len(exploded) - 1)]
	}
	log.Println("App name is: ", appName)

	//TODO: git clone the skeleton application
	color.Green("\tcloning repository...")

	//TODO: remove the .gitignore

	// TODO: create a .env file

	//TODO: update the go.mod file

	//TODO: update existing .go files with correct name/imports

	// TODO: run go mod tidy
}
