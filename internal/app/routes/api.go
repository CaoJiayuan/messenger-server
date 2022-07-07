package routes

import (
	"messenger-server/internal/app/handlers/api"

	"github.com/enorith/http/router"
)

func ApiRoutes(r *router.Wrapper) {
	var apiHandler api.AuthHandler

	r.Post("login", apiHandler.Login)
	r.Post("register", apiHandler.Register)

	r.Group(func(r *router.Wrapper) {
		r.Get("user", apiHandler.User)

		bHandler := api.IoHandler{}
		r.Post("broadcast", bHandler.Broadcast)
	}).Middleware("auth")
}
