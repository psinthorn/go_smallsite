package renders

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	functions = template.FuncMap{}
	tmplCache = map[string]*template.Template{}
)

// RenderTemplate ช่วยในการ render html template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// Use single template
	tmplCache, err := createSingleTemplateCache(tmpl)

	// tmplCache, err := createTemplateCache(tmpl)
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

func createTemplateCache(tmpl string) (map[string]*template.Template, error) {

	//tmplCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		fmt.Println(fmt.Sprintf("error on check file path %s", err))
		return nil, err
	}

	fmt.Println("agrs name is: ", tmpl)
	fmt.Println("---------------------")

	for _, page := range pages {

		pageName := filepath.Base(page)
		fmt.Println("page name is: ", pageName)
		fmt.Println("---------------------")

		// ตรวจสอบว่ามี page ที่ชื่อตรงกับ args หรือไม่ หากทีให้ทำการสร้าง template ขึ้นมาตามชื่อที่ส่งเข้ามา (args)
		if pageName == tmpl {
			tmplSet, err := template.New(pageName).Funcs(functions).ParseFiles(page)
			if err != nil {
				fmt.Println(fmt.Sprintf("error can't create new template set error: %s ", err))
				return nil, err
			}

			// ตรวจสอบว่ามีไฟล์ที่ลงท้ายด้วยนามสกุล .layout.html หรือไม่หาก
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

	}

	return tmplCache, nil
}

// createSingleTemplateCache use to create and parse layout single template cache
func createSingleTemplateCache(tmpl string) (map[string]*template.Template, error) {

	page, err := filepath.Glob("./templates/" + tmpl)
	if err != nil {
		return nil, err
	}

	if len(page) > 0 {
		fmt.Println("tmpl name: ", tmpl)
		fmt.Println("page name: ", tmpl)
		fmt.Println("------------------------")

		pageURL := "./templates/" + tmpl
		tmplSet, err := template.New(tmpl).Funcs(functions).ParseFiles(pageURL)
		if err != nil {
			return nil, err
		}

		// ตรวจสอบว่ามีไฟล์ที่ลงท้ายด้วยนามสกุล .layout.html หรือไม่หาก
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

			tmplCache[tmpl] = tmplSet

		}

	} // endif

	return tmplCache, nil
}
