package routes

import (
	controller "golang-exercise/controller"

	"github.com/gofiber/fiber/v2"
)

// here we will setup our routes
func Setup(app *fiber.App) {

	app.Get("/", hellox)
	app.Get("/list-all-product", controller.GetListAllProduct)
	app.Get("/listproduct/:pid", controller.GetListProductByPid)

}

// Handler
func hellox(c *fiber.Ctx) error {
	return c.SendString("Hello anak anak ikan golangx!")
}
