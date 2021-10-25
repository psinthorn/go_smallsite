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

// // PromotionTypeList
// func (rp *Repository) PromotionsTypeList(w http.ResponseWriter, r *http.Request) {
// 	st := r.URL.Query().Get("status")
// 	if st == "" {
// 		st = "enable"
// 	}
// 	promotions, err := domain.PromotionService.Get(st)
// 	if err != nil {
// 		helpers.ServerError(w, err)
// 	}

// 	fmt.Println(promotions)

// 	data := make(map[string]interface{})
// 	data["promotions"] = promotions
// 	render.Template(w, r, "admin-promotions-list.page.html", &templates.TemplateData{
// 		Data: data,
// 	})
// }

// // AdminPromotionList
// func (rp *Repository) AdminPromotionsTypeList(w http.ResponseWriter, r *http.Request) {
// 	st := r.URL.Query().Get("status")
// 	if st == "" {
// 		st = "enable"
// 	}

// 	promotions, err := domain.PromotionService.AdminGet()
// 	if err != nil {
// 		helpers.ServerError(w, err)
// 	}

// 	data := make(map[string]interface{})
// 	data["promotions"] = promotions
// 	render.Template(w, r, "admin-promotions-list.page.html", &templates.TemplateData{
// 		Data: data,
// 	})
// }

// AdminPromotionTypes
func (rp *Repository) AdminPromotionTypes(w http.ResponseWriter, r *http.Request) {
	// st := r.URL.Query().Get("status")
	// if st == "" {
	// 	st = "enable"
	// }
	promotion_types, err := domain.PromotionTypeService.AdminGet()
	if err != nil {
		helpers.ServerError(w, err)
	}

	data := make(map[string]interface{})
	data["promotion_types"] = promotion_types
	render.Template(w, r, "admin-promotion-types.page.html", &templates.TemplateData{
		Data: data,
	})
}

// Promotion return a promotion information
func (rp *Repository) PromotionType(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	fmt.Println(id)

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

	pms, err := domain.PromotionService.Get("enable")
	if err != nil {
		helpers.ServerError(w, err)
	}

	st := r.URL.Query().Get("status")
	if st == "" {
		st = "enable"
	}
	promotionType, err := domain.PromotionTypeService.GetById(id)
	if err != nil {
		helpers.ServerError(w, err)
	}

	data := make(map[string]interface{})
	data["promotions"] = pms
	data["promotion_type"] = promotionType
	render.Template(w, r, "admin-promotion-type-details.page.html", &templates.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})

}

// AdminPromotionForm form for create new promotion
func (rp *Repository) PromotionTypeForm(w http.ResponseWriter, r *http.Request) {
	st := r.URL.Query().Get("status")
	if st == "" {
		st = "enable"
	}
	promotionTypes, err := domain.PromotionTypeService.Get(st)
	if err != nil {
		helpers.ServerError(w, err)
	}

	fmt.Println(promotionTypes)
	data := make(map[string]interface{})
	data["promotion_types"] = promotionTypes

	render.Template(w, r, "admin-promotion-type-add-form.page.html", &templates.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// ReservationLists Show all current reservations
func (rp *Repository) AddPromotionType(w http.ResponseWriter, r *http.Request) {
	form := forms.New(r.PostForm)
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	title := r.Form.Get("title")
	description := r.Form.Get("description")
	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")
	status := r.Form.Get("status")

	// form.Has("first_name", r)
	form.Required("title", "description", "start_date", "end_date", "status")
	// minimum require on input field
	form.MinLength("title", 8, r)

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

	pmt := domain_promotions.PromotionType{
		Title:       title,
		Description: description,
		Status:      status,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// if data not valid then return current data to form
	if !form.Valid() {
		stringMap := make(map[string]string)
		stringMap["start_date"] = sd
		stringMap["end_date"] = ed
		stringMap["title"] = title
		stringMap["description"] = description
		stringMap["status"] = status

		data := make(map[string]interface{})
		data["promotion_type"] = pmt
		render.Template(w, r, "admin-promotion-type-add-form.page.html", &templates.TemplateData{
			Form:      form,
			Data:      data,
			StringMap: stringMap,
		})
		return
	}

	_, err = domain.PromotionTypeService.Create(pmt)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rp.App.Session.Put(r.Context(), "promotion_type", pmt)
	rp.App.Session.Put(r.Context(), "success", "New promotion type added :)")
	http.Redirect(w, r, "/admin/promotions-types", http.StatusSeeOther)

}

// Update promotion
func (rp *Repository) UpdatePromotionType(w http.ResponseWriter, r *http.Request) {
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

	pmt, err := domain.PromotionTypeService.GetById(id)

	pmt.Title = r.Form.Get("title")
	pmt.Description = r.Form.Get("description")
	pmt.StartDate = startDate
	pmt.EndDate = endDate
	pmt.Status = r.Form.Get("status")

	_ = domain.PromotionTypeService.Update(pmt)

	rp.App.Session.Put(r.Context(), "success", "promotion type is updated")
	http.Redirect(w, r, "/admin/promotions-types", http.StatusSeeOther)
}

// Delete
func (rp *Repository) DeletePromotionType(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	fmt.Println(id)

	_ = domain.PromotionTypeService.Delete(id)

	rp.App.Session.Put(r.Context(), "success", "promotion type is deleted")
	http.Redirect(w, r, "/admin/promotions-types", http.StatusSeeOther)
}
