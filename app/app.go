package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/psinthorn/go_smallsite/pkg/handlers"
	"github.com/psinthorn/go_smallsite/pkg/renders"
)

const portNumber = ":8080"

// Start use to start new server
func StartApp() {

	CreateSession()

	// Create new template
	tmplCache, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	appConfig.TemplateCache = tmplCache
	appConfig.UseCache = false

	newRepo := handlers.NewRepository(&appConfig)
	handlers.NewHandlers(newRepo)

	renders.NewTemplate(&appConfig)

	// Serve server service
	serveErr := http.ListenAndServe(":8080", routes(&appConfig))
	fmt.Println(fmt.Sprintf("Server is started on port %s", portNumber))
	if serveErr != nil {
		fmt.Println("------------------------------------------------------------------")
		fmt.Println(fmt.Sprintf(":( sorry can't start server on port %s ", portNumber))
		fmt.Println(fmt.Sprintf("error: %s", serveErr))
		fmt.Println("------------------------------------------------------------------")
	}

	// // another serve server config
	// srv := &http.Server{
	// 	Addr:    portNumber,
	// 	Handler: routes(&appConfig),
	// }

	// err = srv.ListenAndServe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
