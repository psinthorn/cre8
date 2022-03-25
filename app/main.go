package main

import "github.com/psinthorn/cre8"

type application struct {
	App *cre8.Cre8
}

func main() {
	cre8 := initApplication()
	cre8.App.ListenAndServe()
}
