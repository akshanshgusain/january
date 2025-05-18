package main

import (
	"fmt"
	"github.com/fatih/color"
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
		j.AppName = appName
	}
	log.Println("App name is: ", appName)

	// clone the started app from the template directory
	color.Green("\tbuilding starter app...")

	err := os.MkdirAll("./"+appName, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating app directory:", err)
		return
	}

	err = copyFilesFromDir("templates/appTemplate", "./"+appName)
	if err != nil {
		exitGracefully(err)
	}

	//create a .gitignore file
	color.Yellow("\tcreating .gitignore file")
	data, err := templateFS.ReadFile("templates/gitignore.txt")
	if err != nil {
		exitGracefully(err)
	}
	env := string(data)
	err = copyDataToFile([]byte(env), fmt.Sprintf("./%s/.gitignore", appName))
	if err != nil {
		exitGracefully(err)
	}

	//create a .env file
	color.Yellow("\tcreating .env file")
	data, err = templateFS.ReadFile("templates/env.txt")
	if err != nil {
		exitGracefully(err)
	}

	env = string(data)
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
	// remove unused Makefile templates from the skeleton
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
