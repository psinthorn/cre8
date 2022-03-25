package cre8

import (
	"log"

	"github.com/go-chi/chi/v5"
)

type Cre8 struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
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
