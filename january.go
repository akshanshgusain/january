package january

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

const version = "1.0.0"

type January struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	config   configuration
}

type configuration struct {
	port           string
	templateEngine string
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
	j.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	j.Version = version
	j.RootPath = rootPath

	// create loggers
	infoLog, errorLog := j.startLoggers()
	j.ErrorLog = errorLog
	j.InfoLog = infoLog

	// configuration
	j.config = configuration{
		port:           os.Getenv("PORT"),
		templateEngine: os.Getenv("TEMPLATE_ENGINE"),
	}
	return nil
}

func (j *January) Init(p initPaths) error {
	// root path of web app
	root := p.rootPath

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

func (j *January) startLoggers() (*log.Logger, *log.Logger) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}
