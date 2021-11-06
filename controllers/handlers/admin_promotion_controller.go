package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	domain "github.com/psinthorn/go_smallsite/domain/promotions"
	domain_promotions "github.com/psinthorn/go_smallsite/domain/promotions"
	"github.com/psinthorn/go_smallsite/domain/rates"
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

// PromotionList
func (rp *Repository) AdminPromotionsList(w http.ResponseWriter, r *http.Request) {
	st := r.URL.Query().Get("status")
	if st == "" {
		st = "enable"
	}
	promotions, err := domain.PromotionService.AdminGet()
	if err != nil {
		helpers.ServerError(w, err)
	}

	//fmt.Println(promotions)

	data := make(map[string]interface{})
	data["promotions"] = promotions
	render.Template(w, r, "admin-promotions-list.page.html", &templates.TemplateData{
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

	//fmt.Println(promotionTypes)

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
	// pmtTitle := r.Form.Get("promotion_type_title")
	pmtTypeId := r.Form.Get("promotion_type_id")
	pmtTypeIdInt, err := strconv.Atoi(pmtTypeId)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	price := r.Form.Get("price")
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	status := r.Form.Get("status")

	// form.Has("first_name", r)
	form.Required("title", "description", "promotion_type_id", "start_date", "end_date", "price", "status")
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
		PromotionTypeId: pmtTypeIdInt,
		StartDate:       startDate,
		EndDate:         endDate,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// if data not valid then return current data to form
	if !form.Valid() {
		stringMap := make(map[string]string)
		stringMap["start_date"] = sd
		stringMap["end_date"] = ed
		stringMap["title"] = title
		stringMap["description"] = description
		stringMap["price"] = price
		stringMap["promotion_type_id"] = pmtTypeId
		stringMap["status"] = status

		data := make(map[string]interface{})
		data["promotion"] = pm
		render.Template(w, r, "admin-promotion-add-form.page.html", &templates.TemplateData{
			Form:      form,
			Data:      data,
			StringMap: stringMap,
		})
		return
	}

	_, err = domain.PromotionService.Create(pm)
	if err != nil {
		helpers.ServerError(w, err)
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

	var pmrs []rates.PromotionRate
	pmrs, err = rates.PromotionRateService.GetRatesByPromotionId(id)
	if err != nil {
		helpers.ServerError(w, err)
	}

	st := r.URL.Query().Get("status")
	if st == "" {
		st = "enable"
	}
	promotionTypes, err := domain.PromotionTypeService.Get(st)
	if err != nil {
		helpers.ServerError(w, err)
	}

	data := make(map[string]interface{})
	data["promotions_rates"] = pmrs
	data["promotion"] = pm
	data["promotion_types"] = promotionTypes
	render.Template(w, r, "admin-promotion-details.page.html", &templates.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})

}

// Update promotion
func (rp *Repository) UpdatePromotion(w http.ResponseWriter, r *http.Request) {
	//form := forms.New(r.PostForm)
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	price, err := strconv.Atoi(r.Form.Get("price"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	pmtId, err := strconv.Atoi(r.Form.Get("promotion_type_id"))
	fmt.Println(pmtId)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")
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

	pm, err := domain.PromotionService.GetById(id)

	pm.Title = r.Form.Get("title")
	pm.Description = r.Form.Get("description")
	pm.Price = price
	pm.PromotionTypeId = pmtId
	pm.StartDate = startDate
	pm.EndDate = endDate
	pm.Status = r.Form.Get("status")

	_ = domain.PromotionService.Update(pm)

	rp.App.Session.Put(r.Context(), "success", "promotion package is updated")
	http.Redirect(w, r, "/admin/promotions", http.StatusSeeOther)
}

// Delete
func (rp *Repository) DeletePromotion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	fmt.Println(id)

	_ = domain.PromotionService.Delete(id)

	rp.App.Session.Put(r.Context(), "success", "promotion package is deleted")
	http.Redirect(w, r, "/admin/promotions", http.StatusSeeOther)
}
