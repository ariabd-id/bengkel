package main

import (
	"bengkel/internal/component"
	"bengkel/internal/config"
	"bengkel/internal/module/customer"
	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := config.Get()

	dbConnection := component.GetDatabaseConnection(conf)

	customerRepository := customer.NewRepository(dbConnection)
	customerService := customer.NewService(customerRepository)

	app := fiber.New()
	customer.NewAPI(app, customerService)

	_ = app.Listen(conf.Srv.Host + ":" + conf.Srv.Port)

}
