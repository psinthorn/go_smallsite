package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/psinthorn/go_smallsite/pkg/configs"
	"github.com/psinthorn/go_smallsite/pkg/renders"
)

const portNumber = ":8080"

// Start use to start new server
func Start() {
	fmt.Println(fmt.Sprintf("Server is started on port %s", portNumber))
	var app configs.AppConfig

	// Create new template
	tmplCache, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tmplCache
	renders.NewTemplate(&app)

	// Serve server service
	serveErr := http.ListenAndServe(":8080", routes())
	fmt.Println(fmt.Sprintf("Server is started on port %s", portNumber))
	if serveErr != nil {
		fmt.Println("------------------------------------------------------------------")
		fmt.Println(fmt.Sprintf(":( sorry can't start server on port %s ", portNumber))
		fmt.Println(fmt.Sprintf("error: %s", serveErr))
		fmt.Println("------------------------------------------------------------------")
	}

}
