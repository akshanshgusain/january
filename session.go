package january

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	SessionType    string
	CookieSecure   string
}

func (s *Session) InitSession() *scs.SessionManager {
	var persist = false
	var secure = false

	// session duration
	minutes, err := strconv.Atoi(s.CookieLifetime)
	if err != nil {
		minutes = 60
	}

	// should cookie persist?
	if strings.ToLower(s.CookiePersist) == "true" {
		persist = true
	}

	// should cookie be secure?
	if strings.ToLower(s.CookieSecure) == "true" {
		secure = true
	}

	// create session
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = s.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = s.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// select session store
	switch strings.ToLower(s.SessionType) {
	case "redis":
	case "mysql", "mariadb":
	case "postgres", "postgresql":
	default:
	}

	return session
}
