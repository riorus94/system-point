package contoller

import (
	"system-point/delivery/http/dto/request"
	"system-point/domain"
	"system-point/shared/util"

	"github.com/gofiber/fiber/v2"
)

type RegisterController struct {
	domain domain.Domain
}

func NewRegisterController(domain domain.Domain) RegisterController {
	return RegisterController{
		domain: domain,
	}
}

func (input *RegisterController) CreateUser(ctx *fiber.Ctx) error {
	var reqister request.RequestRegisterDTO
	if err := ctx.BodyParser(&reqister); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := input.domain.UserRegistrationUsecase.Register(reqister.ToRegister()); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to register")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/")
}
