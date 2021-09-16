package main

import (
	"net/http"

	"github.com/wolmirgarbin/go-web-alura/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
