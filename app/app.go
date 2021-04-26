package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/psinthorn/go_smallsite/internal/handlers"
	"github.com/psinthorn/go_smallsite/internal/renders"
	"github.com/psinthorn/go_smallsite/internal/utils"
)

const portNumber = ":8080"

var infoLog *log.Logger
var errorLog *log.Logger

// Start use to start new server
func StartApp() {

	utils.Utils.IsProduction(&appConfig)

	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("Current environment isProduction: ", appConfig.IsProduction)

	infoLog = log.New(os.Stdout, "Info\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	appConfig.InfoLog = infoLog
	appConfig.ErrorLog = errorLog

	CreateSession()

	// Connect to databast
	// log.Println("Connecting to Database......")
	// // db, err := drivers.ConnectSQL("host=localhost port=5432 dbname=go_smallsite_bookngs user=postgres password=")
	// if err != nil {
	// 	log.Fatal("Can't connect to database :(")
	// }

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
