package router

import (
	"system-point/delivery/http/contoller"
	"system-point/domain"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, domain domain.Domain) {
	kindActivityController := contoller.NewKindActivityController(domain)

	app.Get("/", kindActivityController.SumAllPointsHandler)
	app.Post("/activity", kindActivityController.AddActivity)
}
