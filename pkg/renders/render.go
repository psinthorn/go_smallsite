package renders

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate ช่วยในการ render html template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("error parsing template %s ", err))
		return
	}
}

func createTemplateCache() {

	tmplCache := map[string]*template.Template{}
	fmt.Println(tmplCache)

}
