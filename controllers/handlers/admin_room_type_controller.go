package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/psinthorn/go_smallsite/domain/dbrepo"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/helpers"
	"github.com/psinthorn/go_smallsite/internal/render"
)

// GetRoomStatus
func (rp *Repository) GetAllRoomType(w http.ResponseWriter, r *http.Request) {

}

func (rp *Repository) AddNewRoomTypeForm(w http.ResponseWriter, r *http.Request) {
	var emptyRoomStatus dbrepo.RoomStatus
	data := make(map[string]interface{})
	data["room"] = emptyRoomStatus

	render.Template(w, r, "admin-roomtype-add-form.page.html", &templates.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})

}

func (rp *Repository) AddNewRoomType(w http.ResponseWriter, r *http.Request) {
	// Parse Form
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// Receive form value and pass to room status models
	roomType := dbrepo.RoomType{
		Title:       r.FormValue("title"),
		Description: r.Form.Get("description"),
		Status:      r.Form.Get("status"),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	//
	form := forms.New(r.PostForm)
	if !form.Valid() {
		data := make(map[string]interface{})
		data["room"] = roomType
		render.Template(w, r, "admin-roomtype-add-form.page.html", &templates.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	_, err = dbrepo.RoomTypeService.Create(roomType)
	if err != nil {
		fmt.Println(err)
		return
	}

	rp.App.Session.Put(r.Context(), "room_type", roomType)
	rp.App.Session.Put(r.Context(), "success", "New room type is added :)")
	http.Redirect(w, r, "/admin/rooms/roomtype", http.StatusSeeOther)

}
