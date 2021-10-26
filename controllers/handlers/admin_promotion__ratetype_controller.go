package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	domain "github.com/psinthorn/go_smallsite/domain/promotions"
	domain_promotions "github.com/psinthorn/go_smallsite/domain/promotions"
	"github.com/psinthorn/go_smallsite/domain/rooms"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/helpers"
	"github.com/psinthorn/go_smallsite/internal/render"
	"github.com/psinthorn/go_smallsite/internal/utils"
)

// AdminPromotionTypes
func (rp *Repository) AdminPromotionRateTypes(w http.ResponseWriter, r *http.Request) {
	var rts []rooms.RoomType
	var pms []domain.Promotion
	//rts, err := rooms.RoomTypeService.GetRoomTypeByID()
	pms, err := domain.PromotionService.AdminGet()
	if err != nil {
		helpers.ServerError(w, err)
	}

	data := make(map[string]interface{})
	data["room_types"] = rts
	data["promotions"] = pms
	//data["promotion_ratetypes"] = promotion_types
	render.Template(w, r, "admin-promotion-ratetypes.page.html", &templates.TemplateData{
		Data: data,
	})
}

// Promotion return a promotion information
func (rp *Repository) PromotionRateType(w http.ResponseWriter, r *http.Request) {
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
	render.Template(w, r, "admin-promotion-ratetype-details.page.html", &templates.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})

}

// AdminPromotionForm form for create new promotion
func (rp *Repository) PromotionRateTypeForm(w http.ResponseWriter, r *http.Request) {
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

	render.Template(w, r, "admin-promotion-ratetype-add-form.page.html", &templates.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// ReservationLists Show all current reservations
func (rp *Repository) AddPromotionRateType(w http.ResponseWriter, r *http.Request) {
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
		render.Template(w, r, "admin-promotion-ratetype-add-form.page.html", &templates.TemplateData{
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
	http.Redirect(w, r, "/admin/promotions-ratetypes", http.StatusSeeOther)

}

// Update promotion
func (rp *Repository) UpdatePromotionRateType(w http.ResponseWriter, r *http.Request) {
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
func (rp *Repository) DeletePromotionRateType(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	fmt.Println(id)

	_ = domain.PromotionTypeService.Delete(id)

	rp.App.Session.Put(r.Context(), "success", "promotion type is deleted")
	http.Redirect(w, r, "/admin/promotions-ratetypes", http.StatusSeeOther)
}
