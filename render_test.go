package january

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pd = []struct {
	name           string
	templateEngine string
	template       string
	errorExpected  bool
	errorMessage   string
}{
	{"go_page", "go", "home", false, "error rendering go template"},
	{"go_page_no_template", "go", "no=file", true, "error rendering non-existent go template, when one is expected"},
	{"jet_page", "jet", "home", false, "error rendering jet template"},
	{"jet_page_no_template", "jet", "no=file", true, "error rendering non-existent jet template, when one is expected"},
	{"invalid_template_engine", "na", "home", true, "no error returned while rendering with invalid template engine"},
}

func TestRender_Page(t *testing.T) {

	for _, c := range pd {
		r, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Error(err)
		}
		w := httptest.NewRecorder()
		te.TemplateEngine = c.templateEngine
		te.RootPath = "./testData"

		err = te.Page(w, r, c.template, nil, nil)
		if c.errorExpected {
			if err == nil {
				t.Errorf("%s: %s", c.name, c.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s: %s", c.name, c.errorMessage, err.Error())
			}
		}
	}
}

func TestRender_GoPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}

	te.TemplateEngine = "go"
	te.RootPath = "./testData"

	err = te.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("error rendering page", err)
	}
}

func TestRender_JetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}

	te.TemplateEngine = "jet"

	err = te.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("error rendering page", err)
	}
}
