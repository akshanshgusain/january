package january

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

// MigrateUp run migrations which have not yet been run
func (j *January) MigrateUp(dsn string) error {
	m, err := migrate.New("file://"+j.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()
	if err = m.Up(); err != nil {
		log.Println("error running migration: ", err)
		return err
	}
	return nil
}

func (j *January) MigrateDown(dsn string) error {
	m, err := migrate.New("file://"+j.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()
	if err = m.Down(); err != nil {
		log.Println("error running migration: ", err)
		return err
	}
	return nil
}

func (j *January) Steps(n int, dsn string) error {
	m, err := migrate.New("file://"+j.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Steps(n); err != nil {
		return err
	}
	return nil
}

func (j *January) MigrationForce(dsn string) error {
	m, err := migrate.New("file://"+j.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()
	if err = m.Force(-1); err != nil {
		return err
	}
	return nil
}

func (j *January) MigrateDownAll(dsn string) error {
	m, err := migrate.New("file://"+j.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Down(); err != nil {
		return err
	}

	return nil
}
