package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	domain "github.com/psinthorn/go_smallsite/domain/promotions"
	domain_promotions "github.com/psinthorn/go_smallsite/domain/promotions"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/helpers"
	"github.com/psinthorn/go_smallsite/internal/render"
	"github.com/psinthorn/go_smallsite/internal/utils"
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

	fmt.Println(promotions)

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
	st := r.URL.Query().Get("status")
	if st == "" {
		st = "enable"
	}
	promotionTypes, err := domain.PromotionTypeService.Get(st)
	if err != nil {
		helpers.ServerError(w, err)
	}

	fmt.Println(promotionTypes)

	var emptyPromotion domain.Promotion
	data := make(map[string]interface{})
	data["promotion"] = emptyPromotion
	data["promotion_types"] = promotionTypes

	render.Template(w, r, "admin-promotion-add-form.page.html", &templates.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// ReservationLists Show all current reservations
func (rp *Repository) AddPromotion(w http.ResponseWriter, r *http.Request) {
	form := forms.New(r.PostForm)
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")

	title := r.Form.Get("title")
	description := r.Form.Get("description")

	fmt.Println(title)
	fmt.Println(description)

	pmtTitle := r.Form.Get("promotion_type_title")
	fmt.Println(pmtTitle)
	pmtTypeId := r.Form.Get("promotion_type_id")
	fmt.Println(pmtTypeId)
	// pmtTypeIdInt, err := strconv.Atoi(pmtTypeId)
	// if err != nil {
	// 	helpers.ServerError(w, err)
	// 	return
	// }

	price := r.Form.Get("price")
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	status := r.Form.Get("status")

	// form.Has("first_name", r)
	form.Required("title", "description", "promotion_type_id", "start_date", "end_date")
	// minimum require on input field
	form.MinLength("title", 12, r)

	// convert from string date to time.Time format
	startDate, err := utils.UtilsService.StringToTime(sd)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	endDate, err := utils.UtilsService.StringToTime(ed)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	pm := domain_promotions.Promotion{
		Title:           title,
		Description:     description,
		Price:           priceInt,
		Status:          status,
		PromotionTypeId: 1,
		StartDate:       startDate,
		EndDate:         endDate,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	fmt.Println("data from pm add form")
	fmt.Println(pm)

	// if data not valid then return current data to form
	if !form.Valid() {
		stringMap := make(map[string]string)
		stringMap["start_date"] = sd
		stringMap["end_date"] = ed
		stringMap["title"] = title
		stringMap["description"] = description
		stringMap["price"] = price

		data := make(map[string]interface{})
		data["promotion"] = pm
		render.Template(w, r, "admin-promotion-add-form.page.html", &templates.TemplateData{
			Form:      form,
			Data:      data,
			StringMap: stringMap,
		})
		return
	}

	rp.App.Session.Put(r.Context(), "promotion", pm)
	rp.App.Session.Put(r.Context(), "success", "New promotion added :)")
	http.Redirect(w, r, "/admin/promotions", http.StatusSeeOther)

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
