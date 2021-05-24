package router

import (
	"go-fiber/src/controller"

	"github.com/gofiber/fiber/v2"
)

func CustomerRouter(app *fiber.App) {
	app.Get("api/v1/customers", controller.GetCustomers)
	app.Post("api/v1/customers", controller.StoreCustomer)
	app.Get("api/v1/customers/:id", controller.GetCustomer)
	app.Put("api/v1/customers/:id", controller.UpdateCustomer)
	app.Delete("api/v1/customers/:id", controller.RemoveCustomer)
}
