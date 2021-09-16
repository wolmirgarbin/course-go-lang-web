package routes

import (
	"github.com/wolmirgarbin/go-web-alura/domains/produto"
	"github.com/wolmirgarbin/go-web-alura/domains/user"
)

func LoadRoutes() {
	produto.CreateRoutes()
	user.CreateRoutes()
}
