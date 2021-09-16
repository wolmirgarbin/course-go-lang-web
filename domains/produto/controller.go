package produto

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func GetViewListProducts(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Index", ListProducts())
}

func GetViewFormNew(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func GetViewEdit(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Edit", GetById(r.URL.Query().Get("id")))
}

func PostInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Save(ProductReceiver(r))
		http.Redirect(w, r, "/", 301)
	}
	http.NotFound(w, r)
}

func GetDelete(w http.ResponseWriter, r *http.Request) {
	Delete(r.URL.Query().Get("id"))
	http.Redirect(w, r, "/", 301)
}

func PostUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Edit(ProductReceiver(r))
		http.Redirect(w, r, "/", 301)
	}
}
