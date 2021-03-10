package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/psinthorn/go_smallsite/pkg/configs"
	"github.com/psinthorn/go_smallsite/pkg/renders"
)

const portNumber = ":8080"

func Start() {
	var app configs.AppConfig

	tmplCache, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tmplCache

	mapUrls()

	fmt.Println(fmt.Sprintf("Server is started on port %s", portNumber))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("can't start server on port %s", portNumber))
	}

}
