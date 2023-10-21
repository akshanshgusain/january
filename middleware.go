package january

import (
	"github.com/justinas/nosurf"
	"net/http"
	"strconv"
)

func (j *January) LoadSession(n http.Handler) http.Handler {
	j.InfoLog.Println("load session called")
	return j.Session.LoadAndSave(n)
}

func (j *January) JanuaryCSRF(n http.Handler) http.Handler {
	csrfHandler := nosurf.New(n)
	secure, _ := strconv.ParseBool(j.config.cookie.secure)

	csrfHandler.ExemptGlob("/api/*")

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   secure,
		SameSite: http.SameSiteStrictMode,
		Domain:   j.config.cookie.domain,
	})

	return csrfHandler
}
