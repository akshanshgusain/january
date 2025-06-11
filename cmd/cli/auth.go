package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
	"time"
)

var appName string

func extractAppName() string {
	file, err := os.Open("go.mod")
	if err != nil {
		exitGracefully(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			moduleName := strings.TrimSpace(strings.TrimPrefix(line, "module"))
			return moduleName
		}
	}

	if err := scanner.Err(); err != nil {
		exitGracefully(err)
	} else {
		exitGracefully(fmt.Errorf("No module declaration found in go.mod"))
	}

	return ""
}

func doAuth() error {
	// extract the appName from go.mod
	appName = extractAppName()
	appURL = appName

	// create migrations
	dbType := j.DB.DataType

	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := j.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := j.RootPath + "/migrations/" + fileName + ".down.sql"

	err := copyFileFromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile([]byte("drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens;"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	// run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	// copy files over
	err = copyFileFromTemplate("templates/data/user.go.txt", j.RootPath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/token.go.txt", j.RootPath+"/data/token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/rememberToken.go.txt", j.RootPath+"/data/remember_token.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy middlewares
	err = copyFileFromTemplate("templates/middleware/auth.go.txt", j.RootPath+"/middleware/auth.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/authToken.go.txt", j.RootPath+"/middleware/authToken.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/remember.go.txt", j.RootPath+"/middleware/remember.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy handlers
	err = copyFileFromTemplate("templates/handlers/authHandlers.go.txt", j.RootPath+"/handlers/authHandlers.go")
	if err != nil {
		exitGracefully(err)
	}

	color.Green("changing auth")
	updateSource()

	//data, err := templateFS.ReadFile("templates/handlers/authHandlers.go.txt")
	//if err != nil {
	//	exitGracefully(err)
	//}
	//handler := string(data)
	//handler = strings.ReplaceAll(handler, "$APP_NAME$", appURL)
	//err = ioutil.WriteFile(j.RootPath+"/handlers/authHandlers.go", []byte(handler), 0644)
	//if err != nil {
	//	exitGracefully(err)
	//}

	// copy views

	// views-mailer
	err = copyFileFromTemplate("templates/mailer/password-reset.html.tmpl", j.RootPath+"/mail/password-reset.html.tmpl")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/mailer/password-reset.plain.tmpl", j.RootPath+"/mail/password-reset.plain.tmpl")
	if err != nil {
		exitGracefully(err)
	}

	// views - views
	err = copyFileFromTemplate("templates/views/login.jet", j.RootPath+"/views/login.html")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/views/forgot.jet", j.RootPath+"/views/forgot.html")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/views/reset-password.jet", j.RootPath+"/views/reset-password.html")
	if err != nil {
		exitGracefully(err)
	}

	//updateSource()

	color.Yellow("   - users, tokens, and remember_tokens migrations created and executed")
	color.Yellow("   - user and tokens models created")
	color.Yellow("   - auth middleware created")
	color.Yellow("")
	color.Yellow("Don't forget to add user, token and remember_token models in data/models.go, and to add appropriate middleware to your routes!")

	return nil
}
