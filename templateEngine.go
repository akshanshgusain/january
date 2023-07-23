package january

import (
	"fmt"
	"html/template"
	"net/http"
)

type TemplateEngine struct {
	TemplateEngine string
	RootPath       string
	Secure         bool
	Port           string
	ServerName     string
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
}

func (t *TemplateEngine) Page(w http.ResponseWriter, r *http.ResponseWriter, view string, variables, data interface{}) error {
	switch t.TemplateEngine {
	case "go":
		return t.GoPage(w, r, view, variables, data)
	case "jet":
	}
	return nil
}

func (t *TemplateEngine) GoPage(w http.ResponseWriter, r *http.ResponseWriter, view string, variables, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", t.RootPath, view))
	if err != nil {
		return err
	}
	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}
	if err = tmpl.Execute(w, &td); err != nil {
		return err
	}
	return nil
}
