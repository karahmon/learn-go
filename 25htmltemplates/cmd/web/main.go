package main

import (
	"fmt"
	"log"
	"main/25htmltemplates/pkg/config"
	"main/25htmltemplates/pkg/handlers"
	"main/25htmltemplates/pkg/render"
	"net/http"
)

const portNumber = ":3000"

func main() {
	var app config.Appconfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Cannot Create Template Cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting Application on port", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
