package router

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("Nothing to see here") })

	apiRoutes := app.Group("/api")
	v1APIRoutes := apiRoutes.Group("/v1")
	{
		// IMPLEMENT ROUTE
		v1APIRoutes.Get("/health", func(c *fiber.Ctx) error { return c.SendString("OK") })
	}
}
