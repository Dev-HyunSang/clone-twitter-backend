package middleware

import (
	"github.com/dev-hyunsang/clone-twitter-backend/cmd"
	"github.com/gofiber/fiber/v2"
)

func Middleware(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/join", cmd.JoinUserHandler)
	auth.Post("/login", cmd.LoginUserHandler)
	auth.Post("/logout", cmd.LogoutUserHandler)

}
