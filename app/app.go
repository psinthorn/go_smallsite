package main

import (
	"fmt"
	"log"
	"os"

	"github.com/psinthorn/go_smallsite/internal/drivers"
	"github.com/psinthorn/go_smallsite/internal/handlers"
	"github.com/psinthorn/go_smallsite/internal/renders"
	"github.com/psinthorn/go_smallsite/internal/utils"
)

var infoLog *log.Logger
var errorLog *log.Logger

// Start use to start new server
func StartApp() (*drivers.DB, error) {

	// Check env is production
	utils.Utils.IsProduction(&appConfig)

	// Write server and client error log to logs file
	infoLog = log.New(os.Stdout, "Info\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	appConfig.InfoLog = infoLog
	appConfig.ErrorLog = errorLog

	// Start session
	CreateSession()

	// Create new template
	tmplCache, err := renders.CreateTemplateCache()
	if err != nil {
		return nil, err
	}

	// Create and load config to templates
	appConfig.TemplateCache = tmplCache
	appConfig.UseCache = false
	newRepo := handlers.NewRepository(&appConfig)
	handlers.NewHandlers(newRepo)
	renders.NewTemplate(&appConfig)

	// Connect to postgress databast
	fmt.Println("Connecting to Database...")
	dsn := "host=localhost port=5432 dbname=go_smallsite_bookings user=postgres password="
	dbConn, err := drivers.ConnectSQL("pgx", dsn)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connecting to Database Success fully :)")

	// return database connect to startApp function
	return dbConn, nil

}
