package january

import (
	"fmt"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

type January struct {
	AppName string
	Debug   bool
	Version string
}

func (j *January) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	if err := j.Init(pathConfig); err != nil {
		return err
	}

	if err := j.checkDotEnv(rootPath); err != nil {
		return err
	}
	// reading .env with extern lib: joho/godotenv
	if err := godotenv.Load(rootPath + "/.env"); err != nil {
		return err
	}

	return nil
}

func (j *January) Init(p initPaths) error {
	// root path of web app
	root := p.rootPath

	// TODO: get a list of all folder at the root path, create them if not exist
	for _, path := range p.folderNames {
		// create the folder if not present
		err := j.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (j *January) checkDotEnv(p string) error {
	if err := j.CreateFileIfNotExist(fmt.Sprintf("%s/.env", p)); err != nil {
		return err
	}
	return nil
}
