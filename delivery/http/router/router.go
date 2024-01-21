package router

import (
	"system-point/delivery/http/contoller"
	"system-point/domain"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, domain domain.Domain) {
	registrationController := contoller.NewRegisterController(domain)

	app.Post("/", registrationController.CreateUser)
}
