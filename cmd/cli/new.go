package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var appURL string

func doNew(appName string) {
	appName = strings.ToLower(appName)
	appURL = appName

	//sanitise the application name
	if strings.Contains(appName, "/") {
		exploded := strings.SplitAfter(appName, "/")
		appName = exploded[(len(exploded) - 1)]
	}
	log.Println("App name is: ", appName)

	//git clone the skeleton application
	color.Green("\tcloning repository...")
	_, err := git.PlainClone("./"+appName, false, &git.CloneOptions{
		URL:      "https://github.com/akshanshgusain/january_starter_app",
		Progress: os.Stdout,
		Depth:    1,
	})

	if err != nil {
		exitGracefully(errors.New("error cloning the skeleton project: "))
		exitGracefully(err)
	}

	//remove the .gitignore
	err = os.RemoveAll(fmt.Sprintf("./%s/.git", appName))
	if err != nil {
		exitGracefully(err)
	}

	//create a .env file
	color.Yellow("\tcreating .env file")
	data, err := templateFS.ReadFile("templates/env.txt")
	if err != nil {
		exitGracefully(err)
	}

	env := string(data)
	env = strings.ReplaceAll(env, "${APP_NAME}", appName)
	env = strings.ReplaceAll(env, "${KEY}", j.RandomString(32))

	err = copyDataToFile([]byte(env), fmt.Sprintf("./%s/.env", appName))
	if err != nil {
		exitGracefully(err)
	}

	// creat a Makefile
	if runtime.GOOS == "windows" {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.windows", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer source.Close()

		destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			exitGracefully(err)
		}
	} else {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.mac", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer source.Close()

		destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			exitGracefully(err)
		}
	}
	_ = os.Remove("./" + appName + "/Makefile.mac")
	_ = os.Remove("./" + appName + "/Makefile.windows")

	// update the go.mod file

	color.Yellow("\tcreating go.mod file...")
	_ = os.Remove("./" + appName + "/go.mod")

	data, err = templateFS.ReadFile("templates/go.mod.txt")
	if err != nil {
		exitGracefully(err)
	}
	mod := string(data)
	mod = strings.ReplaceAll(mod, "${APP_NAME}", appURL)

	err = copyDataToFile([]byte(mod), "./"+appName+"/go.mod")
	if err != nil {
		exitGracefully(err)
	}

	//update existing .go files with correct name/imports
	color.Yellow("\tupdating source files...")
	os.Chdir("./" + appName)
	updateSource()

	//  run go mod tidy
	color.Yellow("\trunning go mod tidy...")
	cmd := exec.Command("go", "mod", "tidy")
	err = cmd.Start()
	if err != nil {
		exitGracefully(err)
	}

	color.Green("done building " + appURL)
	color.Green("go build something awesome!")
}
