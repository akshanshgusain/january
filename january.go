package january

import (
	"database/sql"
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/akshanshgusain/january/cache"
	"github.com/akshanshgusain/january/mailer"
	"github.com/alexedwards/scs/v2"
	"github.com/dgraph-io/badger/v3"
	"github.com/go-chi/chi/v5"
	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const version = "1.0.0"

var myRedisCache *cache.RedisCache
var myBadgerCache *cache.BadgerCache
var redisPool *redis.Pool
var badgerConn *badger.DB

type January struct {
	AppName        string
	Debug          bool
	Version        string
	ErrorLog       *log.Logger
	InfoLog        *log.Logger
	RootPath       string
	Routes         *chi.Mux
	TemplateEngine *TemplateEngine
	Session        *scs.SessionManager
	DB             Database
	JetViews       *jet.Set
	config         configuration
	EncryptionKey  string
	Cache          cache.Cache
	Scheduler      *cron.Cron
	Mail           mailer.Mail
	Server         Server
}

type configuration struct {
	port           string
	templateEngine string
	cookie         cookieConfig
	sessionType    string
	database       databaseConfig
	redis          redisConfig
}

type Server struct {
	ServerName string
	Port       string
	Secure     bool
	URL        string
}

func (j *January) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "mail", "data", "public", "tmp", "logs", "middleware"},
	}

	if err := j.Init(pathConfig); err != nil {
		return err
	}

	// TODO: add a bette way to do sanity check

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

	// create mailer
	j.Mail = j.createMailer()

	// check if redis is available
	if os.Getenv("CACHE") == "redis" || os.Getenv("SESSION_TYPE") == "redis" {
		myRedisCache = j.createClientRedisCache()
		j.Cache = myRedisCache
		redisPool = myRedisCache.Conn
	}

	// set scheduler
	scheduler := cron.New()
	j.Scheduler = scheduler

	// check if badger is available
	if os.Getenv("CACHE") == "badger" {
		myBadgerCache = j.createClientBadgerCache()
		j.Cache = myBadgerCache
		badgerConn = myBadgerCache.Conn

		// garbage collection/clean up
		_, err := j.Scheduler.AddFunc("@daily", func() {
			_ = myBadgerCache.Conn.RunValueLogGC(0.7)
		})
		if err != nil {
			return err
		}
	}

	// connect to database
	if os.Getenv("DATABASE_TYPE") != "" {
		db, err := j.OpenDBConnection(os.Getenv("DATABASE_TYPE"), j.BuildDSN())
		if err != nil {
			errorLog.Println(err)
			os.Exit(1)
		}
		j.DB = Database{
			DataType: os.Getenv("DATABASE_TYPE"),
			Pool:     db,
		}
	}

	// configuration
	j.config = configuration{
		port:           os.Getenv("PORT"),
		templateEngine: os.Getenv("TEMPLATE_ENGINE"),
		cookie: cookieConfig{
			name:     os.Getenv("COOKIE_NAME"),
			lifetime: os.Getenv("COOKIE_LIFETIME"),
			persist:  os.Getenv("COOKIE_PERSIST"),
			secure:   os.Getenv("COOKIE_SECURE"),
			domain:   os.Getenv("COOKIE_DOMAIN"),
		},
		sessionType: os.Getenv("SESSION_TYPE"),
		database: databaseConfig{
			database: os.Getenv("DATABASE_TYPE"),
			dsn:      j.BuildDSN(),
		},
		redis: redisConfig{
			host:     os.Getenv("REDIS_HOST"),
			password: os.Getenv("REDIS_PASSWORD"),
			prefix:   os.Getenv("REDIS_PREFIX"),
		},
	}

	// fill in server details
	secure := true
	if strings.ToLower(os.Getenv("SECURE")) == "false" {
		secure = false
	}

	j.Server = Server{
		ServerName: os.Getenv("SERVER_NAME"),
		Port:       os.Getenv("PORT"),
		Secure:     secure,
		URL:        os.Getenv("APP_URL"),
	}

	// create session
	s := Session{
		CookieLifetime: j.config.cookie.lifetime,
		CookiePersist:  j.config.cookie.persist,
		CookieName:     j.config.cookie.name,
		SessionType:    j.config.sessionType,
		CookieDomain:   j.config.cookie.domain,
	}

	// switch session storage
	switch j.config.sessionType {
	case "redis":
		s.RedisPool = myRedisCache.Conn
	case "mysql", "postgres", "mariadb", "postgresql":
		s.DBPool = j.DB.Pool
	}

	j.Session = s.InitSession()
	j.EncryptionKey = os.Getenv("KEY")

	// add routes
	j.Routes = j.routes().(*chi.Mux)

	// jet views
	if j.Debug {
		j.JetViews = jet.NewSet(
			jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
			jet.InDevelopmentMode(),
		)
	} else {
		j.JetViews = jet.NewSet(
			jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		)
	}

	// add Template Engine
	j.createTemplateEngine()

	// mailer-background task: go routine to listen to mails in the background
	go j.Mail.ListenForMail()

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

func (j *January) createTemplateEngine() {
	j.TemplateEngine = &TemplateEngine{
		TemplateEngine: j.config.templateEngine,
		RootPath:       j.RootPath,
		Port:           j.config.port,
		JetViews:       j.JetViews,
		Session:        j.Session,
	}
}

func (j *January) createMailer() mailer.Mail {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	m := mailer.Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Templates:   j.RootPath + "/mail",
		Host:        os.Getenv("SMTP_HOST"),
		Port:        port,
		Username:    os.Getenv("SMTP_USERNAME"),
		Password:    os.Getenv("SMTP_PASSWORD"),
		Encryption:  os.Getenv("SMTP_ENCRYPTION"),
		FromName:    os.Getenv("FROM_NAME"),
		FromAddress: os.Getenv("FROM_ADDRESS"),

		// TODO: figure out a way to dynamically change the channel size
		Jobs:    make(chan mailer.Message, 20), // 20 Jobs at a time
		Results: make(chan mailer.Result, 20),
		API:     os.Getenv("MAILER_API"),
		APIKey:  os.Getenv("MAILER_KEY"),
		APIUrl:  os.Getenv("MAILER_URL"),
	}
	return m
}

func (j *January) RunServer() {
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     j.ErrorLog,
		Handler:      j.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// TODO: close db connection
	defer func(Pool *sql.DB) {
		err := Pool.Close()
		if err != nil {
			j.ErrorLog.Fatal(err)
		}
	}(j.DB.Pool)

	defer func(redisPool *redis.Pool) {
		err := redisPool.Close()
		if err != nil {

		}
	}(redisPool)

	defer func(badgerConn *badger.DB) {
		err := badgerConn.Close()
		if err != nil {

		}
	}(badgerConn)

	j.InfoLog.Printf("Starting January server at http://127.0.0.1:%s/", os.Getenv("PORT"))
	j.InfoLog.Printf("Quit the server with control+c")
	if err := s.ListenAndServe(); err != nil {
		j.ErrorLog.Fatal(err)
	}

}

func (j *January) BuildDSN() string {
	var dsn string
	switch os.Getenv("DATABASE_TYPE") {
	case "postgres", "postgresql":
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_NAME"),
			os.Getenv("DATABASE_SSL_MODE"))
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("%s password=%s",
				dsn,
				os.Getenv("DATABASE_PASS"))
		}
	default:
	}

	return dsn
}

func (j *January) createRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   10000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				j.config.redis.host,
				redis.DialPassword(j.config.redis.password),
			)
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func (j *January) createClientRedisCache() *cache.RedisCache {
	cacheClient := cache.RedisCache{
		Conn:   j.createRedisPool(),
		Prefix: j.config.redis.prefix,
	}

	return &cacheClient
}

func (j *January) createBadgerConn() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions(j.RootPath + "/tmp/cache/badger"))
	if err != nil {
		return nil
	}
	return db
}

func (j *January) createClientBadgerCache() *cache.BadgerCache {
	cacheClient := cache.BadgerCache{
		Conn: j.createBadgerConn(),
	}

	return &cacheClient
}
