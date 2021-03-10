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

	// ตรวจสอบว่ามี template ตรงที่ต้องการหรือไหม
	newTmpl, ok := tmplCache[tmpl]
	if !ok {
		log.Fatal("error on parsing template")
	}

	// หากทีให้ทำหารเขียนลงไปที่ bytes buffer
	buff := new(bytes.Buffer)

	// และให้เขียนไปที่ template ใหม่ที่สร้างรอไว้แล้ว (newTmpl)
	_ = newTmpl.Execute(buff, nil)

	// และเขียนส่ง buffer new template ให้ response (w)
	_, err = buff.WriteTo(w)
	if err != nil {
		fmt.Println(fmt.Sprintf("error parsing template to browser %s ", err))
	}
}

// createTemplateCache ตรวจสอบและสร้าง templateห แบบทั้งหมด
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

// createSingleTemplateCache ตรวจสอบและสร้าง template ตาม args ที่ส่งเข้ามาเท่านั้น
func createSingleTemplateCache(tmpl string) (map[string]*template.Template, error) {

	// หาไฟล์ชื่อที่ตรงกับ args ที่ส่งเข้ามา หากไม่ไมีมห้คืนค่า err กลับไป
	page, err := filepath.Glob("./templates/" + tmpl)
	if err != nil {
		return nil, err
	}

	// หากมีไฟล์ชื่อตรงกันให้ทำการเตรียมสร้าง template ขึ้นมาใหม่โดยใช้ชื่อตาม args ที่ส่งเข้ามา (tmpl)
	if len(page) > 0 {
		fmt.Println("tmpl name: ", tmpl)
		fmt.Println("page name: ", tmpl)
		fmt.Println("------------------------")

		pageURL := "./templates/" + tmpl

		// สร้าง template ขึ้นมาใหม่โดยใช้ชื่อตาม args ที่ส่งเข้ามา (tmpl) และ 	URL ของตำแหน่งไฟล์ในโปรเจคให้ถูกต้อง
		newTmpl, err := template.New(tmpl).Funcs(functions).ParseFiles(pageURL)
		if err != nil {
			return nil, err
		}

		// ตรวจสอบว่ามีไฟล์ที่ลงท้ายด้วยนามสกุล .layout.html หรือไม่หาก
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			fmt.Println(fmt.Sprintf("error can't find any layout file in templates folder error: %s ", err))
			return nil, err
		}

		// หากพบ layout.html มากกว่า 0 ให้ทำการเขียนข้อมูลที่มีในไฟล์ไปที่ newTmpl
		if len(matches) > 0 {
			newTmpl, err = newTmpl.ParseGlob("./templates/*.layout.html")
			if err != nil {
				fmt.Println(fmt.Sprintf("error can't parse layout file to templates set error: %s ", err))
				return nil, err
			}

			// เก็บ newTmpl ใน templCache[] อาเรย์
			tmplCache[tmpl] = newTmpl

		}

	} // endif

	// คืนค่า tmplCahe ให้ฟังซ์ชั่น
	return tmplCache, nil
}
