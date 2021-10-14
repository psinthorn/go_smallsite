package controllers

import (
	"net/http"

	domain "github.com/psinthorn/go_smallsite/domain/promotions"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/helpers"
	"github.com/psinthorn/go_smallsite/internal/render"
)

// PromotionList
func (rp *Repository) PromotionsList(w http.ResponseWriter, r *http.Request) {
	st := r.URL.Query().Get("status")
	if st == "" {
		st = "enable"
	}
	promotions, err := domain.PromotionService.Get(st)
	if err != nil {
		helpers.ServerError(w, err)
	}

	data := make(map[string]interface{})
	data["promotions"] = promotions
	render.Template(w, r, "admin-promotions-list.page.html", &templates.TemplateData{
		Data: data,
	})
}

// GetRoomForm form for create new room
func (rp *Repository) AddNewPromotionForm(w http.ResponseWriter, r *http.Request) {
	var emptyPromotion domain.Promotion
	data := make(map[string]interface{})
	data["room"] = emptyPromotion

	render.Template(w, r, "admin-promotion-add-form.page.html", &templates.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// 	rp.App.Session.Remove(r.Context(), "reservation")

// 	data := make(map[string]interface{})
// 	data["reservation"] = reservation
// 	render.Template(w, r, "reservation-summary.page.html", &templates.TemplateData{
// 		Data: data,
// 	})
// }

// // RoomsAll show all rooms
// func (rp *Repository) RoomsAll() (domain.Room, error) {
// 	var roomsAll = domain.Room{}
// 	return roomsAll, nil
// }
