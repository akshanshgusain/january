package january

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"path"
	"path/filepath"
)

func (j *January) WriteJson(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (j *January) WriteXML(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := xml.MarshalIndent(data, "", "")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (j *January) DownloadFile(w http.ResponseWriter, r *http.Request, pathToFile, filename string) error {
	fp := path.Join(pathToFile, filename)
	fileToServer := filepath.Clean(fp)
	w.Header().Set("Content-Type", fmt.Sprintf("attachment; file=\"%s\"", filename))
	http.ServeFile(w, r, fileToServer)
	return nil
}

func (j *January) Error404(w http.ResponseWriter, r *http.Request) {
	j.ErrorStatus(w, r, http.StatusNotFound)
}

func (j *January) Error500(w http.ResponseWriter, r *http.Request) {
	j.ErrorStatus(w, r, http.StatusInternalServerError)
}

func (j *January) ErrorUnauthorised(w http.ResponseWriter, r *http.Request) {
	j.ErrorStatus(w, r, http.StatusUnauthorized)
}

func (j *January) ErrorForbidden(w http.ResponseWriter, r *http.Request) {
	j.ErrorStatus(w, r, http.StatusForbidden)
}

func (j *January) ErrorStatus(w http.ResponseWriter, r *http.Request, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}
