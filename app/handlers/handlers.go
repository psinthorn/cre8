package handlers

import (
	"net/http"

	"github.com/psinthorn/cre8"
)

type Handlers struct {
	App *cre8.Cre8
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	if err := h.App.Render.PageTmpl(w, r, "home", nil, nil); err != nil {
		h.App.ErrorLog.Println("Error on rendering: ", err)
	}
}
