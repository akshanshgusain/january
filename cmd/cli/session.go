package main

import (
	"fmt"
	"time"
)

func doSessionTable() error {
	dbType := j.DB.DataType

	if dbType == "postgresql" {
		dbType = "postgres"
	}
	if dbType == "mariadb" {
		dbType = "mysql"
	}

	fileName := fmt.Sprintf("%d_create_sessions_table", time.Now().UnixMicro())

	upFile := j.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
	downFile := j.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

	// up migration
	err := copyFileFromTemplate("templates/migrations/"+dbType+"_session.sql", upFile)
	if err != nil {
		exitGracefully(err)
	}
	// down migration
	err = copyDataToFile([]byte("drop table sessions"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	return nil
}
