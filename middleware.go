package january

import "net/http"

func (j *January) LoadSession(n http.Handler) http.Handler {
	j.InfoLog.Println("load session called")
	return j.Session.LoadAndSave(n)
}
