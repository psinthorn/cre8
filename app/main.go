package main

import (
	"app/handlers"

	"github.com/psinthorn/cre8"
)

type application struct {
	App      *cre8.Cre8
	Handlers *handlers.Handlers
}

func main() {
	cre8 := initApplication()
	cre8.App.ListenAndServe()
}
