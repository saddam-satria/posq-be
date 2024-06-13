package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains"
	"github.com/saddam-satria/posq-be/domains/apis"
	"github.com/saddam-satria/posq-be/domains/query"
	"github.com/saddam-satria/posq-be/repositories"
)

func GetAggregateReport(ctx *fiber.Ctx) error {
	code := fiber.StatusOK

	var totalOrderNow query.CountAggregate
	var totalProductOutStock query.CountAggregate
	var totalUnavailableService query.CountAggregate
	repositories.CountTotalOrderByToday(&totalOrderNow)
	repositories.CountProductOutStock(&totalProductOutStock)
	repositories.CountUnavailableService(&totalUnavailableService)

	response := apis.ReportAggregateResponse{
		ProductOutStock:    totalProductOutStock.Count,
		TotalOrder:         totalOrderNow.Count,
		UnavailableService: totalUnavailableService.Count,
	}

	ctx.Status(code).JSON(commons.GetResponse(commons.SUCCESS[domains.En], code, response))
	return nil
}
