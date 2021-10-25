package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	domain "github.com/psinthorn/go_smallsite/domain/promotions"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/helpers"
	"github.com/psinthorn/go_smallsite/internal/render"
)

// PromotionList
func (rp *Repository) PromotionRoomRateList(w http.ResponseWriter, r *http.Request) {
	st := r.URL.Query().Get("status")
	if st == "" {
		st = "enable"
	}
	promotionsRoomrate, err := domain.PromotionRoomRateService.Get(st)
	if err != nil {
		helpers.ServerError(w, err)
	}

	data := make(map[string]interface{})
	data["promotions"] = promotionsRoomrate
	render.Template(w, r, "admin-promotions-Roomrate-list.page.html", &templates.TemplateData{
		Data: data,
	})
}

// GetRoomForm form for create new room
func (rp *Repository) PromotionRoomRateForm(w http.ResponseWriter, r *http.Request) {
	var emptyPmr domain.PromotionRoomRate
	data := make(map[string]interface{})
	data["room"] = emptyPmr

	render.Template(w, r, "admin-promotion-roomrate-form.page.html", &templates.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PromotionRoomrate return a promotion information
func (rp *Repository) PromotionRoomRate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	stringMap := make(map[string]string)
	readOnly := "readonly"
	var isView, isEdit string = "true", ""
	edit := r.URL.Query().Get("edit")
	if edit == "true" {
		readOnly = ""
		isView = ""
		isEdit = "true"
	}

	stringMap["is_read"] = readOnly
	stringMap["is_edit"] = isEdit
	stringMap["is_view"] = isView

	pm, err := domain.PromotionService.GetById(id)
	if err != nil {
		helpers.ServerError(w, err)
	}

	data := make(map[string]interface{})
	data["promotion"] = pm
	render.Template(w, r, "admin-promotion-details.page.html", &templates.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})

}
