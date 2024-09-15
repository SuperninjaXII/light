package routes

import (
	"Edge/controllers"

	"github.com/gofiber/fiber/v2"
)

func Apt(app *fiber.App) {
	app.Get("/list", controllers.ListPackages)
	app.Get("/install", controllers.InstallPackage)
	app.Get("/search", controllers.SearchPackage)
	app.Get("/show", controllers.ShowPackage)
}
