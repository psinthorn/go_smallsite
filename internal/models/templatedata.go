package models

// TemplateData struct is holds all data type that we use to send to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	flash     string
	Warnings  string
	Error     string
}
