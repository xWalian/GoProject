package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetOrdersList(app *fiber.App) fiber.Router {
	return app.Get("/orders-list", func(ctx *fiber.Ctx) error {

	})
}
