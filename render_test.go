package january

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_Page(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	te.TemplateEngine = "go"
	te.RootPath = "./testData"

	err = te.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("error rendering page", err)
	}
	err = te.Page(w, r, "no-file", nil, nil)
	if err != nil {
		t.Error("error rendering non-existent go template", err)
	}

	te.TemplateEngine = "jet"
	err = te.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("error rendering page", err)
	}

	err = te.Page(w, r, "no-file", nil, nil)
	if err != nil {
		t.Error("error rendering non-existent jet template", err)
	}
}
