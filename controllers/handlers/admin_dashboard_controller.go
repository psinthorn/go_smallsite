package controllers

import (
	"net/http"

	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/render"
)

// GetRoomForm form for create new room
func (rp *Repository) AdminDashBoard(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "admin-dashboard-summary.page.html", &templates.TemplateData{})
}
