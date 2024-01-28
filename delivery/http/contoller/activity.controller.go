package contoller

import (
	"system-point/delivery/http/dto/request"
	"system-point/domain"
	"system-point/shared/util"

	"github.com/gofiber/fiber/v2"
)

type KindActivityController struct {
	domain domain.Domain
}

func NewKindActivityController(domain domain.Domain) KindActivityController {
	return KindActivityController{
		domain: domain,
	}
}

func (input *KindActivityController) SumAllPointsHandler(ctx *fiber.Ctx) error {
	totalPoints, err := input.domain.KindActivityUsecase.SumAllPoints()
	if err != nil {
		panic(err)
	}

	response := map[string]int{"totalPoints": totalPoints}
	return ctx.Render("resource/views/home", fiber.Map{
		"Points": response,
	})
}

func (input *KindActivityController) AddActivity(ctx *fiber.Ctx) error {
	var kindActivity request.RequestAddKindActivityDTO
	if err := ctx.BodyParser(&kindActivity); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := input.domain.KindActivityUsecase.SaveActivity(kindActivity.ToAddKindActivity()); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to save activity")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/")

}
