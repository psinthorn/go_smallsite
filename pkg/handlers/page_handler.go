package handlers

import (
	"net/http"

	"github.com/psinthorn/go_smallsite/pkg/configs"
	"github.com/psinthorn/go_smallsite/pkg/renders"
)

// Repo
var Repo *Repository

type Repository struct {
	App *configs.AppConfig
}

// NewRepository
func NewRepository(a *configs.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is home page render
func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("rq Home page")
	renders.RenderTemplate(w, "home.page.html")
}

// About is about page render
func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("rq About page")
	renders.RenderTemplate(w, "about.page.html")

}
