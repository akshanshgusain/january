package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
)

// extractRequireBlock extracts the require block from a go.mod file content
func extractRequireBlock(modContent string) string {
	lines := strings.Split(modContent, "\n")
	var requireBlock []string
	inRequireBlock := false

	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "require (") {
			inRequireBlock = true
			continue
		}
		if inRequireBlock {
			if strings.TrimSpace(line) == ")" {
				break
			}
			requireBlock = append(requireBlock, line)
		}
	}

	return strings.Join(requireBlock, "\n")
}

func doAuthJWT() error {
	// extract the appName from go.mod
	appName = extractAppName()
	appURL = appName

	// create migrations
	dbType := j.DB.DataType

	fileName := fmt.Sprintf("%d_create_auth_jwt_tables", time.Now().UnixMicro())
	upFile := j.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := j.RootPath + "/migrations/" + fileName + ".down.sql"

	err := copyFileFromTemplate("templates/migrations/auth_tables_jwt."+dbType+".sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile([]byte("drop table if exists users cascade; drop table if exists tokens_jwt cascade;"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	// run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	// copy files over
	err = copyFileFromTemplate("templates/data/userJWT.go.txt", j.RootPath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/tokenJWT.go.txt", j.RootPath+"/data/tokenJWT.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy middlewares

	err = copyFileFromTemplate("templates/middleware/authToken.go.txt", j.RootPath+"/middleware/authToken.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy handlers
	err = copyFileFromTemplate("templates/handlers/authJWTHandlers.go.txt", j.RootPath+"/handlers/authJWTHandlers.go")
	if err != nil {
		exitGracefully(err)
	}

	// add handlerHelper function to ge the decoded token from request
	err = appendFileFromTemplate("templates/handlers/authJWT-handlerHelper.go.txt", j.RootPath+"/handlers/handlerHelper.go")
	if err != nil {
		exitGracefully(err)
	}

	// handle go.mod file
	color.Yellow("\tadding JWT dependency...")

	cmd := exec.Command("go", "get", "github.com/golang-jwt/jwt/v5")
	err = cmd.Run()
	if err != nil {
		exitGracefully(err)
	}

	color.Green("changing auth")
	updateSource()

	color.Yellow("   - users and token_jwt migrations created and executed")
	color.Yellow("   - user and token_jwt models created")
	color.Yellow("   - auth middleware created")
	color.Yellow("")
	color.Yellow("Don't forget to add user and token_jwt models in data/models.go, and to add appropriate middleware to your routes!")

	return nil
}
