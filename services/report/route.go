package report

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/services/report/endpoints"
)

func ReportRoute(route fiber.Router) {
	route.Get("cashier/dashboard", endpoints.GetAggregateReport)
}
