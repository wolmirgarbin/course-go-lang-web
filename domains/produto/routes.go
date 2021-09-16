package produto

import "net/http"

// Arquivo de rotas
func CreateRoutes() {
	http.HandleFunc("/", GetViewListProducts)
	http.HandleFunc("/new", GetViewFormNew)
	http.HandleFunc("/edit", GetViewEdit)
	http.HandleFunc("/insert", PostInsert)
	http.HandleFunc("/delete", GetDelete)
	http.HandleFunc("/update", PostUpdate)
}
