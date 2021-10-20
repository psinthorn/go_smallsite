package controllers

import (
	"net/http"

	domain "github.com/psinthorn/go_smallsite/domain/contents"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/helpers"
	"github.com/psinthorn/go_smallsite/internal/render"
)

func (rp *Repository) ContentForm(w http.ResponseWriter, r *http.Request) {
	var emptyContent domain.Content
	data := make(map[string]interface{})
	data["room"] = emptyContent

	render.Template(w, r, "admin-content-add-form.page.html", &templates.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (rp *Repository) PostContent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

}

func (rp *Repository) ContentLists(w http.ResponseWriter, r *http.Request) {
	contents, err := domain.ContentService.GetAll()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	data := make(map[string]interface{})
	data["contect"] = contents
	render.Template(w, r, "admin-content-list.page.html", &templates.TemplateData{
		Data: data,
	})
}

func (rp *Repository) ShowContent(w http.ResponseWriter, r *http.Request)     {}
func (rp *Repository) EditContentForm(w http.ResponseWriter, r *http.Request) {}
func (rp *Repository) EditContent(w http.ResponseWriter, r *http.Request)     {}
func (rp *Repository) DeleteContent(w http.ResponseWriter, r *http.Request)   {}
