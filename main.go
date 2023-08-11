package main

import (
	db "golang-exercise/config"
	routes "golang-exercise/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// /// method 1
	// // app.Get("/", func(c *fiber.Ctx) error {
	// // 	return c.SendString("Hello, World!")
	// // })

	// /// method 2
	// // app.Get("/", hellox)

	// app.Use(app)

	// to connect database at config
	db.Connect()

	// to initialize routes
	routes.Setup(app)

	app.Listen(":3000")

}

