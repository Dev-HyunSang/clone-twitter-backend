package main

import (
	"context"
	"log"

	"github.com/dev-hyunsang/clone-twitter-backend/database"
	"github.com/dev-hyunsang/clone-twitter-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.Middleware(app)

	client, err := database.ConnectMySQL()
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatalln(err)
	}

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
