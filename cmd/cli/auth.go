package main

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

func doAuth() error {
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

	// copy middlewares
	err = copyFileFromTemplate("templates/middleware/auth.go.txt", j.RootPath+"/middleware/auth.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/authToken.go.txt", j.RootPath+"/middleware/authToken.go")
	if err != nil {
		exitGracefully(err)
	}

	color.Yellow("   - users, tokens, and remember_tokens migrations created and executed")
	color.Yellow("   - user and tokens models created")
	color.Yellow("   - auth middleware created")
	color.Yellow("")
	color.Yellow("Don't forget to add user and token models in data/models.go, and to add appropriate middleware to your routes!")

	return nil
}
