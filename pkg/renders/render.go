package renders

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate ช่วยในการ render html template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tmplCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	newTmpl, ok := tmplCache[tmpl]
	if !ok {
		log.Fatal("error on parsing template")
	}

	buff := new(bytes.Buffer)
	_ = newTmpl.Execute(buff, nil)
	_, err = buff.WriteTo(w)
	if err != nil {
		fmt.Println(fmt.Sprintf("error parsing template to browser %s ", err))
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	tmplCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		fmt.Println(fmt.Sprintf("error on check file path %s", err))
		return nil, err
	}

	for _, page := range pages {
		pageName := filepath.Base(page)
		// fmt.Println(page)

		tmplSet, err := template.New(pageName).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println(fmt.Sprintf("error can't create new template set error: %s ", err))
			return nil, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			fmt.Println(fmt.Sprintf("error can't find any layout file in templates folder error: %s ", err))
			return nil, err
		}

		if len(matches) > 0 {
			tmplSet, err = tmplSet.ParseGlob("./templates/*.layout.html")
			if err != nil {
				fmt.Println(fmt.Sprintf("error can't parse layout file to templates set error: %s ", err))
				return nil, err
			}
		}

		tmplCache[pageName] = tmplSet

	}

	return tmplCache, nil
}
