package main

import (
	"errors"
	"github.com/fatih/color"
)

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()

	// run the migration command
	switch arg2 {
	case "up":
		err := j.MigrateUp(dsn)
		if err != nil {
			return err
		}

	case "down":
		if arg3 == "all" {
			err := j.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := j.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}
	case "reset":
		err := j.MigrateDownAll(dsn)
		if err != nil {
			return err
		}
		err = j.MigrateUp(dsn)
		if err != nil {
			return err
		}
	default:
		showMigrateHelp()
		return errors.New("invalid migrate command")
	}

	return nil
}

func showMigrateHelp() {
	// use spaces not tabs
	color.Magenta(`Available migrate <sub commands>:

	up                          - runs all up migrations that have not been run previously
	down                        - reverses the most recent migration
	reset                       - runs all down migrations in reverse order, and then all up migrations
	`)
}
