package produto

import "net/http"

// Arquivo de rotas
func CreateRoutes() {
	http.HandleFunc("/", ViewListProducts)
	http.HandleFunc("/new", ViewFormNew)
	http.HandleFunc("/insert", ControllerInsert)
	http.HandleFunc("/delete", ControllerDelete)
}
