package handlers

import (
	"net/http"

	"github.com/psinthorn/go_smallsite/pkg/renders"
)

// Home is home page render
func Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.html")
}

// About is about page render
func About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.page.html")

}
