package january

import "database/sql"

type initPaths struct {
	rootPath    string
	folderNames []string
}

type cookieConfig struct {
	name     string
	lifetime string
	persist  string
	secure   string
	domain   string
}

type databaseConfig struct {
	dsn      string
	database string
}

type Database struct {
	DataType string
	Pool     *sql.DB
}
