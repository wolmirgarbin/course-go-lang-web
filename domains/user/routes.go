package user

import "net/http"

const baseContext = "/user"

func CreateRoutes() {
	http.HandleFunc(baseContext, ListAllUsers)
}
