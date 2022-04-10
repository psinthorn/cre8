package main

import (
	"app/handlers"
	"log"
	"os"

	"github.com/psinthorn/cre8"
)

func initApplication() *application {
	// get cuurent folder path
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init cre8
	cre8 := &cre8.Cre8{
		AppName: "app",
		Debug:   true,
	}

	if err := cre8.New(path); err != nil {
		log.Fatal(err)
	}

	cre8.AppName = "app"
	appHandlers := &handlers.Handlers{
		App: cre8,
	}

	app := &application{
		App:      cre8,
		Handlers: appHandlers,
	}

	// เพื่อให้ cre8 รู้จัก routes ที่เราสร้างขึ้นมา
	app.App.Routes = app.routes()

	return app
}
