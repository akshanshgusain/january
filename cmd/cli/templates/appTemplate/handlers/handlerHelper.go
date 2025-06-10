package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/akshanshgusain/january"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"os"
	"regexp"
)

func (h *Handlers) render(w http.ResponseWriter, r *http.Request, tmpl string, variables, data interface{}) error {
	return h.App.TemplateEngine.Page(w, r, tmpl, variables, data)
}

func (h *Handlers) sessionPut(ctx context.Context, key string, val interface{}) {
	h.App.Session.Put(ctx, key, val)
}

func (h *Handlers) sessionHas(ctx context.Context, key string) bool {
	return h.App.Session.Exists(ctx, key)
}

func (h *Handlers) sessionGet(ctx context.Context, key string) interface{} {
	return h.App.Session.Get(ctx, key)
}

func (h *Handlers) sessionRemove(ctx context.Context, key string) {
	h.App.Session.Remove(ctx, key)
}

func (h *Handlers) sessionRenew(ctx context.Context) error {
	return h.App.Session.RenewToken(ctx)
}

func (h *Handlers) sessionDestroy(ctx context.Context) error {
	return h.App.Session.Destroy(ctx)
}

func (h *Handlers) randomString(n int) string {
	return h.App.RandomString(n)
}

func (h *Handlers) encrypt(text string) (string, error) {
	e := january.Encryption{Key: []byte(h.App.EncryptionKey)}

	encrypted, err := e.Encrypt(text)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func (h *Handlers) decrypt(cypher string) (string, error) {
	e := january.Encryption{Key: []byte(h.App.EncryptionKey)}

	encrypted, err := e.Decrypt(cypher)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func ResponseJSON(w http.ResponseWriter, httpStatusCode int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err.Error())
	}
}

// Fetch query params by key
func (h *Handlers) queryParam(r *http.Request, k string) string {
	v := r.URL.Query().Get(k)
	return v
}

// Returns an error if given param is not present
func (h *Handlers) requireQueryParam(r *http.Request, key string) (string, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return "", fmt.Errorf("missing required query parameter: %s", key)
	}
	return value, nil
}

func (h *Handlers) pathParam(r *http.Request, k string) string {
	v := chi.URLParam(r, k)
	return v
}

func (h *Handlers) rawBody(r *http.Request, i any) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.New("Failed to read body")
	}
	defer r.Body.Close()

	if err := json.Unmarshal(b, i); err != nil {
		return errors.New("Invalid JSON")
	}
	return nil
}

func (h *Handlers) formValue(r *http.Request, k string) string {
	return r.FormValue(k)
}

func (h *Handlers) formFile(r *http.Request, k string, dir string, maxSize int64) (string, error) {
	// Parse multipart form
	if maxSize > 0 {
		err := r.ParseMultipartForm(maxSize << 20) // maxSize MB
		if err != nil {
			return "", errors.New("Error parsing form")
		}
	}
	file, header, err := r.FormFile(k)
	if err != nil {
		return "", errors.New("Error retrieving the file")
	}
	defer file.Close()

	// Save file to local disk (example)

	// check if the dir name is valid
	re := regexp.MustCompile(`^\.\/[a-zA-Z0-9_\-\.]+$`)
	if !re.MatchString(dir) {
		return "", errors.New("Invalid file directory path")
	}

	//dstPath format:  "./uploads/" + Filename
	dstPath := dir + "/" + header.Filename
	os.MkdirAll(dir, os.ModePerm)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", errors.New("Error saving file")

	}
	defer dst.Close()
	io.Copy(dst, file)

	return dstPath, nil
}
