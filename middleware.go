package january

import "net/http"

func (j *January) LoadSession(n http.Handler) http.Handler {
	return j.Session.LoadAndSave(n)
}
