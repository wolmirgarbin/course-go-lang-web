package produto

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func ViewListProducts(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Index", ListProducts())
}

func ViewFormNew(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func ControllerInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Save(ProductReceiver(r))
		http.Redirect(w, r, "/", 301)
	}
	http.NotFound(w, r)
}

func ControllerDelete(w http.ResponseWriter, r *http.Request) {
	Delete(r.URL.Query().Get("id"))
	http.Redirect(w, r, "/", 301)
}
