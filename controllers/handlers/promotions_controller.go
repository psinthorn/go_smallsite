// routes -> controllers -> services -> domain

package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	domain "github.com/psinthorn/go_smallsite/domain/promotions"
	domain_reservation "github.com/psinthorn/go_smallsite/domain/reservations"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/helpers"
	"github.com/psinthorn/go_smallsite/internal/render"
)

// Promotions is page render
func (rp *Repository) Promotions(w http.ResponseWriter, r *http.Request) {
	st := r.URL.Query().Get("status")
	if st == "" {
		st = "enable"
	}

	promotions, err := domain.PromotionService.Get(st)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	data := make(map[string]interface{})
	data["promotions"] = promotions
	render.Template(w, r, "promotions.page.html", &templates.TemplateData{
		Data: data,
	})
}

// // Promotions is page render
// func (rp *Repository) PromotionTypes(w http.ResponseWriter, r *http.Request) {
// 	pts, err := domain.PromotionTypeService.Get()
// 	if err != nil {
// 		helpers.ServerError(w, err)
// 		return
// 	}

// 	data := make(map[string]interface{})
// 	data["promotions_types"] = pts
// 	render.Template(w, r, "promotions.page.html", &templates.TemplateData{
// 		Data: data,
// 	})
// }

// SearchPromotionAvailability is search promotion availability page render
func (rp *Repository) PromotionRoomType(w http.ResponseWriter, r *http.Request) {
	proId, err := strconv.Atoi(chi.URLParam(r, "id"))
	proTypeId, err := strconv.Atoi(chi.URLParam(r, "type"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	pm, err := domain.PromotionService.GetByID(proId)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	data := make(map[string]interface{})
	data["promotion"] = pm

	stringMap := make(map[string]string)
	rsvn := domain_reservation.Reservation{
		IsPromotion:     true,
		PromotionId:     proId,
		PromotionTypeId: proTypeId,
		// StartDate: ,
		// EndDate:   ,
	}
	rsvn.PromotionId = proId
	//rsvn.Room.Pro = proId
	rp.App.Session.Put(r.Context(), "promotion", pm)
	rp.App.Session.Put(r.Context(), "reservation", rsvn)
	render.Template(w, r, "promotions-roomtype.page.html", &templates.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})

}
