package january

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"path/filepath"
)

func (j *January) ReadJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576 // one megabyte
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only have a single json value")
	}

	return nil
}

// WriteJson Marshals data
// Sets the appropriate headers
// write the marshaled data to http.ResponseWriter
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

// WriteXML Marshals data
// Sets the appropriate headers
// write the marshaled data to http.ResponseWriter
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
	//w.Header().Set("Content-Type", fmt.Sprintf("attachment; file=\"%s\"", filename))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; file=\"%s\"", filename))
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
