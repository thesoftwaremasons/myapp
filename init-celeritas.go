package main

import (
	"log"
	"myapp/handlers"
	"os"

	"github.com/thesoftwaremasons/celeritas"
)

func initApplication() *application {

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	//init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}
	cel.AppName = "myapp"
	myHandlers := &handlers.Handlers{
		App: cel,
	}
	//fmt.Println("Debug is set to", cel.Debug)

	app := &application{
		App:      cel,
		Handlers: myHandlers,
	}
	app.App.Routes = app.routes()
	return app

}
