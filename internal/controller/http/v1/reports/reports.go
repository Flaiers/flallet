package reports

import (
	"github.com/flaiers/flallet/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func New(config Config) func(router fiber.Router) {
	controller := &Controller{service: config.Service}

	return func(router fiber.Router) {
		router.Post("/csv", controller.getCSVReport)
	}
}

// @summary  Get CSV Report
// @tags     reports
// @id       getCSVReport
// @accept   json
// @produce  json
// @param    reportFilter body     dto.ReportFilter true " "
// @response 200          {object} dto.ReportRead
// @response 404          {object} utils.JSONError
// @response 422          {object} utils.JSONError
// @response 500          {object} utils.JSONError
// @router   /reports/csv [post]
func (c *Controller) getCSVReport(ctx *fiber.Ctx) error {
	reportFilter := &dto.ReportFilter{}
	if err := ctx.BodyParser(reportFilter); err != nil {
		return err
	}
	url, err := c.service.GetCSV(reportFilter)
	if err != nil {
		return err
	}

	return ctx.JSON(dto.ReportRead{URL: url})
}
