package main

import (
	"fmt"
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

	err = copyDataToFile([]byte("drop table if exists users cascade"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	// run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}
	// copy files over

	return nil
}
