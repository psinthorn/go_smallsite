package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

// Home is home page render
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About is about page render
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")

}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("error parsing template %s ", err))
		return
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Server is started on port %s", portNumber))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("can't start server on port %s", portNumber))
	}
}
