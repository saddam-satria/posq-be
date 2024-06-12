package endpoints

import (
	"fmt"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains"
	"github.com/saddam-satria/posq-be/models"
	"github.com/saddam-satria/posq-be/repositories"
)

func GetProducts(ctx *fiber.Ctx) error {
	query := ctx.Queries()
	var products []models.Product

	filter := query["filter"]
	order := query["order"]
	sortQuery := query["sort"]
	keyword := query["keyword"]

	code := fiber.StatusBadRequest
	if order != "" && (order != "desc" && order != "asc") {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.BAD_REQUEST[domains.En], code, nil))
		return nil
	}

	if sortQuery != "" && (sortQuery != "name" && sortQuery != "price") {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.BAD_REQUEST[domains.En], code, nil))
		return nil
	}

	repositories.GetProducts(filter, order, sortQuery, keyword, &products)

	for i := range products {
		totalPrice := float32(0)
		for _, item := range products[i].ProductVariants {
			totalPrice = totalPrice + item.BasePrice + item.SalePrice
		}
		products[i].Price = totalPrice
		maxPrice := products[i].ProductVariants[len(products[i].ProductVariants)-1].BasePrice
		minPrice := products[i].ProductVariants[0].BasePrice
		priceRange := fmt.Sprintf("%.0f - %.0f", minPrice, maxPrice)

		if minPrice == maxPrice {
			priceRange = fmt.Sprintf("%.0f", maxPrice)
		}

		imageUrl := "https://down-id.img.susercontent.com/file/sg-11134201-22100-e3jorlgpxsivae"
		products[i].Image = &imageUrl

		products[i].RangePrice = priceRange
	}

	sort.Slice(products, func(i, j int) bool {
		if order == "desc" {
			return products[i].Price > products[j].Price
		}
		return products[i].Price < products[j].Price
	})

	ctx.Status(fiber.StatusOK).JSON(commons.GetResponse("success", fiber.StatusOK, products))

	return nil
}
