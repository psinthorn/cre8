package main

import (
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
	cre8.InfoLog.Println("Debug is set to: ", cre8.Debug)

	app := &application{
		App: cre8,
	}

	return app
}
