package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	domain "github.com/psinthorn/go_smallsite/domain/promotions"
	"github.com/psinthorn/go_smallsite/domain/rates"
	"github.com/psinthorn/go_smallsite/domain/rooms"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/helpers"
	"github.com/psinthorn/go_smallsite/internal/render"
	"github.com/psinthorn/go_smallsite/internal/utils"
)

// AdminPromotionTypes
func (rp *Repository) AdminPromotionRate(w http.ResponseWriter, r *http.Request) {
	var rts []rates.RateType
	rts, err := rates.RateTypeService.AdminGet()
	if err != nil {
		helpers.ServerError(w, err)
	}

	data := make(map[string]interface{})
	data["rate_types"] = rts
	render.Template(w, r, "admin-rate-types.page.html", &templates.TemplateData{
		Data: data,
	})
}

// Promotion return a ratetype information
func (rp *Repository) PromotionRate(w http.ResponseWriter, r *http.Request) {
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

	st := r.URL.Query().Get("status")
	if st == "" {
		st = "enable"
	}
	rateType, err := rates.RateTypeService.GetById(id)
	if err != nil {
		helpers.ServerError(w, err)
	}

	data := make(map[string]interface{})
	data["rate_type"] = rateType
	render.Template(w, r, "admin--rate-type-details.page.html", &templates.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})

}

// AdminPromotionForm form for create new promotion
func (rp *Repository) PromotionRateForm(w http.ResponseWriter, r *http.Request) {
	st := r.URL.Query().Get("status")
	if st == "" {
		st = "enable"
	}

	// get promotion
	promotions, err := domain.PromotionService.Get(st)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// get room type
	roomTypes, err := rooms.RoomTypeService.Get(st)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// get room type
	rateTypes, err := rates.RateTypeService.Get(st)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// get rate type

	data := make(map[string]interface{})
	data["room_types"] = roomTypes
	data["promotions"] = promotions
	data["rate_types"] = rateTypes
	render.Template(w, r, "admin-promotion-ratetype-add-form.page.html", &templates.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// ReservationLists Show all current reservations
func (rp *Repository) AddPromotionRate(w http.ResponseWriter, r *http.Request) {
	form := forms.New(r.PostForm)
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	title := r.Form.Get("title")
	acronym := r.Form.Get("acronym")
	description := r.Form.Get("description")
	status := r.Form.Get("status")

	// form.Has("first_name", r)
	form.Required("title", "acronym", "description", "status")
	// minimum require on input field
	form.MinLength("title", 8, r)

	rt := rates.RateType{
		Title:       title,
		Acronym:     acronym,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// if data not valid then return current data to form
	if !form.Valid() {
		stringMap := make(map[string]string)
		stringMap["title"] = title
		stringMap["acronym"] = acronym
		stringMap["description"] = description
		stringMap["status"] = status

		data := make(map[string]interface{})
		data["rate_type"] = rt
		render.Template(w, r, "admin-rate-type-add-form.page.html", &templates.TemplateData{
			Form:      form,
			Data:      data,
			StringMap: stringMap,
		})
		return
	}

	_, err = rates.RateTypeService.Create(rt)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rp.App.Session.Put(r.Context(), "rate_type", rt)
	rp.App.Session.Put(r.Context(), "success", "New rate type added :)")
	http.Redirect(w, r, "/admin/rate-types", http.StatusSeeOther)

}

// Update promotion
func (rp *Repository) UpdatePromotionRate(w http.ResponseWriter, r *http.Request) {
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
	http.Redirect(w, r, "/admin/promotions-ratetypes", http.StatusSeeOther)
}

// Delete
func (rp *Repository) DeletePromotionRate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	fmt.Println(id)

	err = rates.RateTypeService.Delete(id)
	if err != nil {
		helpers.ServerError(w, err)
	}

	rp.App.Session.Put(r.Context(), "success", "promotion type is deleted")
	http.Redirect(w, r, "/admin/promotions-ratetypes", http.StatusSeeOther)
}
