package main

import (
	"errors"
	"fmt"
	"time"
)

func doMake(arg2, arg3 string) error {
	switch arg2 {
	case "migration":
		dbType := j.DB.DataType
		if arg3 == "" {
			exitGracefully(errors.New("you must give this migration a name"))
		}
		fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3)
		upFile := j.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
		downFile := j.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

		//	TODO: templates for migrations

		// up migration
		err := copyFileFromTemplate("templates/migrations/migration."+dbType+".up.sql", upFile)
		if err != nil {
			exitGracefully(err)
		}
		// down migration
		err = copyFileFromTemplate("templates/migrations/migration."+dbType+".down.sql", downFile)
		if err != nil {
			exitGracefully(err)
		}
	case "auth":
		err := doAuth()
		if err != nil {
			exitGracefully(err)
		}
	}
	return nil
}
