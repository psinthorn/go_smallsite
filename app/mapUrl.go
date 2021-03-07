package app

import (
	"net/http"

	"github.com/psinthorn/go_smallsite/pkg/handlers"
)

// mapUrls use to map url with controller func
func mapUrls() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
}
