package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	formField := r.Form.Get(field)
	if formField == "" {
		f.Errors.Add(field, "this field can not be blank")
		return false
	}
	return true
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
