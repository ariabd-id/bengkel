package customer

import (
	"bengkel/domain"
	"github.com/gofiber/fiber/v2"
)

type api struct {
	customerService domain.CustomerService
}

func NewAPI(app *fiber.App, customerService domain.CustomerService) {
	api := api{
		customerService: customerService,
	}
	app.Get("/foo/bar", api.FooBar)
}

func (api api) FooBar(ctx *fiber.Ctx) error {
	return ctx.JSON("asdd")
}
