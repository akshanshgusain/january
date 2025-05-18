package handlers

import (
	"github.com/akshanshgusain/january"
	"januaryApp/data"
	"net/http"
)

type Handlers struct {
	App    *january.January
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	if err := h.render(w, r, "home", nil, nil); err != nil {
		h.App.ErrorLog.Println("error rendering: ", err)
	}
}
