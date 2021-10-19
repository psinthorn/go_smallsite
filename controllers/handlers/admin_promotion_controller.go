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

// PromotionTypeList
func (rp *Repository) PromotionTypesList(w http.ResponseWriter, r *http.Request) {
	st := r.URL.Query().Get("status")
	if st == "" {
		st = "enable"
	}
	promotion_types, err := domain.PromotionTypeService.Get(st)
	if err != nil {
		helpers.ServerError(w, err)
	}

	data := make(map[string]interface{})
	data["promotion_types"] = promotion_types
	render.Template(w, r, "admin-promotion-types.page.html", &templates.TemplateData{
		Data: data,
	})
}

// PromotionForm form for create new promotion
func (rp *Repository) PromotionForm(w http.ResponseWriter, r *http.Request) {
	var emptyPromotion domain.Promotion
	data := make(map[string]interface{})
	data["promotion"] = emptyPromotion

	render.Template(w, r, "admin-promotion-add-form.page.html", &templates.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// Promotion return a promotion information
func (rp *Repository) Promotion(w http.ResponseWriter, r *http.Request) {
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

	pm, err := domain.PromotionService.GetByID(id)
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
