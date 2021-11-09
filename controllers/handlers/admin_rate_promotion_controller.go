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
)

// AdminPromotionRate
func (rp *Repository) AdminPromotionRates(w http.ResponseWriter, r *http.Request) {
	var pmrs []rates.PromotionRate
	pmrs, err := rates.PromotionRateService.AdminGet()
	if err != nil {
		helpers.ServerError(w, err)
	}

	data := make(map[string]interface{})
	data["promotions_rates"] = pmrs
	render.Template(w, r, "admin-promotion-rates.page.html", &templates.TemplateData{
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
	render.Template(w, r, "admin-rate-type-details.page.html", &templates.TemplateData{
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
	pmId, _ := strconv.Atoi(r.Form.Get("promotion_id"))
	roomTypeId, _ := strconv.Atoi(r.Form.Get("room_type_id"))
	rateTypeId, _ := strconv.Atoi(r.Form.Get("rate_type_id"))
	rate, _ := strconv.Atoi(r.Form.Get("rate"))
	status := r.Form.Get("status")

	fmt.Print(r.Form.Get("promotion_id"))
	fmt.Print(r.Form.Get("room_type_id"))
	fmt.Print(r.Form.Get("rate_type_id"))
	fmt.Print(r.Form.Get("rate"))

	// form.Has("first_name", r)
	form.Required("title", "status")
	// minimum require on input field
	form.MinLength("title", 8, r)

	pr := rates.PromotionRate{
		Title:       title,
		PromotionId: pmId,
		RoomTypeId:  roomTypeId,
		Rate:        float32(rate),
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// if data not valid then return current data to form
	if !form.Valid() {
		stringMap := make(map[string]string)
		stringMap["title"] = title
		stringMap["promotion_id"] = string(pmId)
		stringMap["room_type_id"] = string(roomTypeId)
		stringMap["rate_type_id"] = string(rateTypeId)
		stringMap["rate"] = string(rate)
		stringMap["status"] = status

		data := make(map[string]interface{})
		data["rate_type"] = pr
		render.Template(w, r, "admin-promotion-ratetype-add-form.page.html", &templates.TemplateData{
			Form:      form,
			Data:      data,
			StringMap: stringMap,
		})
		return
	}

	id, err := rates.PromotionRateService.Create(pr)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	fmt.Println(id)

	rp.App.Session.Put(r.Context(), "rate_type", pr)
	rp.App.Session.Put(r.Context(), "success", "New rate type added :)")
	http.Redirect(w, r, "/admin/promotion-rates/new", http.StatusSeeOther)

}

// UpdatePromotionRate update promotion rate for each room type and status
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
	rate, err := strconv.Atoi(r.Form.Get("rate"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	status := r.Form.Get("status")

	var pmr rates.PromotionRate
	pmr, err = rates.PromotionRateService.GetById(id)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	pmr.Rate = float32(rate)
	pmr.Status = status

	_ = rates.PromotionRateService.Update(pmr)

	rp.App.Session.Put(r.Context(), "success", "promotion rate is updated")
	http.Redirect(w, r, fmt.Sprintf("/admin/promotions/26?edit=false"), http.StatusSeeOther)
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
