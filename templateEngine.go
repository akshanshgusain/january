package january

import (
	"errors"
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
)

type TemplateEngine struct {
	TemplateEngine string
	RootPath       string
	Secure         bool
	Port           string
	ServerName     string
	JetViews       *jet.Set
	Session        *scs.SessionManager
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
	Error           string
	Flash           string
}

func (t *TemplateEngine) defaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Secure = t.Secure
	td.ServerName = t.ServerName
	td.Port = t.Port
	td.CSRFToken = nosurf.Token(r)

	if t.Session.Exists(r.Context(), "userID") {
		td.IsAuthenticated = true
	}
	return td
}

func (t *TemplateEngine) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	switch t.TemplateEngine {
	case "go":
		return t.GoPage(w, r, view, data)
	case "jet":
		return t.JetPage(w, r, view, variables, data)
	default:

	}
	return errors.New("no template engine specified")
}

func (t *TemplateEngine) GoPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
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

func (t *TemplateEngine) JetPage(w http.ResponseWriter, r *http.Request, templateName string, variables, data interface{}) error {
	var vars jet.VarMap

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}
	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	td = t.defaultData(td, r)

	tmpl, err := t.JetViews.GetTemplate(fmt.Sprintf("%s.jet", templateName))
	if err != nil {
		log.Println(err)
		return err
	}
	if err = tmpl.Execute(w, vars, td); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
