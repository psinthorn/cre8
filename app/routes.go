package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// middleware ต้องมาก่อน route อื่นๆ เสมอ

	// route ที่มีให้บริการเรียกจากตรงนี้
	a.App.Routes.Get("/", a.Handlers.Home)
	a.App.Routes.Get("/jet", func(w http.ResponseWriter, r *http.Request) {
		a.App.Render.JetTmpl(w, r, "test", nil, nil)
	})

	// static route สำหรับกำหนด folder ที่ให้บริการ files, รูปภาพต่างๆ
	// กำหนดให้ใช้ folder ชื่อ public เป็น file server (static folder) โดยใช้ http.Dir()
	filesServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", filesServer))

	return a.App.Routes
}
