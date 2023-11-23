package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"log"
	"os"
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
	_, err := git.PlainClone("./"+appName, false, &git.CloneOptions{
		URL:      "https://github.com/akshanshgusain/january_app_skl",
		Progress: os.Stdout,
		Depth:    1,
	})

	if err != nil {
		exitGracefully(errors.New("error cloning the skeleton project: "))
		exitGracefully(err)
	}

	//TODO: remove the .gitignore
	err = os.RemoveAll(fmt.Sprintf("./%s/.git", appName))
	if err != nil {
		exitGracefully(err)
	}

	// TODO: create a .env file

	//TODO: update the go.mod file

	//TODO: update existing .go files with correct name/imports

	// TODO: run go mod tidy
}
