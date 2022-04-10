package cre8

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/psinthorn/cre8/render"
)

type Cre8 struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Render   *render.Render
	Routes   *chi.Mux
	config   config
}

type config struct {
	port     string
	renderer string
}

type initFolders struct {
	rootPath    string
	folderNames []string
}

// type render struct {
// 	Renderer   string
// 	RootPath   string
// 	Port       string
// 	ServerName string
// }
