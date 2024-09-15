package main

import (
	"Edge/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "home",
		})
	})
	app.Get("/download", func(c *fiber.Ctx) error {
		return c.Render("download", fiber.Map{
			"Title": "download",
		})
	})
	routes.Apt(app)

	log.Fatal(app.Listen(":3000"))
}
